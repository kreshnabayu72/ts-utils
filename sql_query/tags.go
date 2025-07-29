package sql_query

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func traverseNestedStruct[T any](typ reflect.Type) string {
	skippedJsonTag := []string{"", "-"}
	var exprs []string
	var f []reflect.StructField
	isSliceOfStruct := typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Struct

	if isSliceOfStruct {
		for i := 0; i < typ.Elem().NumField(); i++ {
			f = append(f, typ.Elem().Field(i))
		}
	} else {
		for i := 0; i < typ.NumField(); i++ {
			f = append(f, typ.Field(i))
		}
	}

	for i := 0; i < len(f); i++ {
		jsonTag := strings.TrimSuffix(f[i].Tag.Get("json"), ",omitempty")
		columnTag := f[i].Tag.Get("column")

		// Fast‑skip anything we don’t need to project.
		if ArrayIncludes(skippedJsonTag, jsonTag) || columnTag == "" {
			continue
		}

		exprs = append(exprs, fmt.Sprintf("'%s', %s", jsonTag, columnTag))
	}

	if len(exprs) == 0 {
		return ""
	}

	if isSliceOfStruct {
		return fmt.Sprintf("json_agg(json_build_object(%s))", strings.Join(exprs, ", "))
	}
	return fmt.Sprintf("json_build_object(%s)", strings.Join(exprs, ", "))
}

func ExtractJSONTags[T any]() []string {
	skippedJsonTag := []string{"", "-"}
	var cols []string

	typ := reflect.TypeOf((*T)(nil)).Elem()
	if typ.Kind() != reflect.Struct {
		return cols
	}

	for i := 0; i < typ.NumField(); i++ {
		var col string
		f := typ.Field(i)
		ftype := f.Type

		jsonTag := strings.TrimSuffix(f.Tag.Get("json"), ",omitempty")
		columnTag := f.Tag.Get("column")

		// Fast‑skip anything we don’t need to project.
		if ArrayIncludes(skippedJsonTag, jsonTag) {
			continue
		}

		switch {
		case (ftype.Kind() == reflect.Struct && ftype != reflect.TypeOf(time.Time{})) || (ftype.Kind() == reflect.Slice && ftype.Elem().Kind() == reflect.Struct):
			jsonExpr := traverseNestedStruct[T](ftype)
			if jsonExpr == "" {
				continue
			}
			col = fmt.Sprintf(`%s as "%s"`, jsonExpr, jsonTag)
		// If there's no column tag, then column name will derive from JSON name
		case columnTag == "":
			snake := CamelToSnake(jsonTag)
			if snake == jsonTag {
				col = jsonTag
			} else {
				col = fmt.Sprintf(`"%s" as "%s"`, snake, jsonTag)
			}

		case columnTag == "-":
			// Some columns are generated from json_build_object,etc so it should be skipped here and use Select() instead
			continue

		// Just normal column name that use json as alias
		default:
			col = fmt.Sprintf(`%s as "%s"`, columnTag, jsonTag)
		}

		cols = append(cols, col)
	}

	return cols
}

func MakeMapJSONTags[T any]() map[string]string {
	return MakeMapJSONTagsFromType(reflect.TypeOf((*T)(nil)).Elem())
}

func MakeMapJSONTagsFromValue(dto any) map[string]string {
	typ := reflect.TypeOf(dto)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Slice {
		typ = typ.Elem()
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
	}
	return MakeMapJSONTagsFromType(typ)
}

func MakeMapJSONTagsFromType(typ reflect.Type) map[string]string {
	jsonMap := make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		rawJsonTag := field.Tag.Get("json")
		columnTag := field.Tag.Get("column")
		jsonTag := strings.Split(rawJsonTag, ",")[0]

		if jsonTag == "" || jsonTag == "-" || columnTag == "" {
			continue
		}

		jsonMap[jsonTag] = columnTag
	}
	return jsonMap
}
