// create field.Fields from type declaration
package extractor

import (
	"fmt"
	"go/ast"
	"log"

	"github.com/slonegd/structstringer/internal/declaration"
	"github.com/slonegd/structstringer/internal/field"
	"github.com/slonegd/structstringer/internal/packinfo"
)

type Extractor interface {
	ExtractFields(typeSpec declaration.Type) (field.Fields, error)
}

func NewExtractor(finder declaration.Finder, packageName string) Extractor {
	return extractor{finder: finder, packageName: packageName}
}

type extractor struct {
	finder      declaration.Finder
	packageName string
}

func (extractor extractor) ExtractFields(typeSpec declaration.Type) (field.Fields, error) {
	structType, ok := typeSpec.Spec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("type %q not a struct", typeSpec.Spec.Name.Name)
	}
	list := structType.Fields.List

	fields := make(field.Fields, 0, len(list))
	for _, pfield := range list {
		if typeIdent, ok := pfield.Type.(*ast.Ident); ok {
			pathToValue := typeSpec.Path + pfield.Names[0].Name
			fields = append(fields, field.Field{
				Name:        pfield.Names[0].Name,
				PathToValue: pathToValue,
				Type:        typeIdent.Name,
				Package:     extractor.packageName,
				Fields:      extractor.fields(typeIdent.Name, pathToValue+"."),
			})
		}

		if typeSelector, ok := pfield.Type.(*ast.SelectorExpr); ok {
			packageName := typeSelector.X.(*ast.Ident).Name
			pathToValue := typeSpec.Path + pfield.Names[0].Name
			fields = append(fields, field.Field{
				Name:        pfield.Names[0].Name, // TODO >1
				PathToValue: pathToValue,
				Type:        typeSelector.Sel.Name,
				Package:     packageName,
				Fields:      extractor.fieldsInOtherPackage(typeSpec.Imports[packageName], typeSelector.Sel.Name, pathToValue+"."),
			})
		}

	}

	return fields, nil
}

func (extractor extractor) fields(typeName, pathToValue string) field.Fields {
	decl, err := extractor.finder.Find(typeName, pathToValue)
	if err != nil {
		return nil
	}
	fields, _ := extractor.ExtractFields(decl)
	return fields
}

func (extractor extractor) fieldsInOtherPackage(packagePath, typeName, pathToValue string) field.Fields {
	packinfo, err := packinfo.GetByPath(packagePath)
	if err != nil {
		log.Fatalf("get package failed: %s", err)
	}
	extractor.finder = declaration.NewFinder(packinfo.GoFiles)
	extractor.packageName = packinfo.Name
	decl, err := extractor.finder.Find(typeName, pathToValue)
	if err != nil {
		return nil
	}
	fields, _ := extractor.ExtractFields(decl)
	return fields
}
