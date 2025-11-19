package common_builders

import (
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

func IncrementBuilder(
	tableName string,
	query map[string]sql_query.SQLCondition,
	body map[string]any,
) (string, []interface{}) {
	res, args, err := sql_query.NewSQLUpdateBuilder(tableName).
		Increment(body).
		Return("id").
		Where(query).
		Build()
	if err != nil {
		log.Println(err)
	}

	return res, args
}
