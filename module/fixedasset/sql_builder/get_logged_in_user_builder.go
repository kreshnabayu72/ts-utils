package sql_builder

import (
	util_db "gitea.qwertysystem.net/BETS/ts-utils/v2/db"
	util_fixedasset_new_model "gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"
	util_sql_query "gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
)

func GetLoggedInUserBuilder(userId string) (string, []interface{}, error) {
	return util_sql_query.NewSQLSelectBuilder[util_fixedasset_new_model.UserAttribute](util_db.UserTableName).
		Where(map[string]util_sql_query.SQLCondition{
			"id": {
				Operator: util_sql_query.SQLOperatorEqual,
				Value:    userId,
			},
			"is_active": {
				Operator: util_sql_query.SQLOperatorEqual,
				Value:    true,
			},
			"is_deleted": {
				Operator: util_sql_query.SQLOperatorEqual,
				Value:    false,
			},
		}).
		Build()
}
