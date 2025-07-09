package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

func CountBuilder(tableName string, query map[string]sql_query.SQLCondition) (string, []interface{}) {
	res, args, err := sql_query.NewSQLCountBuilder(tableName).
		Where(query).
		Build()

	if err != nil {
		log.Println(err)
	}

	return res, args
}
