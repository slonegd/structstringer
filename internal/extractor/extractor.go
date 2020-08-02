package extractor

import (
	"fmt"
	"go/ast"

	"github.com/slonegd/structstringer/internal/declaration"
	"github.com/slonegd/structstringer/internal/field"
)

type Extractor interface {
	ExtractFields(typeSpec *ast.TypeSpec) (field.Fields, error)
}

func NewExtractor(finder declaration.Finder) Extractor {
	return extractor{finder: finder}
}

type extractor struct {
	finder declaration.Finder
}

func (extractor extractor) ExtractFields(typeSpec *ast.TypeSpec) (field.Fields, error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("type %q not a struct", typeSpec.Name.Name)
	}
	list := structType.Fields.List

	fields := make(field.Fields, 0, len(list))
	for _, pfield := range list {
		typeName := pfield.Type.(*ast.Ident).Name
		fields = append(fields, field.Field{
			Name:   pfield.Names[0].Name,
			Type:   typeName,
			Fields: extractor.fields(typeName),
		})
	}

	return fields, nil
}

func (extractor extractor) fields(typeName string) field.Fields {
	decl, err := extractor.finder.Find(typeName)
	if err != nil {
		return nil
	}
	fields, _ := extractor.ExtractFields(decl)
	return fields
}
