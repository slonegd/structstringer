package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slonegd/structstringer/internal/generator"
)

var (
	typeName = flag.String("type", "", "type name; must be set")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("structstringer: ")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of structstringer:\n")
		fmt.Fprintf(os.Stderr, "\tstructstringer -type T\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	command := strings.Join(os.Args[1:], " ")
	log.Printf("generate structstringer %s", command)

	generator.Generate(*typeName)
}
