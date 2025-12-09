package sql_query

import "reflect"

// To ensure SQLInsertBuilder has its own initial methods
// e.g. InsertBuilder(...).Insert()...Rest
type InsertStarter interface {
	// Insert implements Inserter. (Only able to be called once)
	// Insert builds an INSERT query from either a single struct or a slice of structs.
	// - For a single struct → generates one row.
	// - For a slice → generates multiple rows.
	//
	// By default, it will RETURN the id::text column,
	// but you can pass custom returning columns:
	//
	//	Insert(user, "id", "name")
	//	Insert([]User{u1, u2})
	Insert(values any, returningColumns ...string) Inserter

	// insertSingle handles the insert logic for a single struct.
	// It auto-generates a Snowflake ID and builds the VALUES list
	// based on struct field tags (using "column" or "json").
	cachedInsertSingle(v reflect.Value) Inserter
	// insertMany handles inserting multiple structs at once.
	// It generates a Snowflake ID for each row, creates placeholders
	// for every value, and builds the full VALUES (...), (...), ... list.
	//
	// Example Result:
	//
	//	INSERT INTO users (id, name, email)
	//	VALUES ($1, $2, $3), ($4, $5, $6)
	cachedInsertMany(slice reflect.Value) Inserter
	// preBuild is a helper that assembles the final
	// INSERT INTO ... (columns) VALUES ... statement
	// after the columns and value placeholders have been prepared.
	preBuild(columns, valuePlaceholders []string)
}

// To ensure method .Insert() has its own chaining methods
// e.g. .Insert(...).Build()
type Inserter interface {
	// By default, the Insert builder includes all fields, even if their values are zero or nil.
	// Calling this method tells the builder to skip fields with zero values when generating the SET clause.
	//
	// Note: This option only affects single-row Insert operations.
	// It has no effect on bulk INSERT, because all rows in those operations must have the same set of columns.
	ExcludeEmpty() Inserter
	// Insert implements Inserter. (Only able to be called once, will override previous call)
	// Conflict adds an ON CONFLICT clause to the insert statement.
	// Example:
	//
	//	.Conflict("(id)", "NOTHING")
	//	-> INSERT ... ON CONFLICT (id) DO NOTHING
	Conflict(constraint, do string) Inserter
	// buildInsertQuery finalizes the insert query into SQL string + args.
	// It prevents unsafe cases (like adding filters, joins, or pagination)
	// and appends RETURNING and ON CONFLICT if defined.
	Build() (string, []any, error)
}

type InsertBuilder struct {
	*SQLEloquentQuery
}
