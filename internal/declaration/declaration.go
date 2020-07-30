package declaration

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/slonegd/structstringer/internal/field"
)

func Find(files []string, typeName string) (*ast.TypeSpec, error) {
	fileset := token.NewFileSet()
	for _, fileName := range files {
		file, err := parser.ParseFile(fileset, fileName, nil, 0)
		if err != nil {
			return nil, fmt.Errorf("pars file %q: %w", fileName, err)
		}
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				result := genDecl.Specs[0].(*ast.TypeSpec)
				if result.Name.Name == typeName {
					return result, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("cant find type %q", typeName)
}

func ExtractFields(typeSpec *ast.TypeSpec) (field.Fields, error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, errors.New("type not a struct")
	}
	list := structType.Fields.List
	pfields := make([]*ast.Field, 0, len(list))
	for _, f := range list {
		pfields = append(pfields, f)
	}

	fields := make(field.Fields, 0, len(pfields))
	for _, pfield := range pfields {
		fields = append(fields, field.Field{
			Name: pfield.Names[0].Name,
			Type: pfield.Type.(*ast.Ident).Name,
		})
	}

	return fields, nil
}
