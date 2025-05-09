package pipeline

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	Already used by old repos

	Used to get isManager isMonitoring etc inside aggregation
 	Used after addFields or lookup to user_transaction_roles
	Filter the transactionRoles array by isActive:true and isManagedByParent:{$ne:true} before using this

	example: AddFields(bson.M{"isManager":CheckHasTransactionRoleAggregation("$transactionRoles","transferRole","Manager","$assetGroup._id")})
*/

func CheckHasTransactionRoleAggregation(transactionRoleAttribute string, transactionType string, roleType string, groupIdAttribute string) primitive.M {
	return bson.M{
		"$anyElementTrue": bson.M{
			"$filter": bson.M{
				"input": transactionRoleAttribute,
				"as":    "role",
				"cond": bson.M{
					"$and": bson.A{
						bson.M{"$eq": bson.A{"$$role.transactionGroupAttribute", transactionType}},
						bson.M{"$eq": bson.A{"$$role.roleType", roleType}},
						bson.M{"$or": bson.A{
							bson.M{"$eq": bson.A{groupIdAttribute, "$$role.group._id"}},
							bson.M{"$in": bson.A{groupIdAttribute, "$$role.managedGroups"}},
						}},
					},
				},
			},
		},
	}
}
