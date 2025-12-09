package common_builders

import (
	"log"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/sql_query"
)

func IncrementBuilder(
	tableName string,
	query sql_query.WhereQuery,
	body map[string]any,
) (string, []any) {
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
