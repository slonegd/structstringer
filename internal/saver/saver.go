package saver

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"strings"
)

// Save - format source and save to file
func Save(source, typeName string) error {
	fmtSource, err := format.Source([]byte(source))
	if err != nil {
		return fmt.Errorf("format source: %w", err)
	}

	err = ioutil.WriteFile(strings.ToLower(typeName+"_string.go"), fmtSource, 0777)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
