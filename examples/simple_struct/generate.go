package simple_struct

//go:generate go run ../.. -type=A
type A struct {
	i    int
	flag bool
}
