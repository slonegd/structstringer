package recursive

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
			a: A{i: 42, flag: true, b: B{i: 43}},
			// TODO recursive.B - package name
			want: `
recursive.A{
	i    int  42
	flag bool true
	b    B    {
		i    int  42
		flag bool true
		c    C    {
			i    int  42
			flag bool true
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
