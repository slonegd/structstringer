//nolint
package simple

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA_String(t *testing.T) {
	tests := []struct {
		a       A
		want    string
		fmtWant string
	}{
		{
			a: A{
				flag: true,
				str:  "forty two",
				i:    42,
				i8:   -8,
				i16:  -16,
				i32:  -32,
				i64:  -64,
				ui:   1,
				ui8:  8,
				ui16: 16,
				ui64: math.MaxUint64,
				b:    0xFF,
				p:    0x001122334455667788,
				r:    '©',
				f64:  0.42,
				f32:  0.84,
				c64:  complex(5, -12),
				c128: complex(12, -5),
			},
			want: `
simple.A{
	flag bool       true
	str  string     forty two
	i    int        42
	i8   int8       -8
	i16  int16      -16
	i32  int32      -32
	i64  int64      -64
	ui   uint       1
	ui8  uint8      8
	ui16 uint16     16
	ui64 uint64     18446744073709551615
	b    byte       0xff
	p    uintptr    0x1122334455667788
	r    rune       ©
	f64  float64    4.2e-01
	f32  float32    8.4e-01
	c64  complex64  (5-12i)
	c128 complex128 (12-5i)
}`,
			fmtWant: `simple.A{flag:true, str:"forty two", i:42, i8:-8, i16:-16, i32:-32, i64:-64, ui:0x1, ui8:0x8, ui16:0x10, ui64:0xffffffffffffffff, b:0xff, p:0x1122334455667788, r:169, f64:0.42, f32:0.84, c64:(5-12i), c128:(12-5i)}`,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.a.String())
		assert.Equal(t, tt.fmtWant, fmt.Sprintf("%#v", tt.a))
	}
}

func BenchmarkAString(b *testing.B) {
	b.ResetTimer()
	a := randomA()
	b.StartTimer()
	a.String()
	b.StopTimer()

}

func BenchmarkAfmt(b *testing.B) {
	b.ResetTimer()
	a := randomA()
	b.StartTimer()
	fmt.Sprintf("%#v", a)
	b.StopTimer()

}

func randomA() A {
	return A{
		flag: rand.Int()%2 == 0,
		str:  "forty two",
		i:    rand.Int(),
		i8:   int8(rand.Int()),
		i16:  int16(rand.Int()),
		i32:  int32(rand.Int()),
		i64:  int64(rand.Int()),
		ui:   uint(rand.Int()),
		ui8:  uint8(rand.Int()),
		ui16: uint16(rand.Int()),
		ui64: uint64(rand.Int()),
		b:    byte(rand.Int()),
		p:    uintptr(rand.Int()),
		r:    rune(rand.Int()),
		f64:  float64(rand.Int()),
		f32:  float32(rand.Int()),
		c64:  complex(float32(rand.Int()), float32(rand.Int())),
		c128: complex(float64(rand.Int()), float64(rand.Int())),
	}
}
