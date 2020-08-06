// stringer of the fields for generator
package field

import (
	"fmt"
)

// Field - simple struct represented one field of the struct
type Field struct {
	Name, Type, Package, PathToValue string
	Fields                           Fields
	allignedName, allignedType       string
	tabs                             string
}

func (field Field) String() string {
	return fmt.Sprintf(`
	%s
	%s`, field.generateDescription(), field.generateStringer())
}

// Fields - slice of Field
type Fields []Field

func (fields Fields) String() string {
	result := ""
	if len(fields) == 0 {
		return result
	}
	if fields[0].tabs == "" {
		fields.setTabs("\\t")
	}
	for _, field := range alignWight(fields) {
		result += field.String()
	}
	return result
}

func (fields *Fields) setTabs(tabs string) {
	for i := range *fields {
		(*fields)[i].tabs = tabs
	}
}

func (field Field) generateDescription() string {
	return fmt.Sprintf(`builder.WriteString("\n%s%s %s ")`, field.tabs, field.allignedName, field.allignedType)
}

func (field Field) generateStringer() string {
	if field.Fields != nil {
		field.Fields.setTabs(field.tabs + "\\t")
		// field.Fields.setPathTovalue(field.PathToValue + field.Name + ".")
		return fmt.Sprintf(`builder.WriteRune('{')%s
	builder.WriteString("\n%s}")`, field.Fields, field.tabs)
	}
	switch field.Type {
	case "bool":
		return fmt.Sprintf("builder.WriteString(strconv.FormatBool(t.%s))", field.PathToValue)
	case "string":
		return fmt.Sprintf("builder.WriteString(t.%s)", field.PathToValue)
	case "int":
		return fmt.Sprintf("builder.WriteString(strconv.Itoa(t.%s))", field.PathToValue)
	case "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32":
		return fmt.Sprintf("builder.WriteString(strconv.Itoa(int(t.%s)))", field.PathToValue)
	case "byte", "uintptr":
		return fmt.Sprintf(`builder.WriteString("0x")
		builder.WriteString(strconv.FormatUint(uint64(t.%s), 16))`, field.PathToValue)
	case "uint64":
		return fmt.Sprintf("builder.WriteString(strconv.FormatUint(t.%s, 10))", field.PathToValue)
	case "rune":
		return fmt.Sprintf("builder.WriteString(string(t.%s))", field.PathToValue)
	case "float64":
		return fmt.Sprintf("builder.WriteString(strconv.FormatFloat(t.%s, 'e', -1, 64))", field.PathToValue)
	case "float32":
		return fmt.Sprintf("builder.WriteString(strconv.FormatFloat(float64(t.%s), 'e', -1, 32))", field.PathToValue)
	case "complex64", "complex128":
		return fmt.Sprintf(`builder.WriteString(fmt.Sprintf("%%g", t.%s))`, field.PathToValue)
	default:
		return `builder.WriteString("not_implemented")`
	}
}

func alignWight(fields Fields) Fields {
	nameWight := 0
	typeWight := 0
	for i, field := range fields {
		if len(field.Name) > nameWight {
			nameWight = len(field.Name)
		}
		if field.Fields != nil {
			fields[i].Type = field.Package + "." + field.Type
		}
		if len(fields[i].Type) > typeWight {
			typeWight = len(fields[i].Type)
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
