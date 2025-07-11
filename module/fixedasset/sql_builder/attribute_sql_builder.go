package sql_builder

import (
	"fmt"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type AggregationOptions struct {
	TableName  string
	LocalField string
	Alias      string
	UseArray   bool
	Preserve   bool
}

type AliasCode struct {
	AssetID string `column:"CONCAT_WS(ac.separator, MAX(CASE WHEN acd.name = 'Initial' AND acd.is_checked THEN ac.company_initial END), MAX(CASE WHEN acd.name = 'Group' AND acd.is_checked THEN g.code END), MAX(CASE WHEN acd.name = 'Category' AND acd.is_checked THEN c.code END), MAX(CASE WHEN acd.name = 'Asset Name' AND acd.is_checked THEN an.code END), MAX(CASE WHEN acd.name = 'Purchase Date' AND acd.is_checked THEN TO_CHAR(a.purchase_date, 'MMYY') END), MAX(CASE WHEN acd.name = 'Increment' AND acd.is_checked THEN LPAD(a.sequence_number::text, 4, '0') END))" json:"assetId"`
}

type Brand struct {
	ID   string `column:"brands.id" json:"_id,omitempty"`
	Name string `column:"brands.name" json:"name"`
}

type Category struct {
	ID   string `column:"categories.id" json:"_id,omitempty"`
	Name string `column:"categories.name" json:"name"`
}

type BrandCategory struct {
	ID         string `column:"models.id" json:"_id,omitempty"`
	BrandID    string `column:"brands.id" json:"brandId"`
	CategoryID string `column:"categories.id" json:"categoryId"`
}

type Model struct {
	ID              string `column:"models.id" json:"_id,omitempty"`
	BrandCategoryID string `column:"brand_categories.id" json:"brandCategoryId"`
	Name            string `column:"models.name" json:"name"`
}

type Group struct {
	ID   string `column:"groups.id" json:"_id,omitempty"`
	Name string `column:"groups.name" json:"name"`
}

func AttributesAggregationBuilder(builder *sql_query.SelectBuilder, dto any, options AggregationOptions) {
	joinType := builder.Join
	if options.Preserve {
		joinType = builder.LeftJoin
	}

	joinType(options.TableName, fmt.Sprintf("%s.id = %s", options.TableName, options.LocalField))

	builder.SelectJSONAggregateFunc(func(b *sql_query.SelectBuilder) {
		b.SelectJSONAggregate(options.Alias, dto, options.UseArray)
	})
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
