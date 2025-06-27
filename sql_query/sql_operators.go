package sql_query

type SQLOperators string

const (
	// Comparison
	SQLOperatorEqual       SQLOperators = "="
	SQLOperatorNotEqual    SQLOperators = "!=" // or "<>"
	SQLOperatorGreaterThan SQLOperators = ">"
	SQLOperatorLessThan    SQLOperators = "<"
	SQLOperatorGTE         SQLOperators = ">="
	SQLOperatorLTE         SQLOperators = "<="

	// Set
	SQLOperatorIn    SQLOperators = "IN"
	SQLOperatorNotIn SQLOperators = "NOT IN"

	// Pattern matching
	SQLOperatorLike     SQLOperators = "LIKE"
	SQLOperatorNotLike  SQLOperators = "NOT LIKE"
	SQLOperatorILike    SQLOperators = "ILIKE"     // PostgreSQL only
	SQLOperatorNotILike SQLOperators = "NOT ILIKE" // PostgreSQL only

	// Null checks
	SQLOperatorIsNull    SQLOperators = "IS NULL"
	SQLOperatorIsNotNull SQLOperators = "IS NOT NULL"

	// Range
	SQLOperatorBetween    SQLOperators = "BETWEEN"
	SQLOperatorNotBetween SQLOperators = "NOT BETWEEN"
)
