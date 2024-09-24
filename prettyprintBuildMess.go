package helpers

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func buildPrimitiveMess(a any, lev int) string {
	v := reflect.ValueOf(a)

	if !v.IsValid() {
		return "nil"
	}

	if v.Kind() == reflect.String {
		str := "\"" + v.String() + "\""

		rows := strings.Split(str, "\n")
		indent := printIndent(lev)
		for i, row := range rows {
			if row != "" {
				if i == 0 {
					rows[i] = colorValue(row)
				} else {
					rows[i] = indent + colorValue(row)
				}
			}
		}
		return strings.Join(rows, "\n")
	}

	return colorValue(fmt.Sprintf("%v", v.Interface()))
}

func buildStructMess(a any, lev int, isJSON bool) string {
	v := reflect.ValueOf(a)

	var mess strings.Builder
	mess.WriteString("{")

	numFields := v.NumField()
	for i := 0; i < numFields; i++ {

		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i).Interface()

		typeValue := ""
		if !isJSON {
			typeValue = " " + typeToString(fieldValue)
		}

		mess.WriteString(fmt.Sprintf("\n%s%s%s: %s",
			printIndent(lev),
			colorFieldName(fieldName),
			typeValue,
			prettyPrintElem(fieldValue, lev, isJSON),
		))

		if i != numFields-1 {
			mess.WriteString(",")
		}
	}

	mess.WriteString("\n" + printIndent(lev-1) + "}")

	return mess.String()
}

func buildArrayMess(a any, lev int, isJSON bool) string {
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var mess strings.Builder
	mess.WriteString("[")

	var isValuesOnlyPrimitive = true
	for i := 0; i < v.Len(); i++ {
		vi := v.Index(i)
		if vi.Kind() == reflect.Interface {
			vi = vi.Elem()
		}
		if getEffectiveKind(vi.Kind()) != reflect.Int {
			isValuesOnlyPrimitive = false
			break
		}
	}

	for i := 0; i < v.Len(); i++ {
		vi := v.Index(i)
		ti := getEffectiveKind(vi.Kind())

		isPrimit := ti == reflect.Int // and other primitives
		isIface := ti == reflect.Interface

		if isIface {
			vi = vi.Elem()
			ti = getEffectiveKind(vi.Kind())
			isPrimit = ti == reflect.Int
		}

		newRow := "\n"
		indent := printIndent(lev)
		lastIndent := printIndent(lev - 1)
		space := ""
		if isValuesOnlyPrimitive {
			newRow = ""
			indent = ""
			lastIndent = ""
			space = " "
		}

		mess.WriteString(newRow + indent)

		var isPrintType bool
		//isPrintType = (!isValuesOnlyPrimitive && isIface) || (isValuesOnlyPrimitive && !isPrimit)
		if isIface && isPrimit {
			isPrintType = true
		}

		if isJSON {
			isPrintType = false
		}

		if isPrintType {
			mess.WriteString(colorType(typeToString(vi.Interface())))
			mess.WriteString(" ")
		}

		var vvi any
		if !vi.IsValid() {
			vvi = "unknow"
		} else {
			vvi = vi.Interface()
		}

		mess.WriteString(prettyPrintElem(vvi, lev, isJSON))
		if i == v.Len()-1 {
			mess.WriteString(newRow + lastIndent)
		} else {
			mess.WriteString("," + space)
		}
	}
	mess.WriteString("]")

	return mess.String()
}

func buildMapMess(a any, lev int, isJSON bool) string {
	v := reflect.ValueOf(a)
	var mess strings.Builder
	mess.WriteString("{")

	// sort keys
	keys := v.MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j])
	})

	numFields := v.Len()
	for i, key := range keys {
		vi := v.MapIndex(key)

		// print keys
		valWrap := ""
		typeKey := ""
		switch key.Kind() {
		case reflect.String:
			valWrap = "\""
		case reflect.Interface:
			if key.Elem().Kind() == reflect.String {
				valWrap = "\""
			}
			if !isJSON {
				typeKey = " " + typeToString(key.Elem().Interface())
			}
		}
		mess.WriteString(fmt.Sprintf("%s%s%s%s%s:",
			"\n"+printIndent(lev),
			colorMapKey(valWrap),
			colorMapKey(key),
			colorMapKey(valWrap),
			typeKey,
		))

		// print value
		typeValue := ""
		if vi.Kind() == reflect.Interface && !isJSON {
			typeValue = typeToString(vi.Interface()) + " "
		}
		mess.WriteString(fmt.Sprintf(" %s%s",
			colorType(typeValue),
			prettyPrintElem(vi.Interface(), lev, isJSON),
		))

		if i != numFields-1 {
			mess.WriteString(",")
		}
	}

	mess.WriteString("\n" + printIndent(lev-1) + "}")

	return mess.String()
}
