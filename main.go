package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	typeName = flag.String("type", "", "type name; must be set")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of structstringer:\n")
	fmt.Fprintf(os.Stderr, "\tstructstringer -type T\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("structstringer: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	command := strings.Join(os.Args[1:], " ")
	log.Printf("generate structstringer %s", command)

	pkgs, err := packages.Load(nil, ".")
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}

	log.Printf("files %v", pkgs[0].GoFiles) // TODO debug

	fileset := token.NewFileSet()
	// for work with ast
	file, err := parser.ParseFile(fileset, pkgs[0].GoFiles[0], nil, 0) // FIX pkgs[0].GoFiles[0]
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n%+v", file) // TODO debug

	// TODO inspector
	findTypeSpec := func(typeName string) *ast.TypeSpec {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				result := genDecl.Specs[0].(*ast.TypeSpec)
				if result.Name.Name == typeName {
					return result
				}
			}
		}
		return nil
	}

	typeSpec := findTypeSpec(*typeName)
	log.Printf("\n%+v", typeSpec) // TODO debug

	pfields := typeSpec.Type.(*ast.StructType).Fields.List
	fields := make([]ast.Field, 0, len(pfields))
	for _, f := range pfields {
		fields = append(fields, *f)
	}

	log.Printf("\n%#v", fields) // TODO debug

	fieldNames := make([]Field, 0, len(fields))
	for _, field := range fields {
		fieldNames = append(fieldNames, Field{
			Name: field.Names[0].Name,
			Type: field.Type.(*ast.Ident).Name,
		})
	}

	log.Printf("%+v", fieldNames) // TODO debug

	data := generateFile(pkgs[0].Name, *typeName, alignWight(fieldNames))

	log.Printf("%s", data) // TODO debug

	fmtData, err := format.Source([]byte(data))
	if err != nil {
		log.Fatalf("format source failed: %s", err)
	}

	err = ioutil.WriteFile(strings.ToLower(*typeName+"_string.go"), fmtData, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

type Field struct {
	Name, Type, allignedType string
}

func (field Field) String() string {
	return fmt.Sprintf(`	builder.WriteString(%s)
	builder.WriteString(%s)
`, field.generateDescription(), field.generateStringer())
}

func (field Field) generateDescription() string {
	return fmt.Sprintf(`"\n\t%s %s "`, field.Name, field.allignedType)
}

func (field Field) generateStringer() string {
	switch field.Type {
	case "int":
		return fmt.Sprintf("strconv.Itoa(t.%s)", field.Name)
	case "bool":
		return fmt.Sprintf("strconv.FormatBool(t.%s)", field.Name)
	default:
		return "not_implemented"
	}
}

func alignWight(fields []Field) []Field {
	nameWight := 0
	typeWight := 0
	for _, field := range fields {
		if len(field.Name) > nameWight {
			nameWight = len(field.Name)
		}
		if len(field.Type) > typeWight {
			typeWight = len(field.Type)
		}
	}
	for i, field := range fields {
		fields[i].Name = growWight(field.Name, nameWight)
		fields[i].allignedType = growWight(field.Type, typeWight)
	}
	return fields
}

func growWight(v string, wight int) string {
	for len(v) < wight {
		v += " "
	}
	return v
}

// // это должно получиться
// func (t Type) String() string {
// 	var builder strings.Builder
// 	// TODO use Grow
// 	builder.WriteString(`Type{`)
// 	builder.WriteString(`
// 	i int `)
// 	builder.WriteString(strconv.Itoa(t.i))
// 	builder.WriteString(`
// 	flag bool `)
// 	builder.WriteString(strconv.FormatBool(t.flag))
// 	builder.WriteString(`
// }`)
// 	return builder.String()
// }

func generateFile(packageName, typeName string, fields []Field) string {
	result := strings.ReplaceAll(fileTemplate, "$(package)", packageName)
	result = strings.ReplaceAll(result, "$(type)", typeName)

	fieldsString := ""
	for _, field := range fields {
		fieldsString += field.String()
	}
	result = strings.ReplaceAll(result, "$(fields)", fieldsString)
	return result
}

var fileTemplate = `// Code generated by structstringer
// DO NOT EDIT!
package $(package)

import "strings"
import "strconv"

func (t $(type)) String() string {
	var builder strings.Builder
	builder.Grow(80) // TODO count
	builder.WriteString("\n$(type){")
$(fields)
	builder.WriteString("\n}")
	return builder.String()
}
`
