// find declaration of the type
package declaration

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Finder interface {
	Find(typeName, pathToValue string) (Type, error)
}

type Type struct {
	Spec    *ast.TypeSpec // TODO dont use ast package for out
	Imports map[string]string
	Path    string // TODO move to field
}

// NewFinder - for find type declaration in files
func NewFinder(files []string) Finder {
	return finder{files: files}
}

type finder struct {
	files []string
}

func (finder finder) Find(typeName, pathToValue string) (Type, error) {
	result := Type{}

	fileset := token.NewFileSet()
	for _, fileName := range finder.files {
		file, err := parser.ParseFile(fileset, fileName, nil, 0)
		if err != nil {
			return result, fmt.Errorf("parse file: %w", err)
		}
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				spec := genDecl.Specs[0].(*ast.TypeSpec)
				if spec.Name.Name == typeName {
					result.Spec = spec
					result.Imports = imports(file)
					result.Path = pathToValue
					return result, nil
				}
			}
		}
	}
	return result, fmt.Errorf("cant find type %q", typeName)
}

func imports(file *ast.File) map[string]string {
	result := make(map[string]string)
	for _, importSpec := range file.Imports {
		packageName := ""
		path := strings.ReplaceAll(importSpec.Path.Value, "\"", "")
		if importSpec.Name == nil || importSpec.Name.Name == "." {
			packageName = fromPath(path)
		} else {
			packageName = importSpec.Name.Name
		}
		result[packageName] = path
	}
	return result
}

func fromPath(path string) string {
	names := strings.Split(path, "/")
	if len(names) == 0 {
		return ""
	}
	return names[len(names)-1]
}
