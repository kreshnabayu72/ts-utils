package sql_builder

import (
	"fmt"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/db"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
	util_sql_query "gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

func UpdateAssetAliasCodeBuilder() (string, []interface{}, error) {
	generatedCodeBuilder := generatedCodesCte()

	type editAssetAliasCode struct {
		AliasCode sql_query.UpdateRawSQL `json:"aliasCode" column:"alias_code"`
		UpdatedAt sql_query.UpdateRawSQL `json:"updatedAt" column:"updated_at"`
	}

	return util_sql_query.NewSQLUpdateBuilder(db.AssetTableName).
		Update(editAssetAliasCode{
			AliasCode: sql_query.UpdateRawSQL{
				Expr: `gc.dynamic_code`,
			},
			UpdatedAt: sql_query.UpdateRawSQL{
				Expr: `assets.updated_at`,
			},
		}).
		WithCTEBuilder("generated_codes", generatedCodeBuilder.(*util_sql_query.SelectBuilder).SQLEloquentQuery).
		From([]string{"generated_codes gc"}).
		Return("assets.id").
		Where(map[string]util_sql_query.SQLCondition{
			"": {
				Operator: sql_query.SQLOperatorRaw,
				Value:    "assets.id = gc.id",
			},
		}).
		Build()
}

func generatedCodesCte() util_sql_query.SQLSelectChainBuilder {
	return util_sql_query.NewSQLSelectBuilder[any]("raw_asset_data r").
		WithCTEBuilder("raw_asset_data", rawAssetDataCte().(*util_sql_query.SelectBuilder).SQLEloquentQuery).
		WithCTEBuilder("parts_map(id, part_name, part_value)", partsMapDataBuilder().(*util_sql_query.SelectBuilder).SQLEloquentQuery).
		Select(
			"r.id",
			`STRING_AGG(
				p.part_value, 
				(case 
					when r.separator != 'None' 
						then r.separator 
					else '' 
				end) 
				ORDER BY acd."order" ASC
			) AS dynamic_code`).
		Join("parts_map p", "p.id = r.id").
		Join("alias_codes_data acd", "p.part_name = acd.name").
		Where(map[string]util_sql_query.SQLCondition{
			"acd.is_checked": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    true,
			},
		}).
		GroupBy("r.id")
}

func rawAssetDataCte() util_sql_query.SQLSelectChainBuilder {
	return util_sql_query.NewSQLSelectBuilder[any]("assets a").
		Select(
			"a.id",
			"ac.company_initial",
			"ac.separator",
			"g.code AS group_code",
			"c.code AS category_code",
			"COALESCE(TO_CHAR(a.purchase_date, 'MMYY'), '0000') AS purchase_code",
			"LPAD(a.sequence_number::text, 4, '0') AS sequence_code",
			"COALESCE(NULLIF(an.code, ''), '0000') AS name_code",
		).
		LeftJoin("alias_codes ac", "TRUE").
		LeftJoin("asset_names an", "an.id = a.name_id").
		LeftJoin(`"groups" g`, "g.id = a.group_id").
		LeftJoin(`models m`, "m.id = a.model_id").
		LeftJoin(`brand_categories bc`, "bc.id = m.brand_category_id").
		LeftJoin(`categories c`, "c.id = bc.category_id")
}

func partsMapDataBuilder() util_sql_query.SQLSelectChainBuilder {
	return util_sql_query.NewSQLSelectBuilder[any]("").
		UnionAll(
			getSingleRawAssetMapData("Initial", "company_initial"),
			getSingleRawAssetMapData("Group", "group_code"),
			getSingleRawAssetMapData("Category", "category_code"),
			getSingleRawAssetMapData("Asset Name", "name_code"),
			getSingleRawAssetMapData("Purchase Date", "purchase_code"),
			getSingleRawAssetMapData("Asset Increment", "sequence_code"),
		)
}

func getSingleRawAssetMapData(aliasCodeKey string, assetAttributeKey string) *util_sql_query.SQLEloquentQuery {
	return util_sql_query.NewSQLSelectBuilder[any]("raw_asset_data").
		Select(
			"id",
			fmt.Sprintf("'%s'", aliasCodeKey),
			assetAttributeKey,
		).(*util_sql_query.SelectBuilder).SQLEloquentQuery
}
