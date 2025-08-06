package sql_query

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SQLSelectChainBuilder interface {
	// GetCurrentArgIndex returns the current number of arguments in the query.
	// Useful for calculating placeholder positions when building parameterized SQL.
	GetCurrentArgIndex() int
	// AddArgs appends one or more arguments to the query's argument list.
	// Can be used to manually manage parameter placeholders.
	AddArgs(arg ...interface{}) SQLSelectChainBuilder

	// Distinct implements SQLSelectChainBuilder.
	// Distinct defines one or more columns for the DISTINCT ON(...) statement.
	//
	// Example:
	//
	//	builder.Distinct("u.id", "u.name")
	Distinct(columns ...string) SQLSelectChainBuilder

	// Select implements SQLSelectChainBuilder.
	// Select defines one or more columns for the SELECT statement.
	// If a column alias already exists, it will be replaced.
	//
	// Example:
	//
	//	builder.Select("u.id AS user_id", "u.name")
	Select(columns ...string) SQLSelectChainBuilder
	// SelectCaseWhen adds a CASE WHEN expression as a column.
	//
	// Example:
	//
	//	builder.SelectCaseWhen("'Yes'", "'No'", "is_admin", "role = 'admin'")
	//
	// Generates:
	//
	//	CASE WHEN role = 'admin' THEN 'Yes' ELSE 'No' END AS is_admin
	SelectCaseWhen(thenExpr, elseExpr, alias string, whenClause string, whenArgs ...interface{}) SQLSelectChainBuilder
	// SelectBoolOr adds a bool_or aggregate column with an alias.
	//
	// Example:
	//
	//	builder.SelectBoolOr("is_active", "any_active")
	//
	// Generates:
	//
	//	bool_or(is_active) AS any_active
	SelectBoolOr(expr, alias string, args ...interface{}) SQLSelectChainBuilder

	SelectArrayAggregation(alias string, source string, config ArrayAggConfig) SQLSelectChainBuilder

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
	SelectJSONArrayElements(alias string, arrayElements []map[string]string) SQLSelectChainBuilder
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
	SelectJSONAggregate(alias string, dto any, condition string, asArrayAggregation bool, orderByClauses ...string) SQLSelectChainBuilder
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
	SelectJSONAggregateFunc(alias string, fn func(builder *SelectBuilder)) SQLSelectChainBuilder

	// Where implements SQLSelectChainBuilder. (Accumulates previous value if called again)
	Where(filters map[string]SQLCondition) SQLSelectChainBuilder
	// WhereOr implements SQLSelectChainBuilder. (Accumulates previous value if called again)
	WhereOr(filters ...map[string]SQLCondition) SQLSelectChainBuilder

	// Search implements SQLSelectChainBuilder. (Accumulates previous value if called again).
	// Search adds a case-insensitive ILIKE condition across multiple fields.
	// Each field is compared with the same keyword, combined using OR.
	//
	// Example:
	//
	//	builder.Search("john", []string{"first_name", "last_name"})
	//
	// Generates:
	//
	//	(first_name ILIKE $1 OR last_name ILIKE $1)
	Search(keyword string, fields []string) SQLSelectChainBuilder
	// SetLimit sets a fixed LIMIT value for the query (overwrites any previous limit).
	//
	// Example:
	//
	//	builder.SetLimit(5)
	SetLimit(limit int) SQLSelectChainBuilder
	// Join adds an INNER JOIN clause with the specified ON condition.
	//
	// Example:
	//
	//	builder.Join("roles r", "r.id = u.role_id")
	Join(table string, onCondition string) SQLSelectChainBuilder
	// LeftJoin adds a LEFT JOIN clause with the specified ON condition.
	//
	// Example:
	//
	//	builder.LeftJoin("roles r", "r.id = u.role_id")
	LeftJoin(table string, onCondition string) SQLSelectChainBuilder
	// Paginate implements SQLSelectChainBuilder. (Overrides previous value if called again).
	// Paginate applies LIMIT, OFFSET, and ORDER BY using a Pagination struct.
	// It supports single or multiple sorting rules.
	//
	// Example:
	//
	//	builder.Paginate(Pagination{Page: 2, Limit: 10, SortBy: "name", SortOrder: 1})
	Paginate(query Pagination) SQLSelectChainBuilder
	// OrderBy adds sorting rules to the query. Multiple calls accumulate sorting.
	//
	// Example:
	//
	//	builder.OrderBy([]string{"created_at"}, false) // DESC
	OrderBy(sortBy []string, asc bool) SQLSelectChainBuilder
	// GroupBy adds one or more columns to the GROUP BY clause.
	// Multiple calls accumulate columns.
	//
	// Example:
	//
	//	builder.GroupBy("department", "role")
	GroupBy(groupBy ...string) SQLSelectChainBuilder
	// Having implements SQLSelectChainBuilder. (Overrides previous value if called again).
	// Having adds a HAVING clause for grouped queries.
	// Overwrites any previous HAVING condition.
	//
	// Example:
	//
	//	builder.GroupBy("role").Having(map[string]SQLCondition{
	//	    "count(*)": {Op: ">", Value: 5},
	//	})
	Having(havingClauses map[string]SQLCondition) SQLSelectChainBuilder

	// WithCTEBuilder adds a Common Table Expression (CTE) to the query.
	// It adjusts argument placeholders to avoid conflicts.
	//
	// Example:
	//
	//	cte := NewSQLSelectBuilder[Order]("orders").Select("id", "user_id")
	//	builder.WithCTEBuilder("recent_orders", cte.(*sql_query.SelectBuilder).SQLEloquentQuery)
	//
	// Generates:
	//
	//	WITH recent_orders AS (SELECT id, user_id FROM orders) ...
	WithCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) SQLSelectChainBuilder

	// Build finalizes the SELECT query and returns the query string and arguments.
	// Returns an error if the query is invalid (e.g., HAVING without GROUP BY).
	Build() (string, []interface{}, error)
}

type SelectBuilder struct {
	*SQLEloquentQuery
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

func (s *SelectBuilder) AddArgs(arg ...interface{}) SQLSelectChainBuilder {
	s.Args = append(s.Args, arg...)
	return s
}

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

func (s *SelectBuilder) Distinct(columns ...string) SQLSelectChainBuilder {
	for _, newCol := range columns {
		s.DistinctClauses = append(s.DistinctClauses, newCol)
	}

	return s
}

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
	boolOrColumn := fmt.Sprintf("bool_or(%s) AS \"%s\"", expr, alias)

	// Check if alias exists in current list
	replaced := false
	for i, existing := range s.Columns {
		if extractAlias(existing) == strings.ToLower(alias) {
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

	arrayAggColumn := fmt.Sprintf("(SELECT array_agg(%s%s) FROM %s) AS %s", config.Expr, orderByClause, source, alias)

	replaced := false
	for i, existing := range s.Columns {
		if extractAlias(existing) == strings.ToLower(alias) {
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

func (s *SelectBuilder) SelectCaseWhen(thenExpr, elseExpr, alias string, whenClause string, whenArgs ...interface{}) SQLSelectChainBuilder {
	caseWhenColumn := fmt.Sprintf("CASE WHEN %s THEN %s ELSE %s END AS \"%s\"", whenClause, thenExpr, elseExpr, alias)

	// Check if alias exists in current list
	replaced := false
	for i, existing := range s.Columns {
		if extractAlias(existing) == strings.ToLower(alias) {
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

// NewSQLSelectBuilder creates a new chainable SELECT builder for a given table.
// It extracts JSON tags from the struct type T as default columns.
//
// Example:
//
//	builder := NewSQLSelectBuilder[User]("users").Select("id", "name")
func NewSQLSelectBuilder[T any](tableName string) SQLSelectChainBuilder {
	columns := ExtractJSONTags[T]()

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

func (s *SQLEloquentQuery) buildSelectQuery() (string, []interface{}, error) {
	if len(s.HavingClauses) > 0 && len(s.Grouping) == 0 {
		return "", nil, errors.New("HAVING clauses only allowed if GROUP BY clause is exists")
	}

	if len(s.Columns) == 0 {
		s.Columns = []string{"*"}
	}

	var sb strings.Builder
	var countQuery string
	sb.Grow(256) // preallocate ~256 bytes

	// WITH
	if len(s.WithClauses) > 0 {
		sb.WriteByte('\n')
		sb.WriteString("WITH ")
		for i, c := range s.WithClauses {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(c)
		}
	}

	// DISTINCT
	var db strings.Builder
	if len(s.DistinctClauses) > 0 {
		db.WriteString("DISTINCT ")
		db.WriteString("(")
		for i, d := range s.DistinctClauses {
			if i > 0 {
				db.WriteString(",")
			}
			db.WriteString(d)
		}
		db.WriteString(")")
		var tmp []string

		// Prepend distinct to the column clause.
		copy(tmp, s.Columns)
		s.Columns = []string{db.String()}
		s.Columns = append(s.Columns, tmp...)
	}

	// SELECT
	sb.WriteByte('\n')
	sb.WriteString("SELECT ")
	for i, col := range s.Columns {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(col)
	}
	sb.WriteByte('\n')
	sb.WriteString("FROM ")
	sb.WriteString(s.Table)
	sb.WriteByte('\n')

	// JOIN
	if len(s.OtherTables) > 0 {
		for _, j := range s.OtherTables {
			sb.WriteString(j)
			sb.WriteByte('\n')
		}
	}

	// WHERE
	if len(s.Filters) > 0 {
		sb.WriteString("WHERE ")
		for i, f := range s.Filters {
			if i > 0 {
				sb.WriteString(" AND ")
			}
			sb.WriteString(f)
		}
		sb.WriteByte('\n')
	}

	// GROUP BY
	if len(s.Grouping) > 0 {
		sb.WriteString("GROUP BY ")
		for i, g := range s.Grouping {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(g)
		}
		sb.WriteByte('\n')
	}

	// HAVING
	if len(s.HavingClauses) > 0 {
		sb.WriteString("HAVING ")
		for i, h := range s.HavingClauses {
			if i > 0 {
				sb.WriteString(" AND ")
			}
			sb.WriteString(h)
		}
		sb.WriteByte('\n')
	}

	countQuery = sb.String()

	// ORDER BY
	if len(s.SortBy) > 0 {
		sb.WriteString("ORDER BY ")
		for i, srt := range s.SortBy {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(srt)
		}
		sb.WriteByte('\n')
	}

	// LIMIT/OFFSET
	if s.UsePagination {
		if s.Limit > 0 {
			sb.WriteString("LIMIT ")
			sb.WriteString(strconv.Itoa(s.Limit))
			sb.WriteByte(' ')
		}
		sb.WriteString("OFFSET ")
		sb.WriteString(strconv.Itoa(s.Offset))
		sb.WriteByte('\n')

		return PaginationQuery(sb.String(), countQuery), s.Args, nil
	}

	return sb.String(), s.Args, nil
}

// Internal Utils for builder select
func extractAlias(column string) string {
	parts := strings.Split(strings.ToLower(column), " as ")
	if len(parts) < 2 {
		return "" // No alias
	}
	return strings.Trim(strings.TrimSpace(parts[len(parts)-1]), "\"")
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
