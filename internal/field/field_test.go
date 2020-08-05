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
			fields: Fields{
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
			fields: Fields{
				{Name: "i", Type: "int"},
				{Name: "flag", Type: "bools"},
			},
			want: `
	builder.WriteString("\n\ti    int   ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tflag bools ")
	builder.WriteString("not_implemented")`,
		},
		{
			name: "recursive",
			fields: Fields{
				{Name: "i", Type: "int"},
				{Name: "b", Type: "B", Package: "field", Fields: Fields{
					{Name: "i", Type: "int"},
					{Name: "flag", Type: "bool"},
				}},
			},
			want: `
	builder.WriteString("\n\ti int     ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tb field.B ")
	builder.WriteRune('{')
	builder.WriteString("\n\t\ti    int  ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\t\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.flag))
	builder.WriteString("\n\t}")`,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.fields.String())
	}
}
