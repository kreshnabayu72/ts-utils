package helper

import (
	"context"
	"fmt"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/jackc/pgx/v5"
)

func GetReaderBySerialNumber(ctx context.Context, service service.PostgreSqlService, serialNumber string) (*new_model.HandheldReader, error) {
	queryReader, argsReader, _ := sql_query.
		NewSQLSelectBuilder[new_model.HandheldReader](db.HandheldReaderTableName).
		Where(map[string]sql_query.SQLCondition{
			"serial_number": {Operator: sql_query.SQLOperatorILike, Value: fmt.Sprintf("%%%s%%", serialNumber)},
		}).
		Build()
	var reader new_model.HandheldReader
	if err := service.SelectOne(&reader, ctx, queryReader, argsReader...); err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return &reader, nil
}
