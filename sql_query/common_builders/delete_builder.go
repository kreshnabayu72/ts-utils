package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

func DeleteBuilder(tableName string, query map[string]sql_query.SQLCondition) (string, []any) {
	res, args, err := sql_query.NewSQLDeleteBuilder(tableName).
		Delete("id").
		Where(query).
		Build()

	if err != nil {
		log.Println(err)
	}

	return res, args
}
