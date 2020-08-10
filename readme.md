[![Go Report Card](https://goreportcard.com/badge/github.com/slonegd/structstringer)](https://goreportcard.com/report/github.com/slonegd/structstringer)
[![github action](https://github.com/slonegd/structstringer/workflows/Go/badge.svg?branch=develop)](https://github.com/slonegd/structstringer/actions?query=branch%3Adevelop++)

structstringer - generator of stringer interface for structs

One example (more in `examples` folder):
```go
package simple

type A struct {
	i    int
	flag bool
	b    B
}

type B struct {
	i    int
	flag bool
}

// generated stringer
simple.A{
	i    int      42
	flag bool     true
	b    simple.B {
		i    int  43
		flag bool false
	}
}

// %#v formater
simple.A{i:42, flag:true, b:simple.B{i:43, flag:false}}

BenchmarkRecursiveAString-4   	1000000000	         0.000780 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecursiveAfmt-4      	1000000000	         0.00186 ns/op	       0 B/op	       0 allocs/op
```

TODO:
 * [v] basic types in POD struct
 * [v] struct fields
 * [v] struct from other package field
 * pointers
 * enums
 * aliases
 * interfaces
 * slices and maps
 * unexported fields