package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

func UpdateEachBuilder(tableName string, rowIdentifier string, query map[string]sql_query.SQLCondition, body interface{}) (string, []interface{}) {
	res, args, err := sql_query.NewSQLUpdateBuilder(tableName).
		UpdateEach(body, rowIdentifier).
		Return("id").
		Where(query).
		Build()
	if err != nil {
		log.Println(err)
	}

	return res, args
}
