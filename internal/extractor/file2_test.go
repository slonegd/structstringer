//nolint
package extractor

import "github.com/slonegd/structstringer/examples/simple"

type B struct {
	i    int
	flag bool
}

type C int

type D struct {
	i int
	b B
}

type E struct {
	i int
	b simple.B
}
