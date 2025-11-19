package sql_builder

import (
	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

type Asset struct {
	ID string `column:"a.id" json:"_id,omitempty"`
}

func NewAssetsSQLBuilder() *sql_query.SelectBuilder {
	builder := sql_query.NewSQLSelectBuilder[Asset]("assets a")
	return builder.(*sql_query.SelectBuilder)
}

// Extendable builder for common asset list, also able to exclude some default fields from this builder.
// Can be extended with select builder's chain methods or used in CTE.
func CommonGetAssetListBuilder(excludeFields ...string) *sql_query.SelectBuilder {
	assetBuilder := NewAssetsSQLBuilder()

	if !sql_query.ArrayIncludes(excludeFields, "group") {
		AttributesAggregationBuilder(assetBuilder, Group{}, AggregationOptions{
			TableName:  "groups",
			LocalField: "a.group_id",
			Alias:      "group",
			UseArray:   false,
			Preserve:   true,
		})
		// assetBuilder.GroupBy("groups.id")
	}

	if !sql_query.ArrayIncludes(excludeFields, "model") {
		AttributesAggregationBuilder(assetBuilder, Model{}, AggregationOptions{
			TableName:  "models",
			LocalField: "a.model_id",
			Alias:      "model",
			UseArray:   false,
			Preserve:   true,
		})
		// assetBuilder.GroupBy("models.id")
	}

	if !sql_query.ArrayIncludes(excludeFields, "category") {
		assetBuilder.
			LeftJoin("brand_categories", "brand_categories.id = models.brand_category_id")
		AttributesAggregationBuilder(assetBuilder, Category{}, AggregationOptions{
			TableName:  "categories",
			LocalField: "brand_categories.category_id",
			Alias:      "category",
			UseArray:   false,
			Preserve:   true,
		})
		// assetBuilder.GroupBy("categories.id")
	}

	if !sql_query.ArrayIncludes(excludeFields, "brand") {
		AttributesAggregationBuilder(assetBuilder, Brand{}, AggregationOptions{
			TableName:  "brands",
			LocalField: "brand_categories.brand_id",
			Alias:      "brand",
			UseArray:   false,
			Preserve:   true,
		})
		// assetBuilder.GroupBy("brands.id")
	}

	if !sql_query.ArrayIncludes(excludeFields, "name") {
		AttributesAggregationBuilder(assetBuilder, Name{}, AggregationOptions{
			TableName:  "asset_names",
			LocalField: "a.name_id",
			Alias:      "name",
			UseArray:   false,
			Preserve:   true,
		})
		// assetBuilder.GroupBy("asset_names.id")
	}

	if !sql_query.ArrayIncludes(excludeFields, "depreciationGroup") {
		AttributesAggregationBuilder(assetBuilder, DepreciationGroup{}, AggregationOptions{
			TableName:  "depreciation_groups",
			LocalField: "asset_names.depreciation_group_id",
			Alias:      "depreciationGroup",
			UseArray:   false,
			Preserve:   true,
		})
		// assetBuilder.GroupBy("depreciation_groups.id")
	}

	if !sql_query.ArrayIncludes(excludeFields, "assetId") {
		GetAssetIDFromAliasCodeBuilder(assetBuilder, "a.id")
	}

	return assetBuilder
}

// Ex-dummy experiment
func GetAllAssetsBuilder(filters map[string]sql_query.SQLCondition) (string, []interface{}, error) {
	return CommonGetAssetListBuilder().Where(filters).Build()
}

// Ex-dummy experiment
func GetAllAssetsWithRoleBuilder(
	filters map[string]sql_query.SQLCondition,
) (string, []interface{}, error) {
	commonBuilder := CommonGetAssetListBuilder()
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
