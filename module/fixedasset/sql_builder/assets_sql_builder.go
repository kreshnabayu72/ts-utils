package sql_builder

import (
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type Asset struct {
	ID string `column:"a.id" json:"_id,omitempty"`
}

func NewAssetsSQLBuilder() *sql_query.SelectBuilder {
	builder := sql_query.NewSQLSelectBuilder[Asset]("assets a")
	return builder.(*sql_query.SelectBuilder)
}

func commonGetAssetListBuilder() *sql_query.SelectBuilder {
	assetBuilder := NewAssetsSQLBuilder()
	AttributesAggregationBuilder(assetBuilder, Group{}, AggregationOptions{
		TableName:  "groups",
		LocalField: "a.group_id",
		Alias:      "group",
		UseArray:   false,
		Preserve:   true,
	})
	AttributesAggregationBuilder(assetBuilder, Model{}, AggregationOptions{
		TableName:  "models",
		LocalField: "a.model_id",
		Alias:      "model",
		UseArray:   false,
		Preserve:   true,
	})
	assetBuilder.
		LeftJoin("brand_categories", "brand_categories.id = models.brand_category_id")
	AttributesAggregationBuilder(assetBuilder, Category{}, AggregationOptions{
		TableName:  "categories",
		LocalField: "brand_categories.category_id",
		Alias:      "category",
		UseArray:   false,
		Preserve:   true,
	})
	AttributesAggregationBuilder(assetBuilder, Brand{}, AggregationOptions{
		TableName:  "brands",
		LocalField: "brand_categories.brand_id",
		Alias:      "brand",
		UseArray:   false,
		Preserve:   true,
	})
	AttributesAggregationBuilder(assetBuilder, Name{}, AggregationOptions{
		TableName:  "asset_names",
		LocalField: "a.name_id",
		Alias:      "name",
		UseArray:   false,
		Preserve:   true,
	})
	AttributesAggregationBuilder(assetBuilder, DepreciationGroup{}, AggregationOptions{
		TableName:  "depreciation_groups",
		LocalField: "asset_names.depreciation_group_id",
		Alias:      "depreciationGroup",
		UseArray:   false,
		Preserve:   true,
	})
	GetAssetIDFromAliasCodeBuilder(assetBuilder, "assets a")

	assetBuilder.
		GroupBy("_id", "models.id", "brands.id", "brand_categories.id", "categories.id", "groups.id", "asset_names.id", "depreciation_groups.id", "a.id")

	return assetBuilder
}

func GetAllAssetsBuilder(filters map[string]sql_query.SQLCondition) (string, []interface{}, error) {
	return commonGetAssetListBuilder().Where(filters).Build()
}

func GetAllAssetsWithRoleBuilder(filters map[string]sql_query.SQLCondition) (string, []interface{}, error) {
	commonBuilder := commonGetAssetListBuilder()
	TransactionRolePermissionAggregation(
		commonBuilder,
		1000000000000000001,
		"a.group_id",
		"damagedRole",
		RoleRestrictionOptions{},
		true,
	)
	return commonBuilder.Where(filters).Build()
}
