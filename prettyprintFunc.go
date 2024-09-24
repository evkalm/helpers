package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func commonTypeName(a any) string {
	v := reflect.ValueOf(a)

	if a == nil {
		return "invalid"
	}

	t := getEffectiveKind(v.Kind())
	switch t {
	case reflect.Int: // and other primitives
		return "primitive"

	case reflect.Array:
		return "array"

	case reflect.Slice:
		return "slice"

	case reflect.Map:
		return "map"

	case reflect.Struct:
		return "struct"

	case reflect.Chan:
		return "chan"

	case reflect.Func:
		return "func"

	case reflect.Interface:
		return "interface"

	case reflect.Ptr:
		if !v.IsZero() {
			return commonTypeName(v.Elem().Interface())
		}

	case reflect.Uintptr:
		return "uintptr"

	case reflect.Invalid:
		return "invalid"

	case reflect.UnsafePointer:
		return "unsafePointer"
	}

	return "invalid"
}

func isPrimitive(t reflect.Kind) bool {
	switch t {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32,
		reflect.Float64, reflect.String, reflect.Complex64, reflect.Complex128:
		// reflect.Uintptr ??
		return true
	default:
		return false
	}
}

func getEffectiveKind(t reflect.Kind) reflect.Kind {
	if isPrimitive(t) {
		return reflect.Int
	}
	return t
}

func ptrValToMess(a any, isJSON bool) string {
	v := reflect.ValueOf(a)

	t := v.Type().Kind()
	if isPrimitive(t) {
		return buildPrimitiveMess(a, 0)
	}
	switch t {
	case reflect.Uintptr:
		return buildPrimitiveMess(a, 0)

	case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map:
		return prettyPrintElem(v.Interface(), 1, isJSON)
	}

	return ""
}
func typeToString(a any) string {
	v := reflect.ValueOf(a)

	if !v.IsValid() {
		return colorType("(invalid)")
	}

	typeStr := v.Type().String()

	switch typeStr {
	case "uint8":
		typeStr = "uint8, byte"
	case "*uint8":
		typeStr = "*uint8, *byte"
	case "int32":
		typeStr = "int32, rune"
	case "*int32":
		typeStr = "*int32, *rune"
	}
	return fmt.Sprintf("%s", colorType("("+typeStr+")"))
}

func printIndent(lev int) string {
	return strings.Repeat(colorPipe("│   "), lev)
}
