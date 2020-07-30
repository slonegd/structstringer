package declaration

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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
