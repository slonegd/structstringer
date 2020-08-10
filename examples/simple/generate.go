package simple

//go:generate go run ../.. -type A

// A - test struct
type A struct {
	flag bool
	str  string
	i    int
	i8   int8
	i16  int16
	i32  int32
	i64  int64
	ui   uint
	ui8  uint8
	ui16 uint16
	ui64 uint64
	b    byte
	p    uintptr
	r    rune
	f64  float64
	f32  float32
	c64  complex64
	c128 complex128
}

// B - test struct
type B struct {
	I int
}

// C - test struct
type C struct {
	I    int
	flag bool
}
