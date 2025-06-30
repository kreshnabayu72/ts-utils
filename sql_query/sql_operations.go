package sql_query

import "strings"

type SQLOperationsTypeEnum string

const (
	SQLOperationsTypeEnumUnknown    = SQLOperationsTypeEnum("unknown")
	SQLOperationsTypeEnumPagination = SQLOperationsTypeEnum("pagination")
	SQLOperationsTypeEnumFindOne    = SQLOperationsTypeEnum("findOne")
	SQLOperationsTypeEnumFind       = SQLOperationsTypeEnum("find")
	SQLOperationsTypeEnumInsertOne  = SQLOperationsTypeEnum("insertOne")
	SQLOperationsTypeEnumInsertMany = SQLOperationsTypeEnum("insertMany")
	SQLOperationsTypeEnumUpdateMany = SQLOperationsTypeEnum("updateMany")
	SQLOperationsTypeEnumUpdateOne  = SQLOperationsTypeEnum("updateOne")
	SQLOperationsTypeEnumDeleteMany = SQLOperationsTypeEnum("deleteMany")
	SQLOperationsTypeEnumDeleteOne  = SQLOperationsTypeEnum("deleteOne")
)

func GetSQLOperationType(query string) SQLOperationsTypeEnum {
	// Normalize and split the query
	normalized := strings.TrimSpace(strings.ToUpper(query))
	words := strings.Fields(normalized)
	if len(words) == 0 {
		return SQLOperationsTypeEnumUnknown
	}

	// Handle paginated SELECT using CTE and json_agg
	if strings.HasPrefix(normalized, "WITH") &&
		strings.Contains(normalized, "JSON_AGG(PAGINATED") &&
		strings.Contains(normalized, "COUNT FROM TOTAL_COUNT") {
		return SQLOperationsTypeEnumPagination
	}

	switch words[0] {
	case "SELECT":
		if strings.Contains(normalized, "LIMIT 1") {
			return SQLOperationsTypeEnumFindOne
		}
		return SQLOperationsTypeEnumFind

	case "INSERT":
		if strings.Count(normalized, "),(") > 0 {
			return SQLOperationsTypeEnumInsertMany
		}
		return SQLOperationsTypeEnumInsertOne

	case "UPDATE":
		if strings.Contains(normalized, "WHERE") &&
			strings.Contains(normalized, "IN (") {
			return SQLOperationsTypeEnumUpdateMany
		}
		if strings.Contains(normalized, "WHERE") {
			return SQLOperationsTypeEnumUpdateOne
		}
		return SQLOperationsTypeEnumUpdateMany // fallback

	case "DELETE":
		if strings.Contains(normalized, "WHERE") &&
			strings.Contains(normalized, "IN (") {
			return SQLOperationsTypeEnumDeleteMany
		}
		if strings.Contains(normalized, "WHERE") {
			return SQLOperationsTypeEnumDeleteOne
		}
		return SQLOperationsTypeEnumDeleteMany

	default:
		return SQLOperationsTypeEnumUnknown
	}
}
