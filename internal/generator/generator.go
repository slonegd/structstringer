package generator

import (
	"log"

	"github.com/slonegd/structstringer/internal/declaration"
	"github.com/slonegd/structstringer/internal/extractor"
	"github.com/slonegd/structstringer/internal/packinfo"
	"github.com/slonegd/structstringer/internal/printer"
	"github.com/slonegd/structstringer/internal/saver"
)

func Generate(typeName string) {
	pkg, err := packinfo.Get()
	catchError(err)

	finder := declaration.NewFinder(pkg.GoFiles)
	typeSpec, err := finder.Find(typeName, "")
	catchError(err)

	extractor := extractor.NewExtractor(finder, pkg.Name)
	fields, err := extractor.ExtractFields(typeSpec)
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
