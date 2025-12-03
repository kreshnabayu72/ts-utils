package sql_query

import "reflect"

type UpdateQuery map[string]any

type UpdateRawSQL struct {
	Expr string
	Args []any
}

// type string struct {
// 	Key       string
// 	Operator  SQLOperators // e.g., '=', '>', '<=', 'LIKE', 'IN', 'IS NULL'
// 	Value     any  // could be a single value, slice, or nil
// 	IsRef     bool         // to determine whether WHERE is targeting literal value or reference
// 	ExtraArgs []any
// }

type UpdateCases interface {
	// Case adds a WHEN clause to the current CASE expression defined by AddCase.
	// It appends the condition and its resulting value.
	//
	// Parameters:
	//   - conditions: a MultiFilterCondition containing AND and OR filters used for this WHEN clause.
	//   - value: the resulting value when the conditions match.
	//   - isRef: whether the value is a SQL column reference (true) or a literal that should be parameterized (false).
	//
	// Example:
	//   c.Case(MultiFilterCondition{
	//     And: WhereQuery{
	//       "status":       {Operator: SQLOperatorEqual, Value: "Not Paired Yet"},
	//       body.TagType:   {SourceIsValue: true, Operator: SQLOperatorEqual, Value: "Non TAG"},
	//     },
	//   }, "Available", false)
	Case(conditions MultiFilterCondition, value any, isRef bool)
	// Else sets the default ELSE value for the current CASE expression.
	// It is used when none of the WHEN conditions match.
	//
	// Parameters:
	//   - value: the fallback value if no WHEN clause matches.
	//   - isRef: whether the value is a SQL column reference (true) or a literal that should be parameterized (false).
	//
	// Example:
	//   c.Else("status", true) // fallback to the original "status" column value
	Else(value any, isRef bool)
}

// To ensure SQLUpdateBuilder has its own initial methods
// e.g. UpdateBuilder(...).Update()...Rest
type UpdateStarter interface {
	// Update builds an UPDATE query using UpdateQuery (map[string]any).
	//
	// It supports two value types:
	//   • Normal values (int, string, time.Time, etc.): converted into positional parameters ($1, $2, …).
	//   • Raw SQL expressions: allows embedding expressions with placeholders ("?")
	//     which will be replaced with proper PostgreSQL-style parameters ($1, $2, …).
	//
	// Example:
	//
	//     builder.Update(sql_query.UpdateQuery{
	//         "used": 5,
	//         "available": sql_query.UpdateRawSQL{
	//             Expr: "CASE WHEN available IS NULL THEN NULL ELSE available - ? END",
	//             Args: []any{5},
	//         },
	//     })
	//
	// By default, it will automatically add `updated_at = NOW()` if not explicitly provided.
	Update(values UpdateQuery) Updater
	// UpdateEach updates multiple rows at once using VALUES() with a slice of structs.
	// Matches rows using the given rowIdentifier (e.g., "id").
	//
	// Example:
	//
	//	builder.UpdateEach([]User{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}, "id")
	//
	// → UPDATE users SET name = v.name, updated_at = NOW()
	//
	//	FROM (VALUES ($1,$2),($3,$4)) AS v(id,name,updated_at)
	//	WHERE users.id = v.id
	UpdateEach(values any, rowIdentifier string) Updater

	// AddCase initializes a conditional CASE expression for the given column in an UPDATE statement.
	// It use completely different CASE expressions from previous AddCase and allows chaining multiple conditional branches using Case() and Else().
	//
	// Example:
	//   builder.AddCase("status", func(b UpdateCases) {
	//       b.Case(...)
	//       b.Else(...)
	//   })
	//
	// Parameters:
	//   - setColumn: the name of the column to be conditionally updated.
	//   - fn: a function that defines the CASE conditions using the provided UpdateCases interface.
	//
	// Returns:
	//   - Updater: the builder itself for method chaining.
	AddCase(setColumn string, fn func(b UpdateCases)) Updater

	// Increment is used to replace Update() for adding int value cause update cant handle that
	// Increment builds an UPDATE query that increases integer columns by a given value.
	// Automatically sets updated_at = NOW().
	//
	// Example:
	//
	//	builder.Increment(map[string]any{"count": 1})
	//
	// → UPDATE table SET "count" = "count" + $1, "updated_at" = NOW()
	Increment(values map[string]any) Updater

	// updateEachClausesGenerator looks at every struct in the slice and builds:
	//  1. The SET part of the query (e.g., "name = v.name"),
	//  2. The column names to use in the VALUES table,
	//  3. The placeholders ($1, $2, …) for each row.
	//
	// It's the core logic behind UpdateEach, letting you bulk-update rows
	// by matching them with a unique rowIdentifier (e.g. "id" or "name").
	//
	// This function loop through slice, returning generated clauses for everything in it
	// Expected result:
	// setClauses contains: [value = v.value, ...etc].
	// valueClauses contains: [value, ...etc].
	// valuePlaceholders contains: [($1, $2, $3, ...etc), ($4, $5, ...etc), ...etc]
	updateEachClausesGenerator(
		slice reflect.Value,
		rowIdentifier string,
	) ([]string, []string, []string)

	// extractUpdateFieldsMap takes a map[string]any and returns a slice of
	// SET clauses for an UPDATE query. It uses the map keys as column names.
	// It also checks if an "updated_at" key exists, so the builder knows
	// whether to add "updated_at = NOW()" automatically.
	extractUpdateFieldsMap(v map[string]any) ([]string, bool)
}

// To ensure method .Update() has its own chaining methods
// e.g. .Update(...).From(...).Build()
type Updater interface {
	// AddCase initializes a conditional CASE expression for the given column in an UPDATE statement.
	// It clears any existing CASE expressions and allows chaining multiple conditional branches using Case() and Else().
	//
	// Example:
	//   builder.AddCase("status", func(b UpdateCases) {
	//       b.Case(...)
	//       b.Else(...)
	//   })
	//
	// Parameters:
	//   - setColumn: the name of the column to be conditionally updated.
	//   - fn: a function that defines the CASE conditions using the provided UpdateCases interface.
	//
	// Returns:
	//   - Updater: the builder itself for method chaining.
	AddCase(setColumn string, fn func(b UpdateCases)) Updater

	// By default, the Update builder includes all fields, even if their values are zero or nil.
	// Calling this method tells the builder to skip fields with zero values when generating the SET clause.
	//
	// Note: This option only affects single-row Update operations.
	// It has no effect on bulk UPDATE, because all rows in those operations must have the same set of columns.
	ExcludeEmpty() Updater
	// Update implements Updater. (Only able to be called once, will override previous call).
	// Conflict adds an ON CONFLICT clause with the specified constraint and action.
	//
	// Example:
	//
	//	builder.Conflict("(id)", "NOTHING")
	//
	// → ON CONFLICT (id) DO NOTHING
	Conflict(constraint, do string) Updater
	// Where implements Updater. (Accumulates previous value if called again)
	Where(filters WhereQuery) Updater
	// WhereOr implements Updater. (Accumulates previous value if called again)
	WhereOr(filters ...WhereQuery) Updater

	// Join adds an INNER JOIN clause with the specified ON condition.
	//
	// Example:
	//
	//	builder.Join("roles r", "r.id = u.role_id")
	Join(table string, onCondition string, additionalConditions ...WhereQuery) Updater
	// LeftJoin adds a LEFT JOIN clause with the specified ON condition.
	//
	// Example:
	//
	//	builder.LeftJoin("roles r", "r.id = u.role_id")
	LeftJoin(table string, onCondition string, additionalConditions ...WhereQuery) Updater

	// WithCTEBuilder adds a Common Table Expression (CTE) to the query.
	// It adjusts argument placeholders to avoid conflicts.
	// This function just add the defined CTE to the top of query.
	// You need to JOIN/LEFT JOIN the CTE builder to let main expression know that it should use CTE.
	//
	// Example:
	//
	//	builder.WithCTEBuilder("recent_orders", cte.(*sql_query.SelectBuilder).SQLEloquentQuery)
	//
	// Generates:
	//
	//	WITH recent_orders AS (SELECT id, user_id FROM orders) ...
	WithCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) Updater

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
	WithRecursiveCTEBuilder(cteName string, cteBuilder *SQLEloquentQuery) Updater

	// Return implements Updater. (Only able to be called once, overrides previous values if re-called).
	// Return sets the columns to return after the update.
	// Defaults to RETURNING id if no column is provided.
	Return(columns ...string) Updater

	// From implements Updater. (Overrides previous value if called again)
	// From adds a FROM clause to the UPDATE query, allowing joins with other tables.
	//
	// Example:
	//
	//	builder.From([]string{"users u", "roles r"})
	//
	// → FROM users u, roles r
	From(tables []string) Updater

	// buildUpdateQuery constructs the final UPDATE query string and its arguments.
	// Ensures that CustomQuery is set and that a WHERE clause exists for safety.
	Build() (string, []any, error)
}

type UpdateBuilder struct {
	*SQLEloquentQuery
}
