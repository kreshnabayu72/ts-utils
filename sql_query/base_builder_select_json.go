package sql_query

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// SelectJSONArrayElements implements SQLSelectChainBuilder
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

// SelectJSONAggregate implements SQLSelectChainBuilder
func (s *SelectBuilder) SelectJSONAggregate(alias string, dto any, aggregateAsArray bool) SQLSelectChainBuilder {
	s.useAggregation = true

	if alias == "" {
		alias = "json_result"
	}

	var mappedJSON map[string]string

	v := reflect.ValueOf(dto)
	if v.Kind() != reflect.Struct {
		mappedJSON = dto.(map[string]string)
	} else {
		mappedJSON = MakeMapJSONTagsFromValue(dto)
	}
	fmt.Println("dto", dto)
	fmt.Println("mappedJSON", mappedJSON)

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

	formattedColumn := fmt.Sprintf("json_agg(json_build_object(%s))", strings.Join(keyValuePairs, ", "))
	if !aggregateAsArray {
		formattedColumn = fmt.Sprintf("json_build_object(%s)", strings.Join(keyValuePairs, ", "))
	}

	s.AggregationAlias = alias
	s.AggregationFields = formattedColumn
	return s
}

// SelectJSONAggregate implements SQLSelectChainBuilder
func (s *SelectBuilder) SelectJSONAggregateFunc(fn func(builder *SelectBuilder)) SQLSelectChainBuilder {
	fn(s)
	var formattedColumn string

	if len(s.AggregationFilters) > 0 {
		formattedColumn = fmt.Sprintf("%s FILTER (WHERE %s)", s.AggregationFields, strings.Join(s.AggregationFilters, " AND "))
	} else {
		formattedColumn = s.AggregationFields
	}

	s.Columns = append(
		s.Columns,
		fmt.Sprintf("%s AS %s", formattedColumn, s.AggregationAlias),
	)

	s.AggregationFields = ""
	s.AggregationAlias = ""
	s.AggregationFilters = []string{}
	s.useAggregation = false
	return s
}
