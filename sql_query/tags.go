package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

func ExtractJSONTags[T any]() []string {
	var columns []string
	typ := reflect.TypeOf((*T)(nil)).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		rawJsonTag := field.Tag.Get("json")
		columnTag := field.Tag.Get("column")
		jsonTag := strings.ReplaceAll(rawJsonTag, ",omitempty", "")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		snakeCaseKey := CamelToSnake(jsonTag)

		if columnTag != "" {
			columns = append(columns, fmt.Sprintf(`"%s" as "%s"`, columnTag, jsonTag))
		} else if snakeCaseKey != jsonTag {
			columns = append(columns, fmt.Sprintf(`"%s" as "%s"`, snakeCaseKey, jsonTag))
		} else if jsonTag == "_id" {
			columns = append(columns, `"id" as "_id"`)
		} else {
			columns = append(columns, fmt.Sprintf(`"%s"`, jsonTag))
		}
	}

	return columns
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
