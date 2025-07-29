package pipeline

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	fixed_asset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	"gitea.qwertysystem.net/BETS/ts-utils/pipeline"
)

// Lookup and unwind user attribute + adding fullName, so user attributes are up to date to db
// Needed since edit user is moved to global-settings
func LookupUserAttributeAggregation(
	builder *pipeline.PipelineBuilder,
	localField string,
	as string,
	preserveUserLookup bool,
) {
	builder.
		Lookup(bson.M{
			"from":         db.UserModelName,
			"localField":   localField,
			"foreignField": "_id",
			"as":           as,
			"pipeline": pipeline.NewPipelineBuilder().
				Project(bson.M{
					"_id":      1,
					"fullName": bson.M{"$concat": bson.A{"$firstName", " ", "$lastName"}},
					"key":      1,
				}).
				Build(),
		}).Unwind(bson.M{"path": fmt.Sprintf("$%s", as), "preserveNullAndEmptyArrays": preserveUserLookup})
}

// Lookup and unwind attributes with _id name key
func LookupAttributesAggregation(
	builder *pipeline.PipelineBuilder,
	modelName string,
	localField string,
	as string,
	preserveNullAndEmptyArrays bool,
) {
	builder.
		Lookup(bson.M{
			"from":         modelName,
			"localField":   localField,
			"foreignField": "_id",
			"as":           as,
			"pipeline": pipeline.NewPipelineBuilder().
				Project(bson.M{
					"_id":      1,
					"name":     1,
					"key":      1,
					"fullPath": 1,
					"code":     1,
				}).
				Build(),
		}).Unwind(bson.M{"path": fmt.Sprintf("$%s", as), "preserveNullAndEmptyArrays": preserveNullAndEmptyArrays})
}

// Lookup and unwind user attributes with _id fullName key
func LookupUserAttributesAggregation(
	builder *pipeline.PipelineBuilder,
	localField string,
	as string,
	preserveNullAndEmptyArrays bool,
) {
	builder.
		Lookup(bson.M{
			"from":         db.UserModelName,
			"localField":   localField,
			"foreignField": "_id",
			"as":           as,
			"pipeline": pipeline.NewPipelineBuilder().
				Project(bson.M{
					"_id":      1,
					"fullName": bson.M{"$concat": bson.A{"$firstName", " ", "$lastName"}},
					"key":      1,
				}).
				Build(),
		}).Unwind(bson.M{"path": fmt.Sprintf("$%s", as), "preserveNullAndEmptyArrays": preserveNullAndEmptyArrays})
}

// Used for "assets" collection aggregation
func GetAssetIdFromAliasCode(
	builder *pipeline.PipelineBuilder,
	aliasCode *fixed_asset_model.AliasCode,
) {
	builder.
		Addfields(bson.M{"aliasCode": aliasCode}).
		Addfields(bson.M{
			"aliasCodeData": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$aliasCode", nil}},
					bson.A{},
					bson.M{
						"$filter": bson.M{
							"input": "$aliasCode.data",
							"as":    "data",
							"cond": bson.M{
								"$eq": bson.A{"$$data.isChecked", true},
							},
						},
					},
				},
			},
			"purchaseDateMonth": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$purchase.date", nil}},
					"00",
					bson.M{
						"$concat": bson.A{
							"0",
							bson.M{"$toString": bson.M{"$month": "$purchase.date"}},
						},
					},
				},
			},
			"purchaseDateYear": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$purchase.date", nil}},
					"00",
					bson.M{
						"$substr": bson.A{
							bson.M{"$toString": bson.M{"$year": "$purchase.date"}},
							2,
							2,
						},
					},
				},
			},
			"increment": bson.M{
				"$arrayElemAt": bson.A{
					bson.M{"$split": bson.A{"$name.nameWithSequence", " - "}},
					1,
				},
			},
		},
		).
		Addfields(bson.M{
			"incrementCode": bson.M{
				"$concat": bson.A{"0000", "$increment"},
			},
		}).
		Addfields(bson.M{
			"purchaseDateCode": bson.M{
				"$concat": bson.A{
					bson.M{
						"$cond": bson.A{
							bson.M{"$eq": bson.A{"$purchaseDateMonth", nil}},
							"00",
							bson.M{
								"$substrCP": bson.A{
									"$purchaseDateMonth",
									bson.M{
										"$subtract": bson.A{
											bson.M{"$strLenCP": "$purchaseDateMonth"},
											2,
										},
									},
									2,
								},
							},
						},
					},
					"$purchaseDateYear",
				},
			},
			"incrementCode": bson.M{
				"$substrCP": bson.A{
					"$incrementCode",
					bson.M{
						"$subtract": bson.A{
							bson.M{"$strLenCP": "$incrementCode"},
							4,
						},
					},
					4,
				},
			},
		},
		).
		Addfields(bson.M{
			"assetIdMap": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$aliasCode", nil}},
					bson.A{},
					bson.M{
						"$map": bson.M{
							"input": "$aliasCodeData",
							"as":    "data",
							"in": bson.M{
								"$switch": bson.M{
									"branches": bson.A{
										bson.M{
											"case": bson.M{"$eq": bson.A{"$$data.name", "Initial"}},
											"then": "$aliasCode.companyInitial",
										},
										bson.M{
											"case": bson.M{"$eq": bson.A{"$$data.name", "Group"}},
											"then": bson.M{"$ifNull": bson.A{"$group.code", "0000"}},
										},
										bson.M{
											"case": bson.M{"$eq": bson.A{"$$data.name", "Category"}},
											"then": bson.M{"$ifNull": bson.A{"$category.code", "0000"}},
										},
										bson.M{
											"case": bson.M{"$eq": bson.A{"$$data.name", "Asset Name"}},
											"then": bson.M{"$ifNull": bson.A{"$name.nameCode", "0000"}},
										},
										bson.M{
											"case": bson.M{"$eq": bson.A{"$$data.name", "Purchase Date"}},
											"then": "$purchaseDateCode",
										},
									},
									"default": "$incrementCode",
								},
							},
						},
					},
				},
			},
		},
		).
		Addfields(bson.M{
			"previousAssetId": "$assetId",
			"assetId": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{bson.M{"$size": "$assetIdMap"}, 0}},
					"$assetId",
					bson.M{
						"$reduce": bson.M{
							"input":        "$assetIdMap",
							"initialValue": "",
							"in": bson.M{
								"$cond": bson.A{
									bson.M{"$eq": bson.A{"$$value", ""}},
									"$$this",
									bson.M{
										"$concat": bson.A{"$$value", "$aliasCode.separator", "$$this"},
									},
								},
							},
						},
					},
				},
			},
			"name.aliasCode": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$aliasCode", nil}},
					"$name.aliasCode",
					"$assetId",
				},
			},
		})
}
