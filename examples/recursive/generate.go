package recursive

//go:generate go run ../.. -type A

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
}