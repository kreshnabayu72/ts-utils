package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

func InsertBuilder(tableName string, body any, returningColumn ...string) (string, []any) {
	res, args, err := sql_query.NewSQLInsertBuilder(tableName).
		Insert(body, returningColumn...).
		Build()
	if err != nil {
		log.Println(err)
	}

	return res, args
}
