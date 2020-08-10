package recursive

import "github.com/slonegd/structstringer/examples/simple"

//go:generate go run ../.. -type A
//go:generate go run ../.. -type D
//go:generate go run ../.. -type E

// A - test struct
type A struct {
	i    int
	flag bool
	b    B
}

// B - test struct
type B struct {
	i    int
	flag bool
	c    C
}

// C - test struct
type C struct {
	i    int
	flag bool
}

// D - test struct
type D struct {
	i    int
	flag bool
	b    simple.B
}

// E - test struct
type E struct {
	i    int
	flag bool
	c    simple.C
}
