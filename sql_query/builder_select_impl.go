package sql_query

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

// NewSQLSelectBuilder creates a new chainable SELECT builder for a given table.
// It extracts JSON tags from the struct type T as default columns.
//
// Example:
//
//	builder := NewSQLSelectBuilder[User]("users","u").Select("id", "name")
func NewSQLSelectBuilder[T any](tableName string, alias ...string) SQLSelectChainBuilder {
	columns := ExtractJSONTags[T]()

	if len(alias) > 0 {
		tableName = fmt.Sprintf("%s %s", tableName, strings.TrimSpace(alias[0]))
	}

	defaultColumns := []string{}
	if len(columns) > 0 {
		defaultColumns = columns
	}
	return &SelectBuilder{
		&SQLEloquentQuery{
			Table:         tableName,
			Filters:       []string{},
			OtherTables:   []string{},
			Columns:       defaultColumns,
			Limit:         0,
			Offset:        0,
			SortBy:        []string{},
			CustomQuery:   "",
			Args:          nil,
			UsePagination: false,
			Mode:          "select",
		},
	}
}

// Example:
//
//	builder := NewSQLCountBuilder("users","u").Where(...)
func NewSQLCountBuilder(tableName string, alias ...string) SQLSelectChainBuilder {
	defaultColumns := []string{"COUNT(*)"}

	if len(alias) > 0 {
		tableName = fmt.Sprintf("%s %s", tableName, strings.TrimSpace(alias[0]))
	}
	return &SelectBuilder{
		&SQLEloquentQuery{
			Table:           tableName,
			Filters:         []string{},
			OtherTables:     []string{},
			UnionAllQueries: []string{},
			Columns:         defaultColumns,
			Limit:           0,
			Offset:          0,
			SortBy:          []string{},
			CustomQuery:     "",
			Args:            nil,
			UsePagination:   false,
			Mode:            "select",
		},
	}
}

func (s *SelectBuilder) Where(filters map[string]SQLCondition) SQLSelectChainBuilder {
	s.SQLEloquentQuery.sharedWhereAndQuery(filters)
	return s
}

func (s *SelectBuilder) WhereOr(filters ...map[string]SQLCondition) SQLSelectChainBuilder {
	s.SQLEloquentQuery.sharedWhereOrQuery(filters...)
	return s
}

func (s *SelectBuilder) GetCurrentArgIndex() int {
	return len(s.Args)
}

func (s *SelectBuilder) AddArgs(arg ...any) SQLSelectChainBuilder {
	s.Args = append(s.Args, arg...)
	return s
}

func (s *SelectBuilder) StartPlaceholderFrom(index int) SQLSelectChainBuilder {
	if index < 1 {
		index = 1
	}

	if len(s.Args) > 0 {
		// Need to pad so that first existing arg becomes $index
		pad := index - 1
		if pad > 0 {
			temp := make([]any, pad+len(s.Args))
			// fill padding with nils
			for i := 0; i < pad; i++ {
				temp[i] = nil
			}
			// copy old args after padding
			copy(temp[pad:], s.Args)
			s.Args = temp
		}
	} else {
		// If empty, just pad with nils up to index-1
		if index > 1 {
			temp := make([]any, index-1)
			for i := 0; i < index-1; i++ {
				temp[i] = nil
			}
			s.Args = temp
		}
	}

	return s
}

func (s *SelectBuilder) Search(keyword string, fields []string) SQLSelectChainBuilder {
	if keyword != "" && len(fields) > 0 {
		var orClauses []string
		orArgs := []any{}

		for _, field := range fields {
			isArrayColumn := strings.HasSuffix(field, ":array")

			if isArrayColumn {
				cleanColumn := strings.TrimSuffix(field, ":array")
				orClauses = append(orClauses, fmt.Sprintf("EXISTS (SELECT 1 FROM unnest(%s) as val WHERE val ILIKE $%d)", cleanColumn, len(s.Args)+len(orArgs)+1))
			} else {
				orClauses = append(orClauses, fmt.Sprintf("%s ILIKE $%d", field, len(s.Args)+len(orArgs)+1))
			}

			orArgs = append(orArgs, "%"+keyword+"%")
		}

		// Combine all clauses with OR inside parentheses
		orClauseGroup := fmt.Sprintf("(%s)", strings.Join(orClauses, " OR "))
		s.Filters = append(s.Filters, orClauseGroup)
		s.Args = append(s.Args, orArgs...)
	}

	return s
}

func (s *SelectBuilder) Distinct(alias string, columns ...string) SQLSelectChainBuilder {
	s.DistinctAlias = alias
	for _, newCol := range columns {
		newAlias := extractAlias(newCol)

		// Check if alias exists in current list
		replaced := false
		for i, existing := range s.DistinctBy {
			extracted := extractAlias(existing)
			if extracted != "" && extracted == newAlias {
				s.DistinctBy[i] = newCol // Overwrite
				replaced = true
				break
			}
		}

		if !replaced {
			s.DistinctBy = append(s.DistinctBy, newCol)
		}
	}

	return s
}

func (s *SelectBuilder) ClearSelects() SQLSelectChainBuilder {
	s.Columns = []string{}
	return s
}

func (s *SelectBuilder) FilterColumnsByAlias(aliases *[]string) SQLSelectChainBuilder {
	if aliases == nil {
		return s
	}

	// make unique set of alias
	aliasSet := make(map[string]struct{}, len(*aliases))
	for _, alias := range *aliases {
		trimmedAlias := trimString(strings.ToLower(alias))

		aliasSet[trimmedAlias] = struct{}{} // struct{}{} uses zero memory
	}

	filteredColumns := make([]string, 0, len(aliasSet))

	for _, existing := range s.Columns {
		// loop through columns
		existingAlias := extractAlias(existing)

		_, foundInSet := aliasSet[existingAlias]

		// _id should always be appended
		if foundInSet || existingAlias == "_id" {
			filteredColumns = append(filteredColumns, existing)
		}
	}

	// overwrite s.Columns to attributes which have alias included in aliases array
	s.Columns = filteredColumns

	return s
}

func (s *SelectBuilder) Select(columns ...string) SQLSelectChainBuilder {
	for _, newCol := range columns {
		newAlias := extractAlias(newCol)

		// Check if alias exists in current list
		replaced := false
		for i, existing := range s.Columns {
			extracted := extractAlias(existing)
			if extracted != "" && extracted == newAlias {
				s.Columns[i] = newCol // Overwrite
				replaced = true
				break
			}
		}

		if !replaced {
			s.Columns = append(s.Columns, newCol)
		}
	}
	return s
}

func (s *SelectBuilder) SelectBoolAnd(expr, alias string, args ...any) SQLSelectChainBuilder {
	boolAndColumn := fmt.Sprintf("bool_and(%s) AS \"%s\"", expr, alias)

	// Check if alias exists in current list
	replaced := false
	for i, existing := range s.Columns {
		extracted := extractAlias(existing)
		if extracted != "" && extracted == strings.ToLower(alias) {
			s.Columns[i] = boolAndColumn // Overwrite
			replaced = true
			break
		}
	}

	if !replaced {
		s.Columns = append(s.Columns, boolAndColumn)
	}

	s.Args = append(s.Args, args...)
	return s
}

func (s *SelectBuilder) SelectBoolOr(expr, alias string, args ...any) SQLSelectChainBuilder {
	boolOrColumn := fmt.Sprintf("bool_or(%s) AS \"%s\"", expr, alias)

	// Check if alias exists in current list
	replaced := false
	for i, existing := range s.Columns {
		extracted := extractAlias(existing)
		if extracted != "" && extracted == strings.ToLower(alias) {
			s.Columns[i] = boolOrColumn // Overwrite
			replaced = true
			break
		}
	}

	if !replaced {
		s.Columns = append(s.Columns, boolOrColumn)
	}

	s.Args = append(s.Args, args...)
	return s
}

func (s *SelectBuilder) SelectArrayAggregation(alias string, source string, config ArrayAggConfig) SQLSelectChainBuilder {
	if config.Expr == "" {
		s.LastError = errors.New("expression should not empty")
	}

	var orderByClause string
	if config.SortBy != "" && config.SortOrder != 0 {
		order := "ASC"
		if config.SortOrder < 0 {
			order = "DESC"
		}
		orderByClause = fmt.Sprintf(" ORDER BY %s %s", config.SortBy, order)
	}

	if source != "" {
		source = " FROM " + source
	}
	arrayAggColumn := fmt.Sprintf(`(SELECT array_agg(%s%s)%s) AS "%s"`, config.Expr, orderByClause, source, alias)

	replaced := false
	for i, existing := range s.Columns {
		extracted := extractAlias(existing)
		if extracted != "" && extracted == strings.ToLower(alias) {
			s.Columns[i] = arrayAggColumn // Overwrite
			replaced = true
			break
		}
	}

	if !replaced {
		s.Columns = append(s.Columns, arrayAggColumn)
	}

	return s
}

func (s *SelectBuilder) SelectCaseWhen(thenExpr, elseExpr, alias string, whenClause string, whenArgs ...any) SQLSelectChainBuilder {
	caseWhenColumn := fmt.Sprintf("CASE WHEN %s THEN %s ELSE %s END AS \"%s\"", whenClause, thenExpr, elseExpr, alias)

	// Check if alias exists in current list
	replaced := false
	for i, existing := range s.Columns {
		extracted := extractAlias(existing)
		if extracted != "" && extracted == strings.ToLower(alias) {
			s.Columns[i] = caseWhenColumn // Overwrite
			replaced = true
			break
		}
	}

	if !replaced {
		s.Columns = append(s.Columns, caseWhenColumn)
	}

	s.Args = append(s.Args, whenArgs...)
	return s
}

func (s *SelectBuilder) SelectMultiCaseWhen(parts []CaseWhenPart, elseExpr, alias string) SQLSelectChainBuilder {
	var caseBuilder strings.Builder
	var allArgs []interface{}

	// Start the CASE statement
	caseBuilder.WriteString("CASE")

	// Loop through all provided parts to build the WHEN...THEN clauses
	for _, part := range parts {
		// Append the clause string
		// Result looks like: CASE WHEN status = 1 THEN 'Active' WHEN status = 2 THEN 'Pending'
		caseBuilder.WriteString(fmt.Sprintf(" WHEN %s THEN %s", part.When, part.Then))

		// Collect arguments in order
		if len(part.Args) > 0 {
			allArgs = append(allArgs, part.Args...)
		}
	}

	// Close with ELSE and END AS
	caseBuilder.WriteString(fmt.Sprintf(" ELSE %s END AS \"%s\"", elseExpr, alias))

	finalColumnStr := caseBuilder.String()

	// --- Existing Logic to Replace or Append Column ---
	replaced := false
	for i, existing := range s.Columns {
		extracted := extractAlias(existing)
		if extracted != "" && extracted == strings.ToLower(alias) {
			s.Columns[i] = finalColumnStr // Overwrite
			replaced = true
			break
		}
	}

	if !replaced {
		s.Columns = append(s.Columns, finalColumnStr)
	}

	// Append all collected arguments to the main builder
	s.Args = append(s.Args, allArgs...)

	return s
}

func (s *SelectBuilder) Join(table string, onCondition string, additionalConditions ...map[string]SQLCondition) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	var filterSb strings.Builder
	if len(additionalConditions) > 0 {
		var filters []string
		s.sharedWhereAndQuery(additionalConditions[0], &filters)

		for i := range filters {
			filterSb.WriteString(" AND ")
			filterSb.WriteString(filters[i])
		}
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("JOIN %s ON %s%s", table, onCondition, filterSb.String()))
	return s
}

func (s *SelectBuilder) LeftJoin(table string, mainCondition string, additionalConditions ...map[string]SQLCondition) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	var filterSb strings.Builder
	if len(additionalConditions) > 0 {
		var filters []string
		s.sharedWhereAndQuery(additionalConditions[0], &filters)

		for i := range filters {
			filterSb.WriteString(" AND ")
			filterSb.WriteString(filters[i])
		}
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("LEFT JOIN %s ON %s%s", table, mainCondition, filterSb.String()))
	return s
}

func (s *SelectBuilder) LeftJoinLateralWithQuery(joinName string, joinQueryBuilder *SQLEloquentQuery, mainCondition string, additionalConditions ...map[string]SQLCondition) SQLSelectChainBuilder {
	joinQuery, joinArgs, err := joinQueryBuilder.Build()
	if err != nil {
		s.LastError = err
		return s
	}

	// Calculate the current argument offset
	offset := len(s.Args)

	// Shift the placeholders in the CTE query
	shiftedCTEQuery := shiftSQLPlaceholders(joinQuery, offset)

	s.Args = append(s.Args, joinArgs...)

	// additional filter
	var filterSb strings.Builder
	if len(additionalConditions) > 0 {
		var filters []string
		s.sharedWhereAndQuery(additionalConditions[0], &filters)

		for i := range filters {
			filterSb.WriteString(" AND ")
			filterSb.WriteString(filters[i])
		}
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("LEFT JOIN LATERAL (%s) %s ON %s%s", shiftedCTEQuery, joinName, mainCondition, filterSb.String()))
	return s
}

func (s *SelectBuilder) GroupBy(groupBy ...string) SQLSelectChainBuilder {
	s.Grouping = append(s.Grouping, groupBy...)
	return s
}

func (s *SelectBuilder) Having(havingClause map[string]SQLCondition) SQLSelectChainBuilder {
	s.useHaving = true
	s.SQLEloquentQuery.sharedWhereAndQuery(havingClause)
	return s
}

func (s *SelectBuilder) OrderBy(sortBy []string, asc bool) SQLSelectChainBuilder {
	direction := "ASC"
	nulls := "FIRST"

	if !asc {
		direction = "DESC"
		nulls = "LAST"
	}

	sortingRule := fmt.Sprintf("%s %s NULLS %s", strings.Join(sortBy, ", "), direction, nulls)
	s.SortBy = append(s.SortBy, sortingRule)
	return s
}

func (s *SelectBuilder) Paginate(query Pagination) SQLSelectChainBuilder {
	var normalizedPage int
	if query.Page > 0 {
		normalizedPage = query.Page - 1
	} else {
		normalizedPage = 0
	}

	s.UsePagination = true
	s.Limit = query.Limit
	s.Offset = normalizedPage * query.Limit

	if query.SortBy != "" && query.SortOrder != 0 {
		s.OrderBy([]string{query.SortBy}, query.SortOrder > 0)
	} else if len(query.MultiSort) > 0 {
		// Overwrite sortBy sortOrder
		// Functions: For nameWithSequence sort (split the sort into 2 sorts)
		s.SortBy = []string{}

		for _, sort := range query.MultiSort {
			s.OrderBy([]string{sort.SortBy}, sort.SortOrder > 0)
		}
	} else if len(query.DefaultSort) > 0 {
		// Overwrite sortBy sortOrder
		// Functions: If all other sorts not defined, define defaultsort
		s.SortBy = []string{}

		for _, sort := range query.DefaultSort {
			s.OrderBy([]string{sort.SortBy}, sort.SortOrder > 0)
		}
	}

	return s
}

func (s *SelectBuilder) SetLimit(limit int) SQLSelectChainBuilder {
	var normalizedPage int

	s.Limit = limit
	s.Offset = normalizedPage * limit

	return s
}

func (s *SelectBuilder) WithCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) SQLSelectChainBuilder {
	cteQuery, cteArgs, err := cteBuilder.Build()
	if err != nil {
		s.LastError = err
		return s
	}

	// Calculate the current argument offset
	offset := len(s.Args)

	// Shift the placeholders in the CTE query
	shiftedCTEQuery := shiftSQLPlaceholders(cteQuery, offset)

	s.WithClauses = append(s.WithClauses, fmt.Sprintf("%s AS (%s)", cteName, shiftedCTEQuery))
	s.Args = append(s.Args, cteArgs...)

	return s
}

func (s *SelectBuilder) WithRecursiveCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) SQLSelectChainBuilder {
	cteQuery, cteArgs, err := cteBuilder.Build()
	if err != nil {
		s.LastError = err
		return s
	}

	// Calculate the current argument offset
	offset := len(s.Args)

	// Shift the placeholders in the CTE query
	shiftedCTEQuery := shiftSQLPlaceholders(cteQuery, offset)

	s.WithClauses = append(s.WithClauses, fmt.Sprintf("%s AS (%s)", cteName, shiftedCTEQuery))
	s.Args = append(s.Args, cteArgs...)

	s.useWithRecursive = true

	return s
}

func (s *SelectBuilder) UnionAll(cteBuilders ...*SQLEloquentQuery) SQLSelectChainBuilder {
	for _, cteBuilder := range cteBuilders {
		s.useUnionAll = true // only set true if len >0
		cteQuery, cteArgs, err := cteBuilder.Build()
		if err != nil {
			s.LastError = err
			return s
		}

		offset := len(s.Args)

		// Shift the placeholders in the CTE query
		shiftedQuery := shiftSQLPlaceholders(cteQuery, offset)

		s.UnionAllQueries = append(s.UnionAllQueries, shiftedQuery)
		s.Args = append(s.Args, cteArgs...)
	}

	return s
}

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

	formatted := fmt.Sprintf(`jsonb_array_elements(%d::jsonb) AS "%s"`, placeholder, alias)
	s.Columns = append(s.Columns, formatted)

	return s
}

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
		formattedColumn = fmt.Sprintf("jsonb_agg(jsonb_build_object(%s)%s)", strings.Join(keyValuePairs, ", "), orderBy)
		if condition != "" {
			formattedColumn = fmt.Sprintf("%s FILTER (WHERE %s)", formattedColumn, condition)
		}
	} else {
		formattedColumn = fmt.Sprintf("jsonb_build_object(%s)", strings.Join(keyValuePairs, ", "))
		if condition != "" {
			formattedColumn = fmt.Sprintf("CASE WHEN %s THEN %s ELSE NULL END", condition, formattedColumn)
		}
	}
	formattedColumn = fmt.Sprintf(`%s AS "%s"`, formattedColumn, alias)

	if s.WrapAggregation && !asArrayAggregation {
		s.NestedAggregation = append(s.NestedAggregation, formattedColumn)
	} else {
		s.Columns = append(s.Columns, formattedColumn)
	}

	return s
}

func (s *SelectBuilder) SelectJSONAggregateCoalesce(alias string, dto any, condition string, asArrayAggregation bool, coalesce string, orderByClauses ...string) SQLSelectChainBuilder {
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
		formattedColumn = fmt.Sprintf("jsonb_agg(jsonb_build_object(%s)%s)", strings.Join(keyValuePairs, ", "), orderBy)

		if condition != "" {
			formattedColumn = fmt.Sprintf("%s FILTER (WHERE %s)", formattedColumn, condition)
		}
		formattedColumn = fmt.Sprintf("COALESCE(%s,%s)", formattedColumn, coalesce)

	} else {
		formattedColumn = fmt.Sprintf("jsonb_build_object(%s)", strings.Join(keyValuePairs, ", "))

		if condition != "" {
			formattedColumn = fmt.Sprintf("CASE WHEN %s THEN %s ELSE NULL END", condition, formattedColumn)
		}
		formattedColumn = fmt.Sprintf("COALESCE(%s,%s)", formattedColumn, coalesce)

	}
	formattedColumn = fmt.Sprintf(`%s AS "%s"`, formattedColumn, alias)

	if s.WrapAggregation && !asArrayAggregation {
		s.NestedAggregation = append(s.NestedAggregation, formattedColumn)
	} else {
		s.Columns = append(s.Columns, formattedColumn)
	}

	return s
}

func (s *SelectBuilder) SelectJSONAggregateDistinct(alias string, dto any, condition string, asArrayAggregation bool, orderByClauses ...string) SQLSelectChainBuilder {
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
		formattedColumn = fmt.Sprintf("jsonb_agg(DISTINCT jsonb_build_object(%s)%s)", strings.Join(keyValuePairs, ", "), orderBy)
		if condition != "" {
			formattedColumn = fmt.Sprintf("%s FILTER (WHERE %s)", formattedColumn, condition)
		}
	} else {
		formattedColumn = fmt.Sprintf("jsonb_build_object(%s)", strings.Join(keyValuePairs, ", "))
		if condition != "" {
			formattedColumn = fmt.Sprintf("CASE WHEN %s THEN %s ELSE NULL END", condition, formattedColumn)
		}
	}
	formattedColumn = fmt.Sprintf(`%s AS "%s"`, formattedColumn, alias)

	if s.WrapAggregation && !asArrayAggregation {
		s.NestedAggregation = append(s.NestedAggregation, formattedColumn)
	} else {
		s.Columns = append(s.Columns, formattedColumn)
	}

	return s
}

func (s *SelectBuilder) SelectJSONAggregateFunc(alias string, fn func(builder *SelectBuilder)) SQLSelectChainBuilder {
	if alias == "" {
		alias = "json_result"
	}

	s.WrapAggregation = true
	fn(s)

	s.Columns = append(
		s.Columns,
		fmt.Sprintf(`jsonb_build_object(%s) AS "%s"`, strings.Join(s.NestedAggregation, ", "), alias),
	)

	s.WrapAggregation = false
	return s
}

func (s *SQLEloquentQuery) buildSelectQuery() (string, []any, error) {
	if len(s.HavingClauses) > 0 && len(s.Grouping) == 0 {
		return "", nil, errors.New("HAVING clauses only allowed if GROUP BY clause is exists")
	}

	if len(s.Columns) == 0 {
		s.Columns = []string{"*"}
	}

	var withSb strings.Builder
	var selectSb strings.Builder
	var joinSb strings.Builder
	var whereSb strings.Builder
	var groupSb strings.Builder
	var orderSb strings.Builder
	var havingSb strings.Builder
	withSb.Grow(256)   // preallocate ~256 bytes
	selectSb.Grow(256) // preallocate ~256 bytes
	joinSb.Grow(256)   // preallocate ~256 bytes
	whereSb.Grow(256)  // preallocate ~256 bytes
	groupSb.Grow(256)  // preallocate ~256 bytes
	orderSb.Grow(256)  // preallocate ~256 bytes
	havingSb.Grow(256) // preallocate ~256 bytes

	// WITH
	if len(s.WithClauses) > 0 {
		withSb.WriteByte('\n')

		// if one of WITH are recursive, add 'RECURSIVE' string
		if s.useWithRecursive {
			withSb.WriteString("WITH RECURSIVE ")
		} else {
			withSb.WriteString("WITH ")
		}

		for i, c := range s.WithClauses {
			if i > 0 {
				withSb.WriteString(", ")
			}
			withSb.WriteString(c)
		}
	}

	// DISTINCT
	var db strings.Builder
	if len(s.DistinctBy) > 0 {
		db.WriteString("DISTINCT ON ")
		db.WriteString("(")
		for i, d := range s.DistinctBy {
			if i > 0 {
				db.WriteString(",")
			}
			db.WriteString(d)
		}
		db.WriteString(")")
		db.WriteString(" " + s.DistinctAlias)

		tmp := make([]string, len(s.Columns))

		// Prepend distinct to the column clause.
		copy(tmp, s.Columns)
		s.Columns = []string{db.String()}
		s.Columns = append(s.Columns, tmp...)
	}

	// SELECT
	// Dont need to use default select if using union All, since union typically uses SELECT inside of it
	if !s.useUnionAll {
		selectSb.WriteByte('\n')
		selectSb.WriteString("SELECT ")
		for i, col := range s.Columns {
			if i > 0 {
				selectSb.WriteByte(',')
			}
			selectSb.WriteString(col)
		}
		selectSb.WriteByte('\n')
		selectSb.WriteString("FROM ")
		selectSb.WriteString(s.Table)
		selectSb.WriteByte('\n')
	} else if len(s.UnionAllQueries) > 0 { // UNION ALL
		for i, u := range s.UnionAllQueries {
			if i > 0 {
				selectSb.WriteString("UNION ALL")
				selectSb.WriteByte('\n')
			}
			selectSb.WriteString(u)
		}
		selectSb.WriteByte('\n')
	}

	// JOIN
	if len(s.OtherTables) > 0 {
		for _, j := range s.OtherTables {
			joinSb.WriteString(j)
			joinSb.WriteByte('\n')
		}
	}

	// WHERE
	if len(s.Filters) > 0 {
		whereSb.WriteString("WHERE ")
		for i, f := range s.Filters {
			if i > 0 {
				whereSb.WriteString(" AND ")
			}
			whereSb.WriteString(f)
		}
		whereSb.WriteByte('\n')
	}

	// GROUP BY
	if len(s.Grouping) > 0 {
		groupSb.WriteString("GROUP BY ")
		for i, g := range s.Grouping {
			if i > 0 {
				groupSb.WriteString(", ")
			}
			groupSb.WriteString(g)
		}
		groupSb.WriteByte('\n')
	}

	// HAVING
	if len(s.HavingClauses) > 0 {
		havingSb.WriteString("HAVING ")
		for i, h := range s.HavingClauses {
			if i > 0 {
				havingSb.WriteString(" AND ")
			}
			havingSb.WriteString(h)
		}
		havingSb.WriteByte('\n')
	}

	// ORDER BY
	if len(s.SortBy) > 0 {
		orderSb.WriteString("ORDER BY ")

		asSplit := regexp.MustCompile(`(?i)\s+AS\s+`)
		stripDirAndNull := regexp.MustCompile(`\s+(?i)((?:asc|desc)(?:\s+NULLS (?:FIRST|LAST))?|(?:NULLS (?:FIRST|LAST)))\s*$`)

		// helper: trim space, trailing comma, and surrounding double quotes
		clean := func(s string) string {
			s = strings.TrimSpace(s)
			s = strings.TrimSuffix(s, ",")
			if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
				s = s[1 : len(s)-1]
			}
			return s
		}

		// map for transforming alias ("dataType") to column ("cf.data_type")
		aliasToExpr := make(map[string]string, len(s.Columns))
		for _, col := range s.Columns {
			col = strings.TrimSpace(col)
			parts := asSplit.Split(col, 2)

			if len(parts) == 2 {
				expr := strings.TrimSpace(parts[0]) // e.g. cf.data_type
				alias := clean(parts[1])            // e.g. "dataType" -> dataType
				key := strings.ToLower(alias)
				aliasToExpr[key] = expr
			} else {
				// support bare columns too (no AS)
				ident := clean(col)
				key := strings.ToLower(ident)
				aliasToExpr[key] = ident
			}
		}

		for i, srt := range s.SortBy {
			if i > 0 {
				orderSb.WriteString(", ")
			}

			srt = strings.TrimSpace(srt)

			// get dir (asc/desc) if present (keep original casing)
			dir := getSortDirection(srt)

			// key without ASC/DESC, then unquote/clean
			key := strings.TrimSpace(stripDirAndNull.ReplaceAllString(srt, ""))
			key = clean(key)
			lookup := strings.ToLower(key)                                       // "updatedAt" asc null last
			lookup = strings.Split(strings.ReplaceAll(lookup, "\"", ""), " ")[0] // updatedAt

			expr, ok := aliasToExpr[lookup]
			// resolve alias -> expression (fallback to key as-is)
			if ok {
				orderSb.WriteString(expr + dir)
			} else {
				orderSb.WriteString(key + dir)
			}
		}
		orderSb.WriteByte('\n')
	}

	// LIMIT/OFFSET
	if s.UsePagination {
		var limitationSb strings.Builder
		if s.Limit > 0 {
			limitationSb.WriteString("LIMIT ")
			limitationSb.WriteString(strconv.Itoa(s.Limit))
			limitationSb.WriteByte(' ')
		}
		limitationSb.WriteString("OFFSET ")
		limitationSb.WriteString(strconv.Itoa(s.Offset))
		limitationSb.WriteByte('\n')

		splittedTableName := strings.Split(s.Table, " ")
		prefix := splittedTableName[0]
		if len(splittedTableName) > 1 {
			prefix = splittedTableName[1]
		}

		mainQuery := selectSb.String() + joinSb.String() + fmt.Sprintf("JOIN paginated_ids ON paginated_ids.id = %s.id\n", prefix) + groupSb.String() + havingSb.String() + orderSb.String()
		filteredData := fmt.Sprintf("SELECT %s.id as id from %s\n", prefix, s.Table) + joinSb.String() + whereSb.String() + groupSb.String() + havingSb.String() + orderSb.String()
		paginatedDataQuery := "SELECT id as id from filtered_ids\n" + limitationSb.String()
		paginatedCountQuery := "SELECT COUNT(id) from filtered_ids\n"
		return PaginationQuery(withSb.String(), mainQuery, filteredData, paginatedDataQuery, paginatedCountQuery), s.Args, nil
	}

	query := withSb.String() + selectSb.String() + joinSb.String() + whereSb.String() + groupSb.String() + havingSb.String() + orderSb.String()
	return query, s.Args, nil
}

// Internal Utils for builder select
func extractAlias(column string) string {
	parts := strings.Split(strings.ToLower(column), " as ")
	if len(parts) < 2 {
		return "" // No alias
	}

	alias := parts[len(parts)-1]

	return trimString(alias)
}

func trimString(s string) string {
	return strings.Trim(strings.TrimSpace(s), "\"")
}

func getSortDirection(srt string) string {
	fields := strings.Fields(srt)
	dir := ""

	// Loop through all fields to find the first instance
	// of "asc" or "desc", ignoring case.
	for _, field := range fields {
		if strings.EqualFold(field, "asc") {
			dir = field // Keep the original casing (e.g., "asc", "ASC")
			break       // Found it, stop searching
		}
		if strings.EqualFold(field, "desc") {
			dir = field // Keep the original casing (e.g., "desc", "DESC")
			break       // Found it, stop searching
		}
	}

	// If we found a direction, add the space prefix
	if dir != "" {
		return " " + dir
	}

	return ""
}
