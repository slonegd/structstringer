package recursive_struct

//go:generate go run ../.. -type A
type A struct {
	i    int
	flag bool
	b    B
}

type B struct {
	i    int
	flag bool
}
