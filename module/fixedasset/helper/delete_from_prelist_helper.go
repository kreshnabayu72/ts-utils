package helper

import (
	"context"
	"fmt"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

type PrelistModuleEnum string

const (
	PrelistModuleEnumAssignment = "assignment"
	PrelistModuleEnumBorrowing  = "borrowing"
	PrelistModuleEnumTransfer   = "transfer"
)

func DeleteAssetFromPrelist(service service.PostgreSqlService, ctx context.Context, assetId int, module PrelistModuleEnum) error {
	prelistTable := fmt.Sprintf("%s_prelists", module)
	prelistRequestTable := fmt.Sprintf("%s_prelist_requests", module)

	queryDeletePrelistRequest, argsDeletePrelistRequest, _ := sql_query.NewSQLDeleteBuilder(prelistRequestTable).
		Delete("prelist_id as prelistId").
		Where(map[string]sql_query.SQLCondition{
			"asset_id": {Operator: sql_query.SQLOperatorEqual, Value: assetId},
		}).
		Build()

	returningValue := new([]struct {
		PrelistID int `json:"prelistId"`
	})
	if err := service.SelectMany(returningValue, ctx, queryDeletePrelistRequest, argsDeletePrelistRequest...); err != nil {
		return err
	}

	prelistIds := []int{}
	for _, prelist := range *returningValue {
		prelistIds = append(prelistIds, prelist.PrelistID)
	}

	if len(prelistIds) > 0 {
		subQuery, _, _ := sql_query.NewSQLSelectBuilder[any](prelistRequestTable).
			Select("1").
			Where(map[string]sql_query.SQLCondition{
				prelistRequestTable + ".prelist_id": {
					Operator: sql_query.SQLOperatorEqual,
					Value:    prelistTable + ".id",
					IsRef:    true,
				},
			}).
			Build()

		queryDeletePrelist, argsDeletePrelist, _ := sql_query.NewSQLDeleteBuilder(prelistTable).
			Delete().
			Where(map[string]sql_query.SQLCondition{
				prelistTable + ".id": {Operator: sql_query.SQLOperatorAny, Value: prelistIds},
				"":                   {Operator: sql_query.SQLOperatorNotExist, Value: "(" + subQuery + ")", IsSubQuery: true},
			}).
			Build()

		if _, err := service.DeleteMany(ctx, queryDeletePrelist, argsDeletePrelist...); err != nil {
			return err
		}
	}

	return nil
}
