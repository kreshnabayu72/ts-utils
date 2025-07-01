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
			columns = append(columns, fmt.Sprintf("%s as %s", columnTag, jsonTag))
		} else if snakeCaseKey != jsonTag {
			columns = append(columns, fmt.Sprintf("%s as %s", snakeCaseKey, jsonTag))
		} else if jsonTag == "_id" {
			columns = append(columns, fmt.Sprintf("%s as %s", "id", "_id"))
		} else {
			columns = append(columns, jsonTag)
		}
	}

	return columns
}
