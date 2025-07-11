package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

func ExtractJSONTags[T any]() []string {
	var cols []string

	typ := reflect.TypeOf((*T)(nil)).Elem()
	if typ.Kind() != reflect.Struct {
		return cols
	}

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		jsonTag := strings.TrimSuffix(f.Tag.Get("json"), ",omitempty")
		columnTag := f.Tag.Get("column")

		// Fast‑skip anything we don’t need to project.
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		var col string
		switch {
		// If there's no column tag, then column name will derive from JSON name
		case columnTag == "":
			snake := CamelToSnake(jsonTag)
			if snake == jsonTag {
				col = jsonTag
			} else {
				col = fmt.Sprintf(`"%s" as "%s"`, snake, jsonTag)
			}

		// Just normal column name that use json as alias
		default:
			col = fmt.Sprintf(`"%s" as "%s"`, columnTag, jsonTag)
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
