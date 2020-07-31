package recursive_struct

import (
	"fmt"
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
			want: `
recursive_struct.A{
	i    int  42
	flag bool true
	b    B    not_implemented
}`,
			fmtWant: "recursive_struct.A{i:42, flag:true, b:recursive_struct.B{i:43, flag:false}}",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.a.String())
		assert.Equal(t, tt.fmtWant, fmt.Sprintf("%#v", tt.a))
	}
}
