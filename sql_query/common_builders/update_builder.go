package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

func UpdateBuilder(tableName string, query map[string]sql_query.SQLCondition, body interface{}, returningColumn ...string) (string, []interface{}) {
	res, args, err := sql_query.NewSQLUpdateBuilder(tableName).
		Update(body).
		Return(returningColumn...).
		Where(query).
		Build()
	if err != nil {
		log.Println(err)
	}

	return res, args
}
