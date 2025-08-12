package sql_builder

import (
	"fmt"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type AggregationOptions struct {
	TableName    string // Lookup target's table
	LocalField   string // Foreign key of main table that reference to lookup target's table
	Alias        string // Alias for generated jsonb
	FilterClause string // Usage example: fmt.Sprintf("%s.id is not null", options.TableName) to serve value as null if doesn't match
	UseArray     bool   // Querying with jsonb_agg or jsonb_build_object based on given boolean
	Preserve     bool   // Preserve null value
}

type AliasCode struct {
	AssetID string `column:"CONCAT_WS(ac.separator, MAX(CASE WHEN acd.name = 'Initial' AND acd.is_checked THEN ac.company_initial END), MAX(CASE WHEN acd.name = 'Group' AND acd.is_checked THEN g.code END), MAX(CASE WHEN acd.name = 'Category' AND acd.is_checked THEN c.code END), MAX(CASE WHEN acd.name = 'Asset Name' AND acd.is_checked THEN an.code END), MAX(CASE WHEN acd.name = 'Purchase Date' AND acd.is_checked THEN TO_CHAR(a.purchase_date, 'MMYY') END), MAX(CASE WHEN acd.name = 'Increment' AND acd.is_checked THEN LPAD(a.sequence_number::text, 4, '0') END))" json:"assetId"`
}

type Brand struct {
	ID   string `column:"brands.id::text" json:"_id,omitempty"`
	Name string `column:"brands.name" json:"name"`
}

type Category struct {
	ID   string `column:"categories.id::text" json:"_id,omitempty"`
	Name string `column:"categories.name" json:"name"`
	Key  int    `column:"categories.key" json:"key"`
}

type BrandCategory struct {
	ID         string `column:"models.id::text" json:"_id,omitempty"`
	BrandID    string `column:"brands.id::text" json:"brandId"`
	CategoryID string `column:"categories.id" json:"categoryId"`
}

type Model struct {
	ID              string `column:"models.id::text" json:"_id,omitempty"`
	BrandCategoryID string `column:"brand_categories.id::text" json:"brandCategoryId"`
	Name            string `column:"models.name" json:"name"`
}

type Name struct {
	ID      string            `column:"asset_names.id::text" json:"_id,omitempty"`
	Name    string            `column:"asset_names.name" json:"name"`
	TagType new_model.TagType `column:"asset_names.tag_type" json:"tagType"`
}

type DepreciationGroup struct {
	ID   string `column:"depreciation_groups.id::text" json:"_id,omitempty"`
	Name string `column:"depreciation_groups.name" json:"name"`
}

type Group struct {
	ID   string `column:"groups.id::text" json:"_id,omitempty"`
	Name string `column:"groups.name" json:"name"`
}

// AttributesAggregationBuilder joins a related table and builds a JSON
// aggregation for it.
//
// If Preserve is true → uses LEFT JOIN (keeps unmatched rows).
// If UseArray is false → automatically GROUP BY <table>.id so the JSON
// represents a single object instead of an array.
//
// Handy for quickly pulling related data into one field without manually
// writing joins, JSON aggregation, or group by logic.
func AttributesAggregationBuilder(builder *sql_query.SelectBuilder, dto any, options AggregationOptions) {
	joinType := builder.Join
	if options.Preserve {
		joinType = builder.LeftJoin
	}

	joinType(options.TableName, options.TableName+".id"+" = "+options.LocalField)

	// Now `AttributesAggregationBuilder` implement filterclause option.
	builder.SelectJSONAggregate(options.Alias, dto, options.FilterClause, options.UseArray)
}

// Tested result WTF/MG/ELEC/LTP/0724/1
func GetAssetIDFromAliasCodeBuilder(builder *sql_query.SelectBuilder, tableName string) {
	res, args, err := sql_query.NewSQLSelectBuilder[AliasCode](tableName).
		Join("alias_codes ac", "true").
		Join("groups g", "g.id = a.group_id").
		Join("models m", "m.id = a.model_id").
		Join("brand_categories bc", "bc.id = m.brand_category_id").
		Join("categories c", "c.id = bc.category_id").
		Join("asset_names an", "an.id = a.name_id").
		LeftJoin("alias_codes_data acd", "acd.alias_code_id = ac.id").
		GroupBy("ac.separator").
		SetLimit(1).
		Build()
	if err != nil {
		log.Println(err)
	}

	builder.Columns = append(builder.Columns, fmt.Sprintf("(%s) AS assetId", res))
	builder.Args = append(builder.Args, args...)
}
