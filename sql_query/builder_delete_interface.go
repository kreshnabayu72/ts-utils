package sql_query

// To ensure SQLDeleteBuilder has its own initial methods
// e.g. DeleteBuilder(...).Delete()...Rest
type SQLDeleteInitBuilder interface {
	// Delete implements SQLDeleteChainBuilder. (Only able to be called once)
	// Delete initializes a DELETE statement for the current table.
	// By default, it returns the deleted "id", but you can pass custom RETURNING columns.
	//
	// Example:
	//
	//	builder.Delete()                     // RETURNING id
	//	builder.Delete("id", "deleted_at")   // RETURNING id, deleted_at
	//
	// Generates:
	//
	//	DELETE FROM table_name RETURNING id
	Delete(returningColumns ...string) SQLDeleteChainBuilder
}

// To ensure method .Delete() has its own chaining methods
// e.g. .Delete(...).Using(...).Where(...).Build()
type SQLDeleteChainBuilder interface {
	// Where implements SQLDeleteChainBuilder. (Accumulates previous value if called again)
	Where(filters map[string]SQLCondition) SQLDeleteChainBuilder
	// WhereOr implements SQLDeleteChainBuilder. (Accumulates previous value if called again)
	WhereOr(filters ...map[string]SQLCondition) SQLDeleteChainBuilder

	// Using implements SQLDeleteChainBuilder. (Overrides previous value if called again)
	// Using adds a USING clause to the DELETE statement.
	// It is useful for multi-table DELETE with a join-like behavior.
	//
	// Example:
	//
	//	builder.Using([]string{"other_table"})
	//
	// Generates:
	//
	//	DELETE FROM table_name USING other_table
	Using(tables []string) SQLDeleteChainBuilder

	// buildDeleteQuery finalizes the DELETE query into a full SQL string + args.
	// It adds USING, WHERE, and RETURNING clauses if provided.
	// Prevents execution if CustomQuery is empty.
	//
	// Example Output:
	//
	//	DELETE FROM users USING roles WHERE users.role_id = roles.id RETURNING id
	Build() (string, []any, error)
}

// <---To wrap builder to its respective interface, used as pointer type in method of each builder--->
type DeleteBuilder struct {
	*SQLEloquentQuery
}
