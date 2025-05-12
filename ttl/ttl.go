package ttl

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var types = map[reflect.Type]string{}

func typeName(t reflect.Type) string {
	n := t.Name()
	if n == "" {
		n = t.String()
	}
	return n
}

// fieldName returns the yaml field if present, otherwise the struct field name.
func fieldName(f reflect.StructField) string {
	if f.Tag.Get("yaml") != "" {
		// parse the yaml tag (we only need the field name)
		tag := f.Tag.Get("yaml")
		parts := strings.Split(tag, ",")
		if len(parts) > 0 {
			yamlFieldName := strings.TrimSpace(parts[0])
			if yamlFieldName != "" {
				return yamlFieldName
			}
		}
	}
	if f.Name == "" {
		fmt.Printf("WARNING: struct field %s has no name\n", f.Type.String())
	}
	return f.Name
}

// TypeToLua constructs a EmmyLua Annotation from a Go type.
func TypeToLua(t reflect.Type) string {
	if s, ok := types[t]; ok {
		return s
	}
	switch k := t.Kind(); k {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	case reflect.String:
		return "string"
	case reflect.Slice:
		return TypeToLua(t.Elem()) + "[]"
	case reflect.Map:
		keyT := TypeToLua(t.Key())
		valT := TypeToLua(t.Elem())
		return fmt.Sprintf("table<%s, %s>", keyT, valT)
	case reflect.Struct:
		fields := make([]string, 0, t.NumField())
		for i := range t.NumField() {
			field := t.Field(i)
			fn := fieldName(field)
			if fn == "-" || strings.HasPrefix(fn, "#") {
				continue
			}
			fields = append(fields, fmt.Sprintf("%s: %s", fn, TypeToLua(field.Type)))
		}
		tStr := fmt.Sprintf("{%s}", strings.Join(fields, ", "))
		types[t] = tStr
		return typeName(t)
	case reflect.Ptr:
		// There are no pointers in Ba Sing Se
		return TypeToLua(t.Elem())
	case reflect.Interface:
		if t.NumMethod() == 0 { // interface{}
			return "any"
		}
		fallthrough
	default:
		fmt.Fprintf(os.Stderr, "Unexpected type %s of kind %s\n", t.String(), k.String())
		return "any"
	}
}

func String() string {
	sb := strings.Builder{}
	for t, def := range types {
		sb.WriteString(fmt.Sprintf("---@alias %s %s\n", typeName(t), def))
	}
	return sb.String()
}
