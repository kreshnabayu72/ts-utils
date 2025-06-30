package sql_query

import (
	"fmt"
	"math"
	"strings"
)

type SQLBuilder struct {
	Table                   string
	SelectCols              []string
	WhereParts              []string
	Args                    []interface{}
	LimitVal                int
	OffsetVal               int
	OrderBy                 string
	argIndex                int
	usePagination           bool
	setTimeStampToCamelCase bool
}

type SQLFilter map[string]SQLCondition

type SQLCondition struct {
	Operator SQLOperators // e.g., '=', '>', '<=', 'LIKE', 'IN', 'IS NULL'
	Value    interface{}  // could be a single value, slice, or nil
	IsRef bool // to determine whether WHERE is targeting literal value or reference
}

// New creates a new SQLBuilder for a given table
func NewSQLBuilder(table string) *SQLBuilder {
	return &SQLBuilder{
		Table:                   table,
		SelectCols:              []string{"*"},
		WhereParts:              []string{},
		Args:                    []interface{}{},
		argIndex:                1,
		setTimeStampToCamelCase: true,
	}
}

// Select sets the columns to select
func (b *SQLBuilder) Select(columns []string) *SQLBuilder {
	if len(columns) > 0 {
		b.SelectCols = columns
	}
	return b
}

/*
Where applies filters from SQLFilter into SQL WHERE conditions.

Example filter:

	filter := SQLFilter{
		"status": {Operator: "=", Value: "active"},
		"age":    {Operator: ">=", Value: 18},
		"tags":   {Operator: "IN", Value: []string{"golang", "dev"}},
	}

Operator Reference (Mongo-like to SQL):

	"$eq"   → "="
	"$ne"   → "!="
	"$gt"   → ">"
	"$lt"   → "<"
	"$gte"  → ">="
	"$lte"  → "<="
	"$in"   → "IN"
	"$nin"  → "NOT IN"
*/
func (b *SQLBuilder) Where(filter SQLFilter) *SQLBuilder {
	for key, condition := range filter {
		operator := condition.Operator
		value := condition.Value

		// Optional: handle IN/NOT IN for slices
		if operator == "IN" || operator == "NOT IN" {
			placeholders := b.makePlaceholders(value)
			b.WhereParts = append(b.WhereParts, fmt.Sprintf(`"%s" %s (%s)`, key, operator, placeholders))
			b.Args = append(b.Args, b.flattenArgs(value)...)
			b.argIndex += b.countArgs(value)
		} else {
			b.WhereParts = append(b.WhereParts, fmt.Sprintf(`"%s" %s $%d`, key, operator, b.argIndex))
			b.Args = append(b.Args, value)
			b.argIndex++
		}
	}
	return b
}

// Order adds an ORDER BY clause
func (b *SQLBuilder) Order(field string, asc bool) *SQLBuilder {
	order := "ASC"
	if !asc {
		order = "DESC"
	}
	b.OrderBy = fmt.Sprintf(`ORDER BY "%s" %s`, field, order)
	return b
}

// Sets limit & offset (skip)
func (b *SQLBuilder) Pagination(page int, limit int) *SQLBuilder {
	b.usePagination = true
	b.LimitVal = math.MaxInt
	b.OffsetVal = 0

	if limit > 0 {
		b.LimitVal = limit
	}
	if page > 0 && limit > 0 {
		b.OffsetVal = (page - 1) * limit
	}

	return b
}

// Set simulates MongoDB's $set by adding computed SELECT fields
func (b *SQLBuilder) Set(aliases map[string]string) *SQLBuilder {
	for alias, expression := range aliases {
		// expression: raw SQL like "first_name || ' ' || last_name"
		b.SelectCols = append(b.SelectCols, fmt.Sprintf("%s AS %s", expression, alias))
	}
	return b
}

func (b *SQLBuilder) Search(keyword string, fields []string) *SQLBuilder {
	if keyword != "" && len(fields) > 0 {
		var orParts []string
		for _, field := range fields {
			orParts = append(orParts, fmt.Sprintf(`"%s" ILIKE $%d`, field, b.argIndex))
			b.Args = append(b.Args, "%"+keyword+"%")
			b.argIndex++
		}
		// Group with parentheses
		b.WhereParts = append(b.WhereParts, "("+strings.Join(orParts, " OR ")+")")
	}
	return b
}

func (b *SQLBuilder) GetPaginationQuery(selectCols string, filter string) string {
	// Subquery for paginated result
	paginatedQuery := fmt.Sprintf(`
			SELECT %s FROM %s
			%s
			%s
			LIMIT %d OFFSET %d
		`,
		selectCols,
		b.Table,
		filter,
		b.OrderBy,
		b.LimitVal,
		b.OffsetVal,
	)

	// Subquery for total count
	countQuery := fmt.Sprintf(`
			SELECT COUNT(*) FROM %s %s
		`, b.Table, filter)

	// Final combined query
	finalQuery := fmt.Sprintf(`
		WITH 
			paginated AS (%s),
			total_count AS (%s)
		SELECT 
			COALESCE((SELECT json_agg(paginated) FROM paginated), '[]') AS data,
			(SELECT COUNT FROM total_count) AS totalRecords;
	`, paginatedQuery, countQuery)
	return finalQuery
}

// Build returns the final SQL query string (with facet if paginated) and args
func (b *SQLBuilder) Build() (string, []interface{}) {
	filter := ""
	if len(b.WhereParts) > 0 {
		filter = "WHERE " + strings.Join(b.WhereParts, " AND ")
	}

	selectCols := strings.Join(b.SelectCols, ", ")
	if selectCols == "" {
		selectCols = "*"
	}

	// If paginated, return facet-style query with total count
	if b.usePagination {
		return b.GetPaginationQuery(selectCols, filter), b.Args
	}

	// If not paginated, return simple SELECT
	query := fmt.Sprintf("SELECT %s FROM %s", selectCols, b.Table)

	if filter != "" {
		query += " " + filter
	}

	if b.OrderBy != "" {
		query += " " + b.OrderBy
	}

	return query, b.Args
}
