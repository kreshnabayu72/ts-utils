package sql_builder

import (
	"fmt"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	fixed_asset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

func UserAttributeAggregationBuilder(builder *sql_query.SelectBuilder, localField string, as string, preserve bool) {
	joinType := builder.Join
	if preserve {
		joinType = builder.LeftJoin
	}

	joinType(db.UserTableName, fmt.Sprintf("%s.id = %s", db.UserTableName, localField))

	type userJoinProjection struct {
		ID       string `json:"_id" column:"users.id"`
		FullName string `json:"fullName" column:"concat(users.first_name, ' ', users.last_name)"`
		Key      string `json:"key" column:"users.key"`
	}

	builder.SelectJSONAggregate(as, userJoinProjection{}, false)
}

func AttributesAggregationBuilder(builder *sql_query.SelectBuilder, dbName string, localField string, as string, preserve bool) {
	joinType := builder.Join
	if preserve {
		joinType = builder.LeftJoin
	}

	joinType(dbName, fmt.Sprintf("%s.id = %s", dbName, localField))

	builder.SelectJSONAggregate(as, map[string]string{
		"_id":      fmt.Sprintf("%s.id::text", dbName),
		"name":     fmt.Sprintf("%s.name", dbName),
		"key":      fmt.Sprintf("%s.key", dbName),
		"fullPath": fmt.Sprintf("%s.full_path", dbName),
	}, true)
}

func UserAttributesAggregationBuilder(builder *sql_query.SelectBuilder, localField string, as string, preserve bool) {
	joinType := builder.Join
	if preserve {
		joinType = builder.LeftJoin
	}

	joinType(db.UserTableName, fmt.Sprintf("%s.id = %s", db.UserTableName, localField))

	type userJoinProjection struct {
		ID       string `json:"_id" column:"users.id"`
		FullName string `json:"fullName" column:"concat(users.first_name, ' ', users.last_name)"`
		Key      string `json:"key" column:"users.key"`
	}

	builder.SelectJSONAggregate(as, userJoinProjection{}, true)
}

func GetAssetIdFromAliasCodeAggregationBuilder(builder *sql_query.SelectBuilder, tableName string) {
	res, args, err := sql_query.NewSQLSelectBuilder[fixed_asset_model.AliasCode](tableName).
		Join("alias_codes ac", "true").
		Join("groups g", "g.id = a.group_id").
		Join("model_types mt", "mt.id = a.model_id").
		Join("brand_categories bc", "bc.id = mt.brand_category_id").
		Join("categories c", "c.id = bc.category_id").
		Join("asset_names an", "an.id = a.name_id").
		LeftJoin("alias_codes_data acd", "acd.alias_code_id = ac.id").
		GroupBy("_id", "ac.separator").
		SetLimit(1).
		Build()
	if err != nil {
		log.Println(err)
	}

	builder.Columns = append(builder.Columns, fmt.Sprintf("%s.assetId AS assetId", res))
	builder.Args = append(builder.Args, args...)
}
