package simple_struct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA_String(t *testing.T) {
	tests := []struct {
		a    A
		want string
	}{
		{
			a: A{i: 42, flag: true},
			want: `
A{
	i    int  42
	flag bool true
}`,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.a.String())
	}
}
