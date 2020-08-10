//nolint
package recursive

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/slonegd/structstringer/examples/simple"
	"github.com/stretchr/testify/assert"
)

func TestAString(t *testing.T) {
	tests := []struct {
		a       A
		want    string
		fmtWant string
	}{
		{
			a: A{i: 42, flag: true, b: B{i: 43}},
			want: `
recursive.A{
	i    int         42
	flag bool        true
	b    recursive.B {
		i    int         43
		flag bool        false
		c    recursive.C {
			i    int  0
			flag bool false
		}
	}
}`,
			fmtWant: "recursive.A{i:42, flag:true, b:recursive.B{i:43, flag:false, c:recursive.C{i:0, flag:false}}}",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.a.String())
		assert.Equal(t, tt.fmtWant, fmt.Sprintf("%#v", tt.a))
	}
}

func TestDString(t *testing.T) {
	tests := []struct {
		d       D
		want    string
		fmtWant string
	}{
		{
			d: D{i: 42, flag: true, b: simple.B{I: 43}},
			want: `
recursive.D{
	i    int      42
	flag bool     true
	b    simple.B {
		I int 43
	}
}`,
			fmtWant: "recursive.D{i:42, flag:true, b:simple.B{I:43}}",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.d.String())
		assert.Equal(t, tt.fmtWant, fmt.Sprintf("%#v", tt.d))
	}
}

func TestEString(t *testing.T) {
	tests := []struct {
		e       E
		want    string
		fmtWant string
	}{
		{
			e: E{i: 42, flag: true, c: simple.C{I: 43}},
			want: `
recursive.E{
	i    int      42
	flag bool     true
	c    simple.C {
		I    int  43
		flag bool not_implemented_unexported_fields
	}
}`,
			fmtWant: "recursive.E{i:42, flag:true, c:simple.C{I:43, flag:false}}",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.e.String())
		assert.Equal(t, tt.fmtWant, fmt.Sprintf("%#v", tt.e))
	}
}

func BenchmarkRecursiveAString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		a := randomA()
		b.StartTimer()
		_ = a.String()
		b.StopTimer()
	}
}

func BenchmarkRecursiveAfmt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		a := randomA()
		b.StartTimer()
		_ = fmt.Sprintf("%#v", a)
		b.StopTimer()
	}
}

func randomA() A {
	return A{
		i:    rand.Int(),
		flag: rand.Int()%2 == 0,
		b: B{
			i:    rand.Int(),
			flag: rand.Int()%2 == 0,
			c: C{
				i:    rand.Int(),
				flag: rand.Int()%2 == 0,
			},
		},
	}
}
