package common_builders

import (
	"log"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/sql_query"
)

func UpdateBuilder(tableName string, query map[string]sql_query.SQLCondition, body any, returningColumn ...string) (string, []any) {
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
