// Code generated by structstringer
// DO NOT EDIT!
package recursive

import (
	"strconv"
	"strings"
)

func (t A) String() string {
	var builder strings.Builder
	builder.Grow(1024) // TODO count
	builder.WriteString("\nrecursive.A{")
	builder.WriteString("\n\ti    int  ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.flag))
	builder.WriteString("\n\tb    B    ")
	builder.WriteRune('{')
	builder.WriteString("\n\t\ti    int  ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\t\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.flag))
	builder.WriteString("\n\t\tc    C    ")
	builder.WriteRune('{')
	builder.WriteString("\n\t\t\ti    int  ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\t\t\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.flag))
	builder.WriteString("\n\t\t}")
	builder.WriteString("\n\t}")
	builder.WriteString("\n}")
	return builder.String()
}
