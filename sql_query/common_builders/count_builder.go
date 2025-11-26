package common_builders

import (
	"log"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/sql_query"
)

func CountBuilder(tableName string, query map[string]sql_query.SQLCondition) (string, []any) {
	res, args, err := sql_query.NewSQLCountBuilder(tableName).
		Where(query).
		Build()

	if err != nil {
		log.Println(err)
	}

	return res, args
}
