package pipeline

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	supply_model "gitea.qwertysystem.net/BETS/ts-utils/module/supply/model"
	"gitea.qwertysystem.net/BETS/ts-utils/pipeline"
)

type TransactionRoleRestriction struct {
	RoleType   supply_model.TransactionRoleType     `json:"roleType"   bson:"roleType"`
	ManageType []supply_model.TransactionManageType `json:"manageType" bson:"manageType"`
}

type SystemRoleRestriction struct {
	RoleType    supply_model.SystemRoleType       `json:"roleType"    bson:"roleType"`
	Permissions supply_model.SystemRolePermission `json:"permissions" bson:"permissions"`
}

type RoleRestrictionOptions struct {
	TransactionRole *TransactionRoleRestriction `json:"transactionRole" bson:"transactionRole"`
	SystemRole      *SystemRoleRestriction      `json:"systemRole"      bson:"systemRole"`
}

func getTransactionRolePermission(
	roleType supply_model.TransactionRoleType,
	manageType supply_model.TransactionManageType,
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
										"$$role.roleType",
										roleType,
									},
								},
								bson.M{
									"$eq": bson.A{
										"$$role.manageType",
										manageType,
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

func getSystemRolePermission(roleType supply_model.SystemRoleType, permission string) bson.M {
	return bson.M{
		"$gt": bson.A{
			bson.M{
				"$size": bson.M{
					"$filter": bson.M{
						"input": "$userSystemRoles",
						"as":    "role",
						"cond": bson.M{
							"$and": bson.A{
								bson.M{
									"$eq": bson.A{
										"$$role.roleType",
										roleType,
									},
								},
								bson.M{
									"$eq": bson.A{
										fmt.Sprintf("$$role.permissions.%s", permission),
										true,
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

// RolePermissionAggregation adds role permission aggregation to the existing pipeline builder
func RolePermissionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	administrativeGroupField string,
) {
	if administrativeGroupField != "" && !strings.HasPrefix(administrativeGroupField, "$") {
		administrativeGroupField = fmt.Sprintf("$%s", administrativeGroupField)
	}

	builder.
		Set(bson.M{
			"loggedInUserId": loggedInUserId,
		}).
		Lookup(bson.M{
			"from":         db.UserTransactionRoleModelName,
			"localField":   "loggedInUserId",
			"foreignField": "user",
			"let": bson.M{
				"administrativeGroupField": administrativeGroupField,
			},
			"as": "userTransactionRoles",
			"pipeline": pipeline.NewPipelineBuilder().
				OptionalMatch(
					bson.M{
						"$expr": bson.M{
							"$eq": bson.A{"$administrativeGroup", "$$administrativeGroupField"},
						},
					},
					administrativeGroupField != "",
				).
				Build(),
		}).
		Lookup(bson.M{
			"from":         db.UserSystemRoleModelName,
			"localField":   "loggedInUserId",
			"foreignField": "user",
			"let": bson.M{
				"administrativeGroupField": administrativeGroupField,
			},
			"as": "userSystemRoles",
			"pipeline": pipeline.NewPipelineBuilder().
				Match(bson.M{
					"$expr": bson.M{
						"$eq": bson.A{"$administrativeGroup", "$$administrativeGroupField"},
					},
				}).
				Build(),
		}).
		Set(
			bson.M{
				"permissions": bson.M{
					"transactionRoles": bson.M{
						"receptionSupplierReturn": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeReceptionSupplierReturn, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeReceptionSupplierReturn, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeReceptionSupplierReturn, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeReceptionSupplierReturn, supply_model.TransactionManageTypeApprover),
						},
						"checkoutInternalReturn": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeCheckoutInternalReturn, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeCheckoutInternalReturn, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeCheckoutInternalReturn, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeCheckoutInternalReturn, supply_model.TransactionManageTypeApprover),
						},
						"transfer": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeTransfer, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeTransfer, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeTransfer, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeTransfer, supply_model.TransactionManageTypeApprover),
						},
						"defect": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeDefect, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeDefect, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeDefect, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeDefect, supply_model.TransactionManageTypeApprover),
						},
						"expired": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeExpired, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeExpired, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeExpired, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeExpired, supply_model.TransactionManageTypeApprover),
						},
						"tracking": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeTracking, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeTracking, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeTracking, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeTracking, supply_model.TransactionManageTypeApprover),
						},
						"missing": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeMissing, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeMissing, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeMissing, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeMissing, supply_model.TransactionManageTypeApprover),
						},
						"disposal": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeDisposal, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeDisposal, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeDisposal, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeDisposal, supply_model.TransactionManageTypeApprover),
						},
						"stockOpname": bson.M{
							"admin":      getTransactionRolePermission(supply_model.TransactionRoleTypeStockOpname, supply_model.TransactionManageTypeAdmin),
							"monitoring": getTransactionRolePermission(supply_model.TransactionRoleTypeStockOpname, supply_model.TransactionManageTypeMonitoring),
							"staff":      getTransactionRolePermission(supply_model.TransactionRoleTypeStockOpname, supply_model.TransactionManageTypeStaff),
							"approver":   getTransactionRolePermission(supply_model.TransactionRoleTypeStockOpname, supply_model.TransactionManageTypeApprover),
						},
					},
					"systemRoles": bson.M{
						"importData": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTypeImportData, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTypeImportData, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTypeImportData, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTypeImportData, "delete"),
						},
						"itemAttribute": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTypeItemAttribute, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTypeItemAttribute, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTypeItemAttribute, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTypeItemAttribute, "delete"),
						},
						"itemPolicy": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTypeItemPolicy, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTypeItemPolicy, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTypeItemPolicy, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTypeItemPolicy, "delete"),
						},
						"openApi": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTypeOpenAPI, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTypeOpenAPI, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTypeOpenAPI, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTypeOpenAPI, "delete"),
						},
						"documentDeletion": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTypeDocumentDeletion, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTypeDocumentDeletion, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTypeDocumentDeletion, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTypeDocumentDeletion, "delete"),
						},
						"group": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTypeGroup, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTypeGroup, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTypeGroup, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTypeGroup, "delete"),
						},
						"userRole": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleUserRole, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleUserRole, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleUserRole, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleUserRole, "delete"),
						},
						"iotReader": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleIOTReader, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleIOTReader, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleIOTReader, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleIOTReader, "delete"),
						},
						"license": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleLicense, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleLicense, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleLicense, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleLicense, "delete"),
						},
						"tag": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleTAG, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleTAG, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleTAG, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleTAG, "delete"),
						},
						"itemStock": bson.M{
							"create": getSystemRolePermission(supply_model.SystemRoleItemStock, "create"),
							"view":   getSystemRolePermission(supply_model.SystemRoleItemStock, "view"),
							"update": getSystemRolePermission(supply_model.SystemRoleItemStock, "update"),
							"delete": getSystemRolePermission(supply_model.SystemRoleItemStock, "delete"),
						},
					},
				},
			},
		)
}

func RoleRestrictionAggregation(
	builder *pipeline.PipelineBuilder,
	loggedInUserId primitive.ObjectID,
	administrativeGroupField string,
	filterOptions RoleRestrictionOptions,
) {
	if administrativeGroupField != "" && !strings.HasPrefix(administrativeGroupField, "$") {
		administrativeGroupField = fmt.Sprintf("$%s", administrativeGroupField)
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
					"foreignField": "user",
					"let": bson.M{
						"administrativeGroupField": administrativeGroupField,
					},
					"as": "userTransactionRoles",
					"pipeline": pipeline.NewPipelineBuilder().
						OptionalMatch(
							bson.M{
								"$expr": bson.M{
									"$eq": bson.A{"$administrativeGroup", "$$administrativeGroupField"},
								},
							},
							administrativeGroupField != "",
						).
						Match(
							bson.M{
								"roleType": filterOptions.TransactionRole.RoleType,
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
					"foreignField": "user",
					"let": bson.M{
						"administrativeGroupField": administrativeGroupField,
					},
					"as": "userSystemRoles",
					"pipeline": pipeline.NewPipelineBuilder().
						OptionalMatch(
							bson.M{
								"$expr": bson.M{
									"$eq": bson.A{"$administrativeGroup", "$$administrativeGroupField"},
								},
							},
							administrativeGroupField != "",
						).
						Match(
							bson.M{
								"$and": bson.A{
									bson.M{
										"roleType": filterOptions.SystemRole.RoleType,
									},
									bson.M{
										"$or": bson.A{
											bson.M{
												"permissions.create": filterOptions.SystemRole.Permissions.Create,
											},
											bson.M{
												"permissions.view": filterOptions.SystemRole.Permissions.View,
											},
											bson.M{
												"permissions.update": filterOptions.SystemRole.Permissions.Update,
											},
											bson.M{
												"permissions.delete": filterOptions.SystemRole.Permissions.Delete,
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
