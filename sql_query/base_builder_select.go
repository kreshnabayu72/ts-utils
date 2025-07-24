package sql_query

import (
	"errors"
	"fmt"
	"strings"
)

// To ensure method .Select() has its own chaining methods
// e.g. .Select(...).Join(...).Where(...).Build()
type SQLSelectChainBuilder interface {
	Select(columns ...string) SQLSelectChainBuilder
	SelectCaseWhen(thenExpr, elseExpr, alias string, whenClause string, whenArgs ...interface{}) SQLSelectChainBuilder
	SelectBoolOr(expr, alias string, args ...interface{}) SQLSelectChainBuilder
	SelectJSONArrayElements(alias string, arrayElements []map[string]string) SQLSelectChainBuilder
	SelectJSONAggregate(alias string, dto any, condition string, asArrayAggregation bool) SQLSelectChainBuilder
	SelectJSONAggregateFunc(alias string, fn func(builder *SelectBuilder), asArray ...bool) SQLSelectChainBuilder

	Where(filters map[string]SQLCondition) SQLSelectChainBuilder
	WhereOr(filters map[string]SQLCondition) SQLSelectChainBuilder

	Search(keyword string, fields []string) SQLSelectChainBuilder
	SetLimit(limit int) SQLSelectChainBuilder
	Join(table string, onCondition string) SQLSelectChainBuilder
	LeftJoin(table string, onCondition string) SQLSelectChainBuilder
	Paginate(query Pagination) SQLSelectChainBuilder
	OrderBy(sortBy []string, asc bool) SQLSelectChainBuilder
	GroupBy(groupBy ...string) SQLSelectChainBuilder
	Having(havingClauses map[string]SQLCondition) SQLSelectChainBuilder

	WithCTEBuilder(cteName string, cteBuilder SQLSelectChainBuilder) SQLSelectChainBuilder

	Build() (string, []interface{}, error)
}

type SelectBuilder struct {
	*SQLEloquentQuery
}

// Where implements SQLSelectChainBuilder. (Accumulates previous value if called again)
func (s *SelectBuilder) Where(filters map[string]SQLCondition) SQLSelectChainBuilder {
	s.SQLEloquentQuery.sharedWhereAndQuery(filters)
	return s
}

// WhereOr implements SQLSelectChainBuilder. (Accumulates previous value if called again)
func (s *SelectBuilder) WhereOr(filters map[string]SQLCondition) SQLSelectChainBuilder {
	s.SQLEloquentQuery.sharedWhereOrQuery(filters)
	return s
}

// Search implements SQLSelectChainBuilder. (Accumulates previous value if called again)
func (s *SelectBuilder) Search(keyword string, fields []string) SQLSelectChainBuilder {
	if keyword != "" && len(fields) > 0 {
		var orClauses []string
		orArgs := []interface{}{}

		for _, field := range fields {
			orClauses = append(orClauses, fmt.Sprintf("%s ILIKE $%d", field, len(s.Args)+len(orArgs)+1))
			orArgs = append(orArgs, "%"+keyword+"%")
		}

		// Combine all clauses with OR inside parentheses
		orClauseGroup := fmt.Sprintf("(%s)", strings.Join(orClauses, " OR "))
		s.Filters = append(s.Filters, orClauseGroup)
		s.Args = append(s.Args, orArgs...)
	}

	return s
}

// Select implements SQLSelectChainBuilder. (Only able to be called once)
func (s *SelectBuilder) Select(columns ...string) SQLSelectChainBuilder {
	for _, newCol := range columns {
		newAlias := extractAlias(newCol)

		// Check if alias exists in current list
		replaced := false
		for i, existing := range s.Columns {
			if extractAlias(existing) == newAlias {
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

func (s *SelectBuilder) SelectBoolOr(expr, alias string, args ...interface{}) SQLSelectChainBuilder {
	s.Columns = append(s.Columns, fmt.Sprintf("bool_or(%s) AS %s", expr, alias))
	s.Args = append(s.Args, args...)
	return s
}

func (s *SelectBuilder) SelectCaseWhen(thenExpr, elseExpr, alias string, whenClause string, whenArgs ...interface{}) SQLSelectChainBuilder {
	s.Columns = append(s.Columns, fmt.Sprintf("CASE WHEN %s THEN %s ELSE %s END AS %s", whenClause, thenExpr, elseExpr, alias))
	s.Args = append(s.Args, whenArgs...)
	return s
}

// JOIN is the equivalent of INNER JOIN
func (s *SelectBuilder) Join(table string, onCondition string) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("JOIN %s ON %s", table, onCondition))
	return s
}

func (s *SelectBuilder) LeftJoin(table string, onCondition string) SQLSelectChainBuilder {
	if table == "" {
		return s
	}

	s.OtherTables = append(s.OtherTables, fmt.Sprintf("LEFT JOIN %s ON %s", table, onCondition))
	return s
}

// GroupBy implements SQLSelectChainBuilder. (Overrides previous value if called again)
func (s *SelectBuilder) GroupBy(groupBy ...string) SQLSelectChainBuilder {
	s.Grouping = append(s.Grouping, groupBy...)
	return s
}

// Having implements SQLSelectChainBuilder. (Overrides previous value if called again)
func (s *SelectBuilder) Having(havingClause map[string]SQLCondition) SQLSelectChainBuilder {
	s.useHaving = true
	s.SQLEloquentQuery.sharedWhereAndQuery(havingClause)
	return s
}

// OrderBy implements SQLSelectChainBuilder. (Accumulates with previous value if called again)
func (s *SelectBuilder) OrderBy(sortBy []string, asc bool) SQLSelectChainBuilder {
	direction := "ASC"

	if !asc {
		direction = "DESC"
	}

	sortingRule := fmt.Sprintf("%s %s", strings.Join(sortBy, ", "), direction)
	s.SortBy = append(s.SortBy, sortingRule)
	return s
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type Pagination struct {
	Page        int    `json:"page"      transform:"int"`
	Limit       int    `json:"limit"     transform:"int"`
	SortBy      string `json:"sortBy"    transform:"string"`
	SortOrder   int    `json:"sortOrder" transform:"int"`
	MultiSort   []Sort `json:"multiSort"`
	DefaultSort []Sort `json:"defaultSort"`
}

// Paginate implements SQLSelectChainBuilder. (Overrides previous value if called again)
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

// SetLimit implements SQLSelectChainBuilder. (Overrides previous value if called again)
func (s *SelectBuilder) SetLimit(limit int) SQLSelectChainBuilder {
	var normalizedPage int

	s.Limit = limit
	s.Offset = normalizedPage * limit

	return s
}

func (s *SelectBuilder) WithCTEBuilder(cteName string, cteBuilder SQLSelectChainBuilder) SQLSelectChainBuilder {
	cteQuery, cteArgs, _ := cteBuilder.Build()

	// Calculate the current argument offset
	offset := len(s.Args)

	// Shift the placeholders in the CTE query
	shiftedCTEQuery := shiftSQLPlaceholders(cteQuery, offset)

	s.WithClauses = append(s.WithClauses, fmt.Sprintf("%s AS (%s)", cteName, shiftedCTEQuery))
	s.Args = append(s.Args, cteArgs...)

	return s
}

func NewSQLSelectBuilder[T any](tableName string) SQLSelectChainBuilder {
	columns := ExtractJSONTags[T]()

	defaultColumns := []string{"*"}
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

func NewSQLCountBuilder(tableName string) SQLSelectChainBuilder {
	defaultColumns := []string{"COUNT(*)"}

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

// buildSelectQuery implements SQLBuilder.\
func (s *SQLEloquentQuery) buildSelectQuery() (string, []interface{}, error) {
	if len(s.Columns) == 0 {
		return "", nil, errors.New("select query must have columns")
	}
	if len(s.HavingClauses) > 0 && len(s.Grouping) < 1 {
		return "", nil, errors.New("HAVING clauses only allowed if GROUP BY clause is exists")
	}

	var dataQuery []string
	var countQuery []string

	if len(s.WithClauses) > 0 {
		dataQuery = append(dataQuery, fmt.Sprintf("WITH %s", strings.Join(s.WithClauses, ", "))+"\n")
		countQuery = append(countQuery, fmt.Sprintf("WITH %s", strings.Join(s.WithClauses, ", "))+"\n")
	}

	dataQuery = append(dataQuery, fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.Columns, ","), s.Table))
	countQuery = dataQuery

	if len(s.OtherTables) > 0 {
		joinClause := strings.Join(s.OtherTables, " ")
		dataQuery = append(dataQuery, joinClause)
		countQuery = append(countQuery, joinClause)
	}

	if len(s.Filters) > 0 {
		whereClause := fmt.Sprintf("WHERE %s", strings.Join(s.Filters, " AND "))
		dataQuery = append(dataQuery, whereClause)
		countQuery = append(countQuery, whereClause)
	}

	if len(s.Grouping) > 0 {
		dataQuery = append(dataQuery, fmt.Sprintf("GROUP BY %s", strings.Join(s.Grouping, ", ")))
		countQuery = append(countQuery, fmt.Sprintf("GROUP BY %s", strings.Join(s.Grouping, ", ")))
	}

	if len(s.HavingClauses) > 0 {
		dataQuery = append(dataQuery, fmt.Sprintf("HAVING %s", strings.Join(s.HavingClauses, " AND ")))
		countQuery = append(countQuery, fmt.Sprintf("HAVING %s", strings.Join(s.HavingClauses, " AND ")))
	}

	if len(s.SortBy) > 0 {
		sortingClause := fmt.Sprintf("ORDER BY %s", strings.Join(s.SortBy, ", "))
		dataQuery = append(dataQuery, sortingClause)
	}

	if s.UsePagination {
		if s.Limit > 0 {
			dataQuery = append(dataQuery, fmt.Sprintf("LIMIT %d", s.Limit))
		}
		dataQuery = append(dataQuery, fmt.Sprintf("OFFSET %d", s.Offset))

		return PaginationQuery(strings.Join(dataQuery, " "), strings.Join(countQuery, " ")), s.Args, nil
	}

	result := strings.Join(dataQuery, " ")
	return result, s.Args, nil
}

// Internal Utils for builder select
func extractAlias(column string) string {
	parts := strings.Split(strings.ToLower(column), " as ")
	if len(parts) < 2 {
		return "" // No alias
	}
	return strings.TrimSpace(parts[len(parts)-1])
}

// func validateGroupByColumns(availableColumns []string, groupByColumns []string) error {
// 	var columnNames []string
// 	for _, column := range availableColumns {
// 		columnNames = append(columnNames, extractAlias(column))
// 	}

// 	for _, each := range groupByColumns {
// 		if !util.ArrayIncludes(columnNames, strings.ToLower(each)) {
// 			return errors.New("GROUP BY columns should exist in selected columns")
// 		}
// 	}

// 	return nil
// }
