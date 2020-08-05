package field

import (
	"fmt"
)

// Field - simple struct represented one field of the struct
type Field struct {
	Name, Type, Package        string
	Fields                     Fields
	allignedName, allignedType string
	tabs                       string
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
		return fmt.Sprintf(`builder.WriteRune('{')%s
	builder.WriteString("\n%s}")`, field.Fields, field.tabs)
	}
	switch field.Type {
	case "bool":
		return fmt.Sprintf("builder.WriteString(strconv.FormatBool(t.%s))", field.Name)
	case "string":
		return fmt.Sprintf("builder.WriteString(t.%s)", field.Name)
	case "int":
		return fmt.Sprintf("builder.WriteString(strconv.Itoa(t.%s))", field.Name)
	case "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32":
		return fmt.Sprintf("builder.WriteString(strconv.Itoa(int(t.%s)))", field.Name)
	case "byte":
		return fmt.Sprintf(`builder.WriteString("0x")
		builder.WriteString(strconv.FormatUint(uint64(t.%s), 16))`, field.Name)
	case "uint64":
		return fmt.Sprintf("builder.WriteString(strconv.FormatUint(t.%s, 10))", field.Name)
	case "uintptr":
		return fmt.Sprintf(`builder.WriteString("0x")
		builder.WriteString(strconv.FormatUint(uint64(t.%s), 16))`, field.Name)
	case "rune":
		return fmt.Sprintf("builder.WriteString(string(t.%s))", field.Name)
	case "float64":
		return fmt.Sprintf("builder.WriteString(strconv.FormatFloat(t.%s, 'e', -1, 64))", field.Name)
	case "float32":
		return fmt.Sprintf("builder.WriteString(strconv.FormatFloat(float64(t.%s), 'e', -1, 32))", field.Name)
	case "complex64", "complex128":
		return fmt.Sprintf(`builder.WriteString(fmt.Sprintf("%%g", t.%s))`, field.Name)
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
