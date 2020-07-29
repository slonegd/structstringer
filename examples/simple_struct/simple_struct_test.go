package simple_struct

import (
	"fmt"
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
			a: A{i: 42, flag: true},
			want: `
A{
	i    int  42
	flag bool true
}`,
			fmtWant: "simple_struct.A{i:42, flag:true}",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.a.String())
		assert.Equal(t, tt.fmtWant, fmt.Sprintf("%#v", tt.a))
	}
}

func Benchmark_A_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		a := randomA()
		b.StartTimer()
		a.String()
		b.StopTimer()
	}
}

func Benchmark_A_fmt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		a := randomA()
		b.StartTimer()
		fmt.Sprintf("%#v", a)
		b.StopTimer()
	}
}

func randomA() A {
	return A{
		i:    rand.Int(),
		flag: rand.Int()%2 == 0,
	}
}
