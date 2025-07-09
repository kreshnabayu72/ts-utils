package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

func InsertOneBuilder(tableName string, body interface{}) (string, []interface{}) {
	res, args, err := sql_query.NewSQLInsertBuilder(tableName).
		Insert(body, "id").
		Build()

	if err != nil {
		log.Println(err)
	}

	return res, args
}
