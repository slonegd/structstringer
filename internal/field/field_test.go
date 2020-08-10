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
				{Name: "i", PathToValue: "i", Type: "int", Package: "field"},
				{Name: "flag", PathToValue: "flag", Type: "bool", Package: "field"},
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
			name: "not implemented unknow type",
			fields: Fields{
				{Name: "i", PathToValue: "i", Type: "int", Package: "field"},
				{Name: "flag", PathToValue: "flag", Type: "bools", Package: "field"},
			},
			want: `
	builder.WriteString("\n\ti    int   ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tflag bools ")
	builder.WriteString("not_implemented_unknow_type")`,
		},
		{
			name: "recursive",
			fields: Fields{
				{Name: "i", PathToValue: "i", Type: "int", Package: "field"},
				{Name: "b", PathToValue: "b", Type: "B", Package: "field", Fields: Fields{
					{Name: "i", PathToValue: "b.i", Type: "int", Package: "field"},
					{Name: "flag", PathToValue: "b.flag", Type: "bool", Package: "field"},
				}},
			},
			want: `
	builder.WriteString("\n\ti int     ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tb field.B ")
	builder.WriteRune('{')
	builder.WriteString("\n\t\ti    int  ")
	builder.WriteString(strconv.Itoa(t.b.i))
	builder.WriteString("\n\t\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.b.flag))
	builder.WriteString("\n\t}")`,
		},
		{
			name: "unexported field from other package",
			fields: Fields{
				{Name: "i", PathToValue: "i", Type: "int", Package: "field"},
				{Name: "b", PathToValue: "b", Type: "B", Package: "other", IsOtherPackage: true, Fields: Fields{
					{Name: "I", PathToValue: "b.I", Type: "int", Package: "other"},
					{Name: "flag", PathToValue: "b.flag", Type: "bool", Package: "other"},
				}},
			},
			want: `
	builder.WriteString("\n\ti int     ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tb other.B ")
	builder.WriteRune('{')
	builder.WriteString("\n\t\tI    int  ")
	builder.WriteString(strconv.Itoa(t.b.I))
	builder.WriteString("\n\t\tflag bool ")
	builder.WriteString("not_implemented_unexported_fields")
	builder.WriteString("\n\t}")`,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.fields.String())
	}
}
