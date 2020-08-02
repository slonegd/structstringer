package declaration

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Finder interface {
	Find(typeName string) (*ast.TypeSpec, error)
}

type finder struct {
	files []string
}

// NewFinder - for find type declaration in files
func NewFinder(files []string) Finder {
	return finder{files: files}
}

func (finder finder) Find(typeName string) (*ast.TypeSpec, error) {
	fileset := token.NewFileSet()
	for _, fileName := range finder.files {
		file, err := parser.ParseFile(fileset, fileName, nil, 0)
		if err != nil {
			return nil, fmt.Errorf("parse file: %w", err)
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
