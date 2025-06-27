package sql_query

import "strings"

type SQLOperationsTypeEnum string

const (
	SQLOperationsTypeEnumUnknown    = SQLOperationsTypeEnum("unknown")
	SQLOperationsTypeEnumFindOne    = SQLOperationsTypeEnum("findOne")
	SQLOperationsTypeEnumFind       = SQLOperationsTypeEnum("find")
	SQLOperationsTypeEnumInsertOne  = SQLOperationsTypeEnum("insertOne")
	SQLOperationsTypeEnumInsertMany = SQLOperationsTypeEnum("insertMany")
	SQLOperationsTypeEnumUpdateMany = SQLOperationsTypeEnum("updateMany")
	SQLOperationsTypeEnumUpdateOne  = SQLOperationsTypeEnum("updateOne")
	SQLOperationsTypeEnumDeleteMany = SQLOperationsTypeEnum("DeleteMany")
	SQLOperationsTypeEnumDeleteOne  = SQLOperationsTypeEnum("DeleteOne")
)

func GetSQLOperationType(query string) SQLOperationsTypeEnum {
	// Normalize and split the query
	normalized := strings.TrimSpace(strings.ToUpper(query))
	words := strings.Fields(normalized)
	if len(words) == 0 {
		return SQLOperationsTypeEnumUnknown
	}

	switch words[0] {
	case "SELECT":
		if strings.Contains(normalized, "LIMIT 1") {
			return "findOne"
		}
		return "find"

	case "INSERT":
		if strings.Count(normalized, "),(") > 0 {
			return "insertMany"
		}
		return "insertOne"

	case "UPDATE":
		if strings.Contains(normalized, "WHERE") &&
			strings.Contains(normalized, "IN (") {
			return "updateMany"
		}
		if strings.Contains(normalized, "WHERE") {
			return "updateOne"
		}
		return "updateMany" // fallback

	case "DELETE":
		if strings.Contains(normalized, "WHERE") &&
			strings.Contains(normalized, "IN (") {
			return "deleteMany"
		}
		if strings.Contains(normalized, "WHERE") {
			return "deleteOne"
		}
		return "deleteMany"

	default:
		return "unknown"
	}
}
