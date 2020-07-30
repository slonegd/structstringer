package package_info

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func Get() (*packages.Package, error) {
	pkgs, err := packages.Load(nil, ".")
	if err != nil {
		return nil, fmt.Errorf("load packages: %w", err)
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("must be only one package, got: %d", len(pkgs))
	}
	return pkgs[0], nil

}
