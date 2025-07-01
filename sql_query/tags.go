package sql_query

import (
	"fmt"
	"reflect"
	"strings"
)

func ExtractJSONTags[T any](tableName string) []string {
	var columns []string
	typ := reflect.TypeOf((*T)(nil)).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		rawJsonTag := field.Tag.Get("json")
		jsonTag := strings.ReplaceAll(rawJsonTag, ",omitempty", "")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		snakeCaseKey := CamelToSnake(jsonTag)

		if snakeCaseKey != jsonTag {
			columns = append(columns, fmt.Sprintf("%s.%s as %s", tableName, snakeCaseKey, jsonTag))
		} else if jsonTag == "_id" {
			columns = append(columns, fmt.Sprintf("%s.%s as %s", tableName, "id", "_id"))
		} else {
			columns = append(columns, jsonTag)
		}
	}

	return columns
}
