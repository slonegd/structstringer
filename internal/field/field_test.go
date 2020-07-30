package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFields_String(t *testing.T) {
	tests := []struct {
		name   string
		fields Fields
		want   string
	}{
		{
			name: "happy path",
			fields: []Field{
				{Name: "i", Type: "int"},
				{Name: "flag", Type: "bool"},
			},
			want: `
	builder.WriteString("\n\ti    int  ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.flag))`,
		},
		{
			name:   "nil arg",
			fields: nil,
			want:   ``,
		},
		{
			name: "not implemented",
			fields: []Field{
				{Name: "i", Type: "int"},
				{Name: "flag", Type: "bools"},
			},
			want: `
	builder.WriteString("\n\ti    int   ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tflag bools ")
	builder.WriteString("not_implemented")`,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.fields.String())
	}
}
