package common_builders

import (
	"log"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/sql_query"
)

func UpdateEachBuilder(tableName string, rowIdentifier string, query sql_query.WhereQuery, body any) (string, []any) {
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
