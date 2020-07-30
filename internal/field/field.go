package field

import (
	"fmt"
)

type Field struct {
	Name, Type                 string
	allignedName, allignedType string
}

func (field Field) String() string {
	return fmt.Sprintf(`
	builder.WriteString(%s)
	builder.WriteString(%s)`, field.generateDescription(), field.generateStringer())
}

type Fields []Field

func (fields Fields) String() string {
	result := ""
	for _, field := range alignWight(fields) {
		result += field.String()
	}
	return result
}

func (field Field) generateDescription() string {
	return fmt.Sprintf(`"\n\t%s %s "`, field.allignedName, field.allignedType)
}

func (field Field) generateStringer() string {
	switch field.Type {
	case "int":
		return fmt.Sprintf("strconv.Itoa(t.%s)", field.Name)
	case "bool":
		return fmt.Sprintf("strconv.FormatBool(t.%s)", field.Name)
	default:
		return `"not_implemented"`
	}
}

func alignWight(fields Fields) Fields {
	nameWight := 0
	typeWight := 0
	for _, field := range fields {
		if len(field.Name) > nameWight {
			nameWight = len(field.Name)
		}
		if len(field.Type) > typeWight {
			typeWight = len(field.Type)
		}
	}
	for i, field := range fields {
		fields[i].allignedName = growWight(field.Name, nameWight)
		fields[i].allignedType = growWight(field.Type, typeWight)
	}
	return fields
}

func growWight(v string, wight int) string {
	for len(v) < wight {
		v += " "
	}
	return v
}
