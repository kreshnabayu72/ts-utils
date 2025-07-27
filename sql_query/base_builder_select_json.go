package sql_query

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// SelectJSONArrayElements selects elements from a Go slice of maps
// and expands them as rows using jsonb_array_elements().
//
// Each element in the slice is marshaled to JSON and passed as a query argument.
//
// Example:
//
//	arr := []map[string]string{{"id": "1"}, {"id": "2"}}
//	builder.SelectJSONArrayElements("items", arr)
//
// Generates:
//
//	jsonb_array_elements($1::jsonb) AS items
func (s *SelectBuilder) SelectJSONArrayElements(alias string, arrayElements []map[string]string) SQLSelectChainBuilder {
	if alias == "" {
		alias = "array_elements"
	}

	jsonBytes, err := json.Marshal(arrayElements)
	if err != nil {
		s.LastError = fmt.Errorf("SelectJSONArrayElements marshal error: %w", err)
		return s
	}

	jsonStr := string(jsonBytes)
	placeholder := len(s.Args) + 1
	s.Args = append(s.Args, jsonStr)

	formatted := fmt.Sprintf("jsonb_array_elements(%d::jsonb) AS %s", placeholder, alias)
	s.Columns = append(s.Columns, formatted)

	return s
}

// SelectJSONAggregate builds a JSON object or JSON array aggregation using json_build_object
// or json_agg(json_build_object(...)). It supports optional filtering and ordering.
//
// The `dto` parameter can be either:
//   - a struct (fields are mapped from JSON tags), or
//   - a map[string]string (keys = JSON keys, values = SQL expressions).
//
// If asArrayAggregation is true, the result is wrapped with json_agg().
//
// Example (array aggregation):
//
//	builder.SelectJSONAggregate(
//	    "order_items",
//	    map[string]string{"id": "items.id", "name": "items.name"},
//	    "items.is_active = TRUE",
//	    true,
//	    "items.created_at",
//	)
//
// or with a struct:
//
//	builder.SelectJSONAggregate(
//	    "order_items",
//	    dto.OrderItem{},
//	    "items.is_active = TRUE",
//	    true,
//	    "items.created_at",
//	)
//
// Generates:
//
//	json_agg(json_build_object('id', items.id, 'name', items.name) ORDER BY items.created_at)
//	  FILTER (WHERE items.is_active = TRUE) AS order_items
func (s *SelectBuilder) SelectJSONAggregate(alias string, dto any, condition string, asArrayAggregation bool, orderByClauses ...string) SQLSelectChainBuilder {
	var mappedJSON map[string]string

	v := reflect.ValueOf(dto)
	if v.Kind() != reflect.Struct {
		mappedJSON = dto.(map[string]string)
	} else {
		mappedJSON = MakeMapJSONTagsFromValue(dto)
	}

	if len(mappedJSON) == 0 {
		return s
	}

	var keyValuePairs []string
	var keys []string
	for k := range mappedJSON {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, jsonKey := range keys {
		columnExpr := mappedJSON[jsonKey]
		keyValuePairs = append(keyValuePairs,
			fmt.Sprintf("'%s'", jsonKey),
			columnExpr,
		)
	}

	var formattedColumn string
	if asArrayAggregation {
		orderBy := ""
		if len(orderByClauses) > 0 {
			orderBy = fmt.Sprintf(" ORDER BY %s", orderByClauses[0])
		}
		formattedColumn = fmt.Sprintf("json_agg(json_build_object(%s)%s)", strings.Join(keyValuePairs, ", "), orderBy)
		if condition != "" {
			formattedColumn = fmt.Sprintf("%s FILTER (WHERE %s)", formattedColumn, condition)
		}
	} else {
		formattedColumn = fmt.Sprintf("json_build_object(%s)", strings.Join(keyValuePairs, ", "))
		if condition != "" {
			formattedColumn = fmt.Sprintf("CASE WHEN %s THEN %s ELSE NULL END", condition, formattedColumn)
		}
	}
	formattedColumn = fmt.Sprintf("%s AS %s", formattedColumn, alias)

	if s.WrapAggregation && !asArrayAggregation {
		s.NestedAggregation = append(s.NestedAggregation, formattedColumn)
	} else {
		s.Columns = append(s.Columns, formattedColumn)
	}

	return s
}

// SelectJSONAggregateFunc builds a nested JSON object by executing a callback
// that itself adds JSON aggregate fields. The resulting fields are combined
// into a single json_build_object aliased as `alias`.
//
// Example:
//
//	builder.SelectJSONAggregateFunc("hasSystemRole", func(b *sql_query.SelectBuilder) {
//	    b.SelectJSONAggregate(
//	        string(rt),
//	        map[string]string{
//	            "create": getSystemRolePermission("create"),
//	            "view":   getSystemRolePermission("view"),
//	            "update": getSystemRolePermission("update"),
//	            "delete": getSystemRolePermission("delete"),
//	        },
//	        fmt.Sprintf("usr.role_attribute = $%d", len(builder.Args)+1),
//	        false,
//	    )
//	})
//
// Generates:
//
//	json_build_object(
//	    'SomeRole', CASE WHEN usr.role_attribute = $1 THEN json_build_object(
//	        'create', <expr>, 'view', <expr>, 'update', <expr>, 'delete', <expr>
//	    ) ELSE NULL END
//	) AS hasSystemRole
func (s *SelectBuilder) SelectJSONAggregateFunc(alias string, fn func(builder *SelectBuilder)) SQLSelectChainBuilder {
	if alias == "" {
		alias = "json_result"
	}

	s.WrapAggregation = true
	fn(s)

	s.Columns = append(
		s.Columns,
		fmt.Sprintf("json_build_object(%s) AS %s", strings.Join(s.NestedAggregation, ", "), alias),
	)

	s.WrapAggregation = false
	return s
}
