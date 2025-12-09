package sql_query

type Selector interface {
	// GetCurrentArgIndex returns the current number of arguments in the query.
	// Useful for calculating placeholder positions when building parameterized SQL.
	GetCurrentArgIndex() int
	// AddArgs appends one or more arguments to the query's argument list.
	// Can be used to manually manage parameter placeholders.
	AddArgs(arg ...any) Selector
	// StartPlaceholderFrom set starting point of next generated placeholders.
	// Useful for creating sub query from this builder for another main query.
	// Placeholders for this sub query will be started from this given value.
	StartPlaceholderFrom(index int) Selector

	// Distinct implements Selector.
	// Distinct defines one or more columns for the DISTINCT ON(...) statement.
	//
	// Example:
	//
	//	builder.Distinct("distinct", "u.id", "u.name")
	Distinct(alias string, columns ...string) Selector

	// Select implements Selector.
	// Select defines one or more columns for the SELECT statement.
	// If a column alias already exists, its expression will be overwritten with the new one.
	//
	// Example:
	//
	//	builder.Select("u.id AS user_id", "u.name")
	Select(columns ...string) Selector

	// USE WITH CAUTION
	// Reset all previous appended selects
	// Useful for this part in GenerateCTEOption
	// cteBuilder := sourceBuilder.
	//		ClearSelects(). -- without this the previous selects inside the source builder will get included
	// 		Distinct(
	// 			fmt.Sprintf(`%s AS "value"`, labelValue),
	// 			labelValue,
	// 		).
	// 		Select(
	// 			fmt.Sprintf(`%s AS "key"`, labelKey),
	// 			fmt.Sprintf(`"%s"."id" AS "id"`, refTable),
	// 		)
	//
	// Example usage:
	//
	//	builder.ClearSelect()
	ClearSelects() Selector

	// FilterColumnsByAlias filters out elements in select array using its aliases.
	// Will be executed after all select functions
	// Used for handling query.Columns for filtering response attribute from client-side
	//
	// Example:
	//	 s.Columns := []string{"id as _id", "asset_id as assetId", "groups"}
	//
	//	builder.FilterColumnsByAlias([]string{"_id","assetId"}) //removed groups from s.Columns
	FilterColumnsByAlias(aliases *[]string) Selector

	// SelectCaseWhen adds a CASE WHEN expression as a column.
	//
	// Example:
	//
	//	builder.SelectCaseWhen("'Yes'", "'No'", "is_admin", "role = 'admin'")
	//
	// Generates:
	//
	//	CASE WHEN role = 'admin' THEN 'Yes' ELSE 'No' END AS is_admin
	SelectCaseWhen(thenExpr, elseExpr, alias string, whenClause string, whenArgs ...interface{}) Selector
	// SelectMultiCaseWhen adds a complex CASE statement with multiple WHEN conditions as a column.
	// It iterates through the provided parts to build the logic and appends the elseExpr as the fallback.
	//
	// Example:
	//
	// 	parts := []CaseWhenPart{
	// 		{When: "quantity > ?", Then: "'High'", Args: []interface{}{100}},
	// 		{When: "quantity > ?", Then: "'Medium'", Args: []interface{}{50}},
	// 	}
	// 	builder.SelectMultiCaseWhen(parts, "'Low'", "stock_level")
	//
	// Generates:
	//
	// 	CASE WHEN quantity > ? THEN 'High' WHEN quantity > ? THEN 'Medium' ELSE 'Low' END AS "stock_level"
	SelectMultiCaseWhen(parts []CaseWhenPart, elseExpr, alias string) Selector
	// SelectBoolAnd adds a bool_or aggregate column with an alias.
	//
	// Example:
	//
	//	builder.SelectBoolAnd("is_active", "any_active")
	//
	// Generates:
	//
	//	bool_and(is_active) AS any_active
	SelectBoolAnd(expr, alias string, args ...any) Selector
	// SelectBoolOr adds a bool_or aggregate column with an alias.
	//
	// Example:
	//
	//	builder.SelectBoolOr("is_active", "any_active")
	//
	// Generates:
	//
	//	bool_or(is_active) AS any_active
	SelectBoolOr(expr, alias string, args ...any) Selector

	SelectArrayAggregation(alias string, source string, config ArrayAggConfig) Selector

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
	SelectJSONArrayElements(alias string, arrayElements []map[string]string) Selector
	// SelectJSONAggregate builds a JSON object or JSON array aggregation using jsonb_build_object
	// or jsonb_agg(jsonb_build_object(...)). It supports optional filtering and ordering.
	//
	// The `dto` parameter can be either:
	//   - a struct (fields are mapped from JSON tags), or
	//   - a map[string]string (keys = JSON keys, values = SQL expressions).
	//
	// If asArrayAggregation is true, the result is wrapped with jsonb_agg().
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
	//	jsonb_agg(jsonb_build_object('id', items.id, 'name', items.name) ORDER BY items.created_at)
	//	  FILTER (WHERE items.is_active = TRUE) AS order_items
	SelectJSONAggregate(alias string, dto any, condition string, asArrayAggregation bool, orderByClauses ...string) Selector
	// Read documentation for SelectJSONAggregate since the function is similar but with additional COALESCE
	// Generates:
	//
	//	COALESCE(jsonb_agg(DISTINCT jsonb_build_object('id', items.id, 'name', items.name) ORDER BY items.created_at)
	//	  FILTER (WHERE items.is_active = TRUE) AS order_items, ${coalesce})
	SelectJSONAggregateCoalesce(alias string, dto any, condition string, asArrayAggregation bool, coalesce string, orderByClauses ...string) Selector
	// Read documentation for SelectJSONAggregate since the function is similar but with additional DISTINCT
	// Generates:
	//
	//	jsonb_agg(DISTINCT jsonb_build_object('id', items.id, 'name', items.name) ORDER BY items.created_at)
	//	  FILTER (WHERE items.is_active = TRUE) AS order_items
	SelectJSONAggregateDistinct(alias string, dto any, condition string, asArrayAggregation bool, orderByClauses ...string) Selector
	// SelectJSONAggregateFunc builds a nested JSON object by executing a callback
	// that itself adds JSON aggregate fields. The resulting fields are combined
	// into a single jsonb_build_object aliased as `alias`.
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
	//	jsonb_build_object(
	//	    'SomeRole', CASE WHEN usr.role_attribute = $1 THEN jsonb_build_object(
	//	        'create', <expr>, 'view', <expr>, 'update', <expr>, 'delete', <expr>
	//	    ) ELSE NULL END
	//	) AS hasSystemRole
	SelectJSONAggregateFunc(alias string, fn func(builder *SelectBuilder)) Selector

	// Where implements Selector. (Accumulates previous value if called again)
	Where(filters WhereQuery) Selector
	// WhereOr implements Selector. (Accumulates previous value if called again)
	WhereOr(filters ...WhereQuery) Selector

	// Search implements Selector and accumulates conditions if called multiple times.
	// Adds a case-insensitive ILIKE condition across multiple fields, combined with OR.
	//
	// To search inside array columns, add suffix ":array" to the column name.
	// This uses EXISTS + unnest for better performance & usabilitiy than array_to_string.
	//
	// Example:
	//
	//	builder.Search("john", []string{"status::text", "last_name", "tags:array"})
	//
	// Generates:
	//
	//	(first_name ILIKE $1 OR last_name ILIKE $1 OR EXISTS (SELECT 1 FROM unnest(tags) AS val WHERE val ILIKE $1))
	Search(keyword string, fields []string) Selector
	// SetLimit sets a fixed LIMIT value for the query (overwrites any previous limit).
	//
	// Example:
	//
	//	builder.SetLimit(5)
	SetLimit(limit int) Selector
	// Join adds an INNER JOIN clause with the specified ON condition.
	//
	// Example:
	//
	//	builder.Join("roles r", "r.id = u.role_id")
	Join(table string, onCondition string, additionalConditions ...WhereQuery) Selector
	// LeftJoin adds a LEFT JOIN clause with the specified ON condition.
	//
	// Example:
	//
	//	builder.LeftJoin("roles r", "r.id = u.role_id")
	LeftJoin(table string, onCondition string, additionalConditions ...WhereQuery) Selector

	// Example:
	//
	// LeftJoinLateralWithQuery("ta", categoryWithRecursiveBuilder().(*sql_query.SelectBuilder).SQLEloquentQuery, "TRUE").
	//
	// Output:
	//
	// 	LEFT JOIN LATERAL (
	//     `Sql queries here`
	//   ) `alias` ON `condition`
	LeftJoinLateralWithQuery(joinName string, joinQueryBuilder *SQLEloquentQuery, mainCondition string, additionalConditions ...WhereQuery) Selector

	// Paginate implements Selector. (Overrides previous value if called again).
	// Paginate applies LIMIT, OFFSET, and ORDER BY using a Pagination struct.
	// It supports single or multiple sorting rules.
	//
	// Example:
	//
	//	builder.Paginate(Pagination{Page: 2, Limit: 10, SortBy: "name", SortOrder: 1})
	Paginate(query Pagination) Selector
	// OrderBy adds sorting rules to the query. Multiple calls accumulate sorting.
	//
	// Example:
	//
	//	builder.OrderBy([]string{"created_at"}, false) // DESC
	OrderBy(sortBy []string, asc bool) Selector
	// GroupBy adds one or more columns to the GROUP BY clause.
	// Multiple calls accumulate columns.
	//
	// Example:
	//
	//	builder.GroupBy("department", "role")
	GroupBy(groupBy ...string) Selector
	// Having implements Selector. (Overrides previous value if called again).
	// Having adds a HAVING clause for grouped queries.
	// Overwrites any previous HAVING condition.
	//
	// Example:
	//
	//	builder.GroupBy("role").Having(WhereQuery{
	//	    "count(*)": {Op: ">", Value: 5},
	//	})
	Having(havingClauses WhereQuery) Selector

	// WithCTEBuilder adds a Common Table Expression (CTE) to the query.
	// It adjusts argument placeholders to avoid conflicts.
	// This function just add the defined CTE to the top of query.
	// You need to JOIN/LEFT JOIN the CTE builder to let main expression know that it should use CTE.
	//
	// Example:
	//
	//	cte := NewSQLSelectBuilder[Order]("orders").Select("id", "user_id")
	//	builder.WithCTEBuilder("recent_orders", cte.(*sql_query.SelectBuilder).SQLEloquentQuery)
	//
	// Generates:
	//
	//	WITH recent_orders AS (SELECT id, user_id FROM orders) ...
	WithCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) Selector

	// WithRecursiveCTEBuilder adds a Common Table Expression (CTE) to the query.
	// It adjusts argument placeholders to avoid conflicts.
	// This function just add the defined CTE to the top of query.
	// You need to JOIN/LEFT JOIN the CTE builder to let main expression know that it should use CTE.
	//
	// Example:
	//
	//	cte := NewSQLSelectBuilder[Order]("orders").Select("id", "user_id")
	//	builder.WithRecursiveCTEBuilder("recent_orders", cte.(*sql_query.SelectBuilder).SQLEloquentQuery)
	//
	// Generates:
	//
	//	WITH RECURSIVE recent_orders AS (SELECT id, user_id FROM orders) ...
	WithRecursiveCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) Selector

	// Add "UNION ALL" in between the queries
	// example:
	//  SELECT id
	// FROM categories
	// WHERE id = 0
	//
	// UNION ALL
	//
	// SELECT c.id
	// FROM categories c
	// INNER JOIN category_tree ct ON c.parent_id = ct.id
	UnionAll(cteBuilders ...*SQLEloquentQuery) Selector

	// Build finalizes the SELECT query and returns the query string and arguments.
	// Returns an error if the query is invalid (e.g., HAVING without GROUP BY).
	Build() (string, []any, error)
}

type SelectBuilder struct {
	*SQLEloquentQuery
}

type Sort struct {
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}

type SQLSort struct {
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"` // ASC | DESC
}

type Pagination struct {
	Page        int    `json:"page"      transform:"int"`
	Limit       int    `json:"limit"     transform:"int"`
	SortBy      string `json:"sortBy"    transform:"string"`
	SortOrder   int    `json:"sortOrder" transform:"int"`
	MultiSort   []Sort `json:"multiSort"`
	DefaultSort []Sort `json:"defaultSort"`
}
