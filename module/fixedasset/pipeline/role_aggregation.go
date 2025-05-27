package pipeline

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	fixedasset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/pipeline"
)

type TransactionRoleRestriction struct {
	ManageType []fixedasset_model.TransactionManageType `json:"manageType" bson:"manageType"`
}

type SystemRoleRestriction struct {
	RoleType    fixedasset_model.RoleType   `json:"roleType"    bson:"roleType"`
	Permissions fixedasset_model.Permission `json:"permissions" bson:"permissions"`
}

type RoleRestrictionOptions struct {
	TransactionRole *TransactionRoleRestriction `json:"transactionRole" bson:"transactionRole"`
	SystemRole      *SystemRoleRestriction      `json:"systemRole"      bson:"systemRole"`
}

func getTransactionRolePermission(
	transactionGroupAttribute fixedasset_model.TransactionGroupAttribute,
	roleType fixedasset_model.TransactionManageType,
) bson.M {
	return bson.M{
		"$gt": bson.A{
			bson.M{
				"$size": bson.M{
					"$filter": bson.M{
						"input": "$userTransactionRoles",
						"as":    "role",
						"cond": bson.M{
							"$and": bson.A{
								bson.M{
									"$eq": bson.A{
										"$$role.transactionGroupAttribute",
										transactionGroupAttribute,
									},
								},
								bson.M{
									"$eq": bson.A{
										"$$role.roleType",
										roleType,
									},
								},
							},
						},
					},
				},
			},
			0,
		},
	}
}

func getSystemRolePermission(roleType fixedasset_model.RoleType, permission string) bson.M {
	return bson.M{
		"$gt": bson.A{
			bson.M{
				"$size": bson.M{
					"$filter": bson.M{
						"input": "$userSystemRoles",
						"as":    "role",
						"cond": bson.M{
							"$eq": bson.A{
								fmt.Sprintf("$$role.managePermission.%s.%s", roleType, permission),
								true,
							},
						},
					},
				},
			},
			0,
		},
	}
}

// RolePermissionAggregation adds role permission aggregation to the existing pipeline builder
func RolePermissionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	groupField string,
	isFieldArray bool,
) {
	if groupField != "" && !strings.HasPrefix(groupField, "$") {
		groupField = fmt.Sprintf("$%s", groupField)
	}

	builder.
		Set(bson.M{
			"loggedInUserId": loggedInUserId,
		}).
		Lookup(bson.M{
			"from":         db.UserTransactionRoleModelName,
			"localField":   "loggedInUserId",
			"foreignField": "user._id",
			"let": bson.M{
				"groupField": groupField,
			},
			"as": "userTransactionRoles",
			"pipeline": pipeline.NewPipelineBuilder().
				OptionalMatch(
					bson.M{
						"$expr": bson.M{"$and": bson.A{
							bson.M{
								"$eq": bson.A{"$group._id", "$$groupField"},
							},
							bson.M{
								"$eq": bson.A{"$isActive", true},
							},
						}},
					},

					groupField != "" && !isFieldArray,
				).
				OptionalMatch(
					bson.M{
						"$expr": bson.M{"$and": bson.A{
							bson.M{
								"$in": bson.A{"$group._id", "$$groupField"},
							},
							bson.M{
								"$eq": bson.A{"$isActive", true},
							},
						}},
					},
					groupField != "" && isFieldArray,
				).
				Build(),
		}).
		Lookup(bson.M{
			"from":         db.UserSystemRoleModelName,
			"localField":   "loggedInUserId",
			"foreignField": "user._id",
			"let": bson.M{
				"groupField": groupField,
			},
			"as": "userSystemRoles",
			"pipeline": pipeline.NewPipelineBuilder().
				OptionalMatch(
					bson.M{
						"$expr": bson.M{"$and": bson.A{
							bson.M{
								"$eq": bson.A{"$group._id", "$$groupField"},
							},
							bson.M{
								"$eq": bson.A{"$isActive", true},
							},
						}},
					},
					groupField != "" && !isFieldArray,
				).
				OptionalMatch(
					bson.M{
						"$expr": bson.M{"$and": bson.A{
							bson.M{
								"$in": bson.A{"$group._id", "$$groupField"},
							},
							bson.M{
								"$eq": bson.A{"$isActive", true},
							},
						}},
					},
					groupField != "" && isFieldArray,
				).
				Build(),
		}).
		Set(
			bson.M{
				"permissions": bson.M{
					"transactionRoles": bson.M{
						"borrowing": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGABorrowingRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGABorrowingRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGABorrowingRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGABorrowingRole, fixedasset_model.TransactionManageTypeApprover),
						},
						"transfer": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGATransferRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGATransferRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGATransferRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGATransferRole, fixedasset_model.TransactionManageTypeApprover),
						},
						"disposal": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGADisposalRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGADisposalRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGADisposalRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGADisposalRole, fixedasset_model.TransactionManageTypeApprover),
						},
						"tracking": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGATrackingRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGATrackingRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGATrackingRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGATrackingRole, fixedasset_model.TransactionManageTypeApprover),
						},
						"repairTicketing": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGARepairRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGARepairRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGARepairRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGARepairRole, fixedasset_model.TransactionManageTypeApprover),
						},
						"maintenance": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGAMaintenanceRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGAMaintenanceRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGAMaintenanceRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGAMaintenanceRole, fixedasset_model.TransactionManageTypeApprover),
						},
						"audit": bson.M{
							"manager":    getTransactionRolePermission(fixedasset_model.TGAAuditRole, fixedasset_model.TransactionManageTypeManager),
							"monitoring": getTransactionRolePermission(fixedasset_model.TGAAuditRole, fixedasset_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(fixedasset_model.TGAAuditRole, fixedasset_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(fixedasset_model.TGAAuditRole, fixedasset_model.TransactionManageTypeApprover),
						},
					},
					"systemRoles": bson.M{
						"importData": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTImportData, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTImportData, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTImportData, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTImportData, "delete"),
						},
						"assetAttribute": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTAssetAttribute, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTAssetAttribute, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTAssetAttribute, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTAssetAttribute, "delete"),
						},
						"assetPolicies": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTAssetPolicies, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTAssetPolicies, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTAssetPolicies, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTAssetPolicies, "delete"),
						},
						"depreciationGroup": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTDepreciationGroup, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTDepreciationGroup, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTDepreciationGroup, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTDepreciationGroup, "delete"),
						},
						"depreciationMethod": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTDepreciationMethod, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTDepreciationMethod, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTDepreciationMethod, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTDepreciationMethod, "delete"),
						},
						"documentDeletion": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTDocumentDeletion, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTDocumentDeletion, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTDocumentDeletion, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTDocumentDeletion, "delete"),
						},
						"group": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTGroup, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTGroup, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTGroup, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTGroup, "delete"),
						},
						"iotReader": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTIotReader, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTIotReader, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTIotReader, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTIotReader, "delete"),
						},
						"purchaseDocument": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTPurchaseDocument, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTPurchaseDocument, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTPurchaseDocument, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTPurchaseDocument, "delete"),
						},
						"purchaseInformation": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTPurchaseInformation, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTPurchaseInformation, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTPurchaseInformation, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTPurchaseInformation, "delete"),
						},
						"userAndRole": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTUserAndRole, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTUserAndRole, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTUserAndRole, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTUserAndRole, "delete"),
						},
						"tag": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTTag, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTTag, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTTag, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTTag, "delete"),
						},
						"registerAsset": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTRegisterAsset, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTRegisterAsset, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTRegisterAsset, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTRegisterAsset, "delete"),
						},
						"tagType": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTTagType, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTTagType, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTTagType, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTTagType, "delete"),
						},
						"openApi": bson.M{
							"create": getSystemRolePermission(fixedasset_model.RTOpenApi, "create"),
							"view":   getSystemRolePermission(fixedasset_model.RTOpenApi, "view"),
							"update": getSystemRolePermission(fixedasset_model.RTOpenApi, "update"),
							"delete": getSystemRolePermission(fixedasset_model.RTOpenApi, "delete"),
						},
					},
				},
			},
		)
}

func FindUserSystemRoles(v any, ctx context.Context, companyCode string, loggedInUserId primitive.ObjectID) error {
	db.ConnectToDB(db.FixedAssetDBName.CompanyDBName(companyCode))
	coll := db.GetCollection(db.UserSystemRoleModelName)

	cursor, err := coll.Find(context.Background(), bson.M{"user._id": loggedInUserId, "isActive": true})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, v)
	if err != nil {
		return err
	}

	return nil
}

func GetHasSystemRoleSection(key string) bson.M {
	rt, ok := fixedasset_model.RoleTypeMap[key]
	if !ok {
		return bson.M{} // or panic/log error
	}

	return bson.M{
		"hasSystemRole": bson.M{
			key: bson.M{
				"create": getSystemRolePermission(rt, "create"),
				"view":   getSystemRolePermission(rt, "view"),
				"update": getSystemRolePermission(rt, "update"),
				"delete": getSystemRolePermission(rt, "delete"),
			},
		},
	}
}

// only system role (for settings, file manager, import, etc)
func SystemRolePermissionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	groupField string,
	isFieldArray bool,
	roleType string,
) {
	if groupField != "" && !strings.HasPrefix(groupField, "$") {
		groupField = fmt.Sprintf("$%s", groupField)
	}

	builder.Set(bson.M{
		"loggedInUserId": loggedInUserId,
	}).
		Lookup(bson.M{
			"from":         db.UserSystemRoleModelName,
			"localField":   "loggedInUserId",
			"foreignField": "user._id",
			"let": bson.M{
				"groupField": groupField,
			},
			"as": "userSystemRoles",
			"pipeline": pipeline.NewPipelineBuilder().
				Match(bson.M{"isActive": true, "type": roleType}).
				OptionalMatch(
					bson.M{
						"$expr": bson.M{
							"$and": bson.A{
								bson.M{"$in": bson.A{"$$groupField", "$group._id"}},
							},
						},
					},
					"group._id" != "" && !isFieldArray,
				).
				OptionalMatch(
					bson.M{
						"$expr": bson.M{
							"$gt": bson.A{
								bson.M{"$size": bson.M{"$setIntersection": bson.A{"$$groupField", "$group._id"}}},
								0,
							},
						},
					},
					"group._id" != "" && isFieldArray,
				).
				Build(),
		}).
		Addfields(
			GetHasSystemRoleSection(roleType),
		)
}

func GetHasTransactionRole(transactionGroupAttribute fixedasset_model.TransactionGroupAttribute) bson.M {
	return bson.M{
		"isManager":    getTransactionRolePermission(transactionGroupAttribute, fixedasset_model.TransactionManageTypeManager),
		"isMonitoring": getTransactionRolePermission(transactionGroupAttribute, fixedasset_model.TransactionManageTypeMonitoring),
	}
}

func TransactionRolePermissionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	groupField string,
	transactionGroupAttribute fixedasset_model.TransactionGroupAttribute,
	filterOptions RoleRestrictionOptions, //for filtering user_transaction_roles
	filterOutEmptyRole ...bool,
) {
	if groupField != "" && !strings.HasPrefix(groupField, "$") {
		groupField = fmt.Sprintf("$%s", groupField)
	}

	filter := true

	if len(filterOutEmptyRole) > 0 {
		filter = filterOutEmptyRole[0]
	}

	builder.
		Set(
			bson.M{
				"loggedInUserId": loggedInUserId,
			},
		).
		Lookup(
			bson.M{
				"from":         db.UserTransactionRoleModelName,
				"localField":   "loggedInUserId",
				"foreignField": "user._id",
				"let": bson.M{
					"groupField": groupField,
				},
				"as": "userTransactionRoles",
				"pipeline": pipeline.NewPipelineBuilder().
					OptionalMatch(
						bson.M{
							"$expr": bson.M{"$and": bson.A{
								bson.M{"$eq": bson.A{"$group._id", "$$groupField"}},
								bson.M{"$eq": bson.A{"$isActive", true}},
							}},
						},
						groupField != "",
					).
					Match(
						bson.M{
							"roleType": bson.M{
								"$in": filterOptions.TransactionRole.ManageType,
							},
							"transactionGroupAttribute": transactionGroupAttribute,
						},
					).
					Build(),
			},
		).
		Match(bson.M{
			"$expr": bson.M{
				"$and": bson.A{
					bson.M{
						"$gt": bson.A{
							bson.M{
								"$size": "$userTransactionRoles",
							},
							0,
						},
					},
					filter,
				},
			},
		}).
		Addfields(GetHasTransactionRole(transactionGroupAttribute))
}

// Example: Transfer needs both origin & destination group role

type GroupFieldAndAttributeName struct {
	Field         string // originGroup._id / destinationGroup._id
	AttributeName string // Origin / Destination (For isDestinationManager)
}

func TransactionRoleMultipleGroupAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	groupFieldsAndAttribute []GroupFieldAndAttributeName,
	transactionGroupAttribute fixedasset_model.TransactionGroupAttribute,
	filterOptions RoleRestrictionOptions, //for filtering user_transaction_roles
) {
	checkedGroupFields := make([]GroupFieldAndAttributeName, 0)

	for _, groupField := range groupFieldsAndAttribute {
		checkedField := groupField.Field

		if checkedField != "" && !strings.HasPrefix(checkedField, "$") {
			checkedField = fmt.Sprintf("$%s", checkedField)
		}

		checkedGroupFields = append(checkedGroupFields, GroupFieldAndAttributeName{
			Field:         checkedField,
			AttributeName: groupField.AttributeName,
		})
	}

	builder.
		Set(
			bson.M{
				"loggedInUserId": loggedInUserId,
			},
		)

	for _, groupFieldAndAttribute := range checkedGroupFields {
		attributeField := strings.ReplaceAll(groupFieldAndAttribute.Field, "$", "")
		attributeField = strings.ReplaceAll(attributeField, ".", "")

		builder.Lookup(
			bson.M{
				"from":         db.UserTransactionRoleModelName,
				"localField":   "loggedInUserId",
				"foreignField": "user._id",
				"let": bson.M{
					attributeField: groupFieldAndAttribute.Field,
				},
				"as": "userTransactionRoles",
				"pipeline": pipeline.NewPipelineBuilder().
					Match(
						bson.M{
							"$expr": bson.M{"$and": bson.A{
								bson.M{"$eq": bson.A{"$group._id", fmt.Sprintf("$$%s", attributeField)}},
								bson.M{"$eq": bson.A{"$isActive", true}},
							}},
						},
					).
					Match(
						bson.M{
							"roleType": bson.M{
								"$in": filterOptions.TransactionRole.ManageType,
							},
							"transactionGroupAttribute": transactionGroupAttribute,
						},
					).
					Build(),
			},
		).
			Addfields(
				bson.M{
					fmt.Sprintf("is%sManager", groupFieldAndAttribute.AttributeName):    getTransactionRolePermission(transactionGroupAttribute, fixedasset_model.TransactionManageTypeManager),
					fmt.Sprintf("is%sMonitoring", groupFieldAndAttribute.AttributeName): getTransactionRolePermission(transactionGroupAttribute, fixedasset_model.TransactionManageTypeMonitoring),
				},
			)
	}
}

// When this function is made, this is intender only for asset list
func TransactionRoleMultiPermissionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	groupField string,
) {
	if groupField != "" && !strings.HasPrefix(groupField, "$") {
		groupField = fmt.Sprintf("$%s", groupField)
	}

	builder.
		Set(
			bson.M{
				"loggedInUserId": loggedInUserId,
			},
		).
		Lookup(
			bson.M{
				"from":         db.UserTransactionRoleModelName,
				"localField":   "loggedInUserId",
				"foreignField": "user._id",
				"let": bson.M{
					"groupField": groupField,
				},
				"as": "userTransactionRoles",
				"pipeline": pipeline.NewPipelineBuilder().
					OptionalMatch(
						bson.M{
							"$expr": bson.M{"$and": bson.A{
								bson.M{"$eq": bson.A{"$group._id", "$$groupField"}},
								bson.M{"$eq": bson.A{"$isActive", true}},
							}},
						},
						groupField != "",
					).
					Match(
						bson.M{
							"roleType": bson.M{
								"$in": bson.A{fixedasset_model.TRTManager, fixedasset_model.TRTMonitoring},
							},
							"transactionGroupAttribute": bson.M{
								"$in": bson.A{fixedasset_model.TGABorrowingRole, fixedasset_model.TGATransferRole, fixedasset_model.TGADisposalRole},
							},
						},
					).
					Build(),
			},
		).
		Addfields(
			bson.M{
				"manager": bson.M{
					"transfer":  getTransactionRolePermission(fixedasset_model.TGATransferRole, fixedasset_model.TransactionManageTypeManager),
					"borrowing": getTransactionRolePermission(fixedasset_model.TGABorrowingRole, fixedasset_model.TransactionManageTypeManager),
					"disposal":  getTransactionRolePermission(fixedasset_model.TGADisposalRole, fixedasset_model.TransactionManageTypeManager),
				},
				"monitoring": bson.M{
					"transfer":  getTransactionRolePermission(fixedasset_model.TGATransferRole, fixedasset_model.TransactionManageTypeMonitoring),
					"borrowing": getTransactionRolePermission(fixedasset_model.TGABorrowingRole, fixedasset_model.TransactionManageTypeMonitoring),
					"disposal":  getTransactionRolePermission(fixedasset_model.TGADisposalRole, fixedasset_model.TransactionManageTypeMonitoring),
				},
			},
		)
}

func RoleRestrictionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	groupField string,
	filterOptions RoleRestrictionOptions,
) {
	if groupField != "" && !strings.HasPrefix(groupField, "$") {
		groupField = fmt.Sprintf("$%s", groupField)
	}

	builder.
		Set(
			bson.M{
				"loggedInUserId": loggedInUserId,
			},
		)

	if filterOptions.TransactionRole != nil {
		builder.
			Lookup(
				bson.M{
					"from":         db.UserTransactionRoleModelName,
					"localField":   "loggedInUserId",
					"foreignField": "user._id",
					"let": bson.M{
						"groupField": groupField,
					},
					"as": "userTransactionRoles",
					"pipeline": pipeline.NewPipelineBuilder().
						OptionalMatch(
							bson.M{
								"$expr": bson.M{"$and": bson.A{
									bson.M{"$eq": bson.A{"$group._id", "$$groupField"}},
									bson.M{"$eq": bson.A{"$isActive", true}},
								}},
							},
							groupField != "",
						).
						Match(
							bson.M{
								"manageType": bson.M{
									"$in": filterOptions.TransactionRole.ManageType,
								},
							},
						).
						Build(),
				},
			).
			Match(
				bson.M{
					"$expr": bson.M{
						"$gt": bson.A{
							bson.M{
								"$size": "$userTransactionRoles",
							},
							0,
						},
					},
				},
			)
	}

	if filterOptions.SystemRole != nil {
		builder.
			Lookup(
				bson.M{
					"from":         db.UserSystemRoleModelName,
					"localField":   "loggedInUserId",
					"foreignField": "user._id",
					"let": bson.M{
						"groupField": groupField,
					},
					"as": "userSystemRoles",
					"pipeline": pipeline.NewPipelineBuilder().
						OptionalMatch(
							bson.M{
								"$expr": bson.M{"$and": bson.A{
									bson.M{"$eq": bson.A{"$group._id", "$$groupField"}},
									bson.M{"$eq": bson.A{"$isActive", true}},
								}},
							},
							groupField != "",
						).
						Match(
							bson.M{
								"$and": bson.A{
									bson.M{
										"type": filterOptions.SystemRole.RoleType,
									},
									bson.M{
										"$or": bson.A{
											bson.M{
												fmt.Sprintf("managePermission.%s.create", filterOptions.SystemRole.RoleType): filterOptions.SystemRole.Permissions.Create,
											},
											bson.M{
												fmt.Sprintf("managePermission.%s.view", filterOptions.SystemRole.RoleType): filterOptions.SystemRole.Permissions.View,
											},
											bson.M{
												fmt.Sprintf("managePermission.%s.update", filterOptions.SystemRole.RoleType): filterOptions.SystemRole.Permissions.Update,
											},
											bson.M{
												fmt.Sprintf("managePermission.%s.delete", filterOptions.SystemRole.RoleType): filterOptions.SystemRole.Permissions.Delete,
											},
										},
									},
								},
							},
						).
						Build(),
				},
			).
			Match(
				bson.M{
					"$expr": bson.M{
						"$gt": bson.A{
							bson.M{
								"$size": "$userSystemRoles",
							},
							0,
						},
					},
				},
			)
	}
}
