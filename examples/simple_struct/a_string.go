// Code generated by structstringer
// DO NOT EDIT!
package simple_struct

import "strings"
import "strconv"

func (t A) String() string {
	var builder strings.Builder
	builder.Grow(80) // TODO count
	builder.WriteString("\nsimple_struct.A{")
	builder.WriteString("\n\ti    int  ")
	builder.WriteString(strconv.Itoa(t.i))
	builder.WriteString("\n\tflag bool ")
	builder.WriteString(strconv.FormatBool(t.flag))
	builder.WriteString("\n}")
	return builder.String()
}
