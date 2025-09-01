package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

func UpdateGroupQuota(service service.PostgreSqlService, ctx context.Context, setValue map[string]any, groupId int, licenseType new_model.GroupQuotaType) error {
	queryQuota, argsQuota, _ := sql_query.NewSQLUpdateBuilder(
		db.GroupQuotaModelName,
	).
		Increment(setValue).
		Where(map[string]sql_query.SQLCondition{
			"group_id": {Operator: sql_query.SQLOperatorEqual, Value: groupId},
			"type":     {Operator: sql_query.SQLOperatorEqual, Value: licenseType},
		}).
		Build()
	if _, err := service.UpdateOne(
		ctx,
		queryQuota,
		argsQuota...,
	); err != nil {
		return err
	}

	return nil
}
