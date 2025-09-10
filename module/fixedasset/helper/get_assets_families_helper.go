package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type GetAsset struct {
	ID         int                   `json:"_id" column:"a.id"`
	ParentID   *int                  `json:"parentId" column:"a.parent_id"`
	FirstImage *string               `json:"firstImage" column:"a.first_image"`
	AssetID    string                `json:"assetId"`
	Status     new_model.AssetStatus `json:"status" column:"a.status"`
	Brand      sql_builder.Brand     `json:"brand"`
	Model      sql_builder.Model     `json:"model"`
	Group      sql_builder.Group     `json:"group"`
	Name       sql_builder.Name      `json:"name"`
}

func GetAssetsWithFamilies(service service.PostgreSqlService, ctx context.Context, assetIds []int) (*[]GetAsset, error) {
	assetFamilies := []GetAsset{}

	for i := range assetIds {
		assetId := assetIds[i]

		queryAsset, argsAsset, _ := sql_builder.CommonGetAssetListBuilder("depreciationGroup").
			Where(
				map[string]sql_query.SQLCondition{
					"a.id": {Operator: sql_query.SQLOperatorEqual, Value: assetId},
				},
			).
			Build()
		var asset GetAsset
		if err := service.SelectOne(&asset, ctx, queryAsset, argsAsset...); err != nil {
			return nil, err
		}

		queryFamily, argsFamily, _ := sql_builder.CommonGetAssetListBuilder("depreciationGroup").
			WhereOr(
				map[string]sql_query.SQLCondition{
					"a.id": {Operator: sql_query.SQLOperatorEqual, Value: assetId},
				},
				map[string]sql_query.SQLCondition{
					"a.id": {Operator: sql_query.SQLOperatorEqual, Value: asset.ParentID},
				},
				map[string]sql_query.SQLCondition{
					"a.parent_id": {Operator: sql_query.SQLOperatorEqual, Value: assetId},
				},
			).
			Build()
		assetFamily := []GetAsset{}
		if err := service.SelectMany(&assetFamily, ctx, queryFamily, argsFamily...); err != nil {
			return nil, err
		}
		assetFamilies = append(assetFamilies, assetFamily...)
	}

	return &assetFamilies, nil
}
