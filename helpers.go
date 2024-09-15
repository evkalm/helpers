package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"reflect"
	"strings"
)

var colorValue func(a ...interface{}) string
var colorName func(a ...interface{}) string
var colorType func(a ...interface{}) string
var colorFieldName func(a ...interface{}) string
var colorPipe func(a ...interface{}) string
var colorMapKey func(a ...interface{}) string
var colorErr func(a ...interface{}) string

// init
func init() {
	colorValue = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(70)).SprintFunc() // 106, 107, 71,,,108, 10
	colorName = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(220)).SprintFunc()
	colorType = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(60)).SprintFunc()
	colorPipe = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(237)).SprintFunc()
	colorFieldName = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(166)).SprintFunc()
	colorMapKey = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(179)).SprintFunc()
	colorErr = color.New(color.Attribute(38), color.Attribute(5), color.Attribute(1)).SprintFunc()
}

func PrettyPrint(a any) {
	var mess strings.Builder //
	mess.WriteString(colorName(commonTypeName(a)))
	mess.WriteString(" " + typeToString(a) + " = ")
	mess.WriteString(prettyPrintElem(reflect.ValueOf(a).Interface(), 0, false))

	fmt.Println(mess.String())
}

func PrettyJSON[T string | json.RawMessage | []byte](dataJSON T) {
	jsonByte := []byte(dataJSON)
	var data map[string]interface{}

	err := json.Unmarshal(jsonByte, &data)

	var mess strings.Builder
	mess.WriteString(colorName("JSON"))
	mess.WriteString(" " + typeToString(dataJSON) + " = ")

	mess.WriteString(" " + colorErr("Note: JSON Keys Sorted ASC") + " ")

	if err != nil {
		mess.WriteString(fmt.Sprintf("%s: %s",
			colorErr("Error unmarshaling JSON"),
			err,
		))
	} else {
		mess.WriteString(prettyPrintElem(data, 0, true))
	}
	fmt.Println(mess.String())
}

func prettyPrintElem(a any, lev int, isJSON bool) string {
	v := reflect.ValueOf(a)
	t := getEffectiveKind(v.Kind())

	switch t {
	case reflect.Int: // and other primitives
		if isJSON && v.Interface() == "unknow" {
			return colorValue("null")
		}
		return buildPrimitiveMess(v.Interface())

	case reflect.Array, reflect.Slice:
		return buildArrayMess(v.Interface(), lev+1, isJSON)

	case reflect.Struct:
		return buildStructMess(v.Interface(), lev+1, isJSON)

	case reflect.Map:
		return buildMapMess(v.Interface(), lev+1, isJSON)

	case reflect.Ptr:
		if !v.IsZero() {
			return prettyPrintElem(v.Elem().Interface(), lev, isJSON)
		} else {
			return colorValue("nil")
		}

	case reflect.Uintptr:
		return ptrValToMess(v.Interface(), isJSON)

		//case reflect.Chan:
		//	return buildChannelMess(v.Interface())

		//case reflect.Func:
		//case reflect.Interface:
		//	return fmt.Sprintf("%s", v)
		//case reflect.Invalid:
		//case reflect.UnsafePointer:
	}
	if isJSON {
		return colorValue("null")
	}

	return "unknow"
}
