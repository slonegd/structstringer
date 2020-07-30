package generator

import (
	"log"

	"github.com/slonegd/structstringer/internal/declaration"
	"github.com/slonegd/structstringer/internal/package_info"
	"github.com/slonegd/structstringer/internal/printer"
	"github.com/slonegd/structstringer/internal/saver"
)

func Generate(typeName string) {
	pkg, err := package_info.Get()
	catchError(err)

	typeSpec, err := declaration.Find(pkg.GoFiles, typeName)
	catchError(err)

	fields, err := declaration.ExtractFields(typeSpec)
	catchError(err)

	data := printer.String(pkg.Name, typeName, fields.String())

	err = saver.Save(data, typeName)
	catchError(err)
}

func catchError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
