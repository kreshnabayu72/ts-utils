package pipeline

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	supply_model "gitea.qwertysystem.net/BETS/ts-utils/module/supply/model"
	"gitea.qwertysystem.net/BETS/ts-utils/pipeline"
)

// ApprovalAggregation adds the necessary aggregation to the existing pipeline builder
// for XXX approval list & options, exclusively used in supply asset
func ApprovalAggregation(
	builder *pipeline.PipelineBuilder,
	approvalFor string,
	loggedInUserId primitive.ObjectID,
	filterStatus []supply_model.ApprovalStatus,
) {
	builder = builder.
		Unwind(bson.M{
			"path": "$approvers",
		}).
		Match(bson.M{
			"for":            approvalFor,
			"isActivated":    true,
			"approvers.user": loggedInUserId,
			"$expr": bson.M{
				"$or": bson.A{
					bson.M{
						"$and": bson.A{
							bson.M{
								"$in": bson.A{
									supply_model.ApprovalStatusApproved,
									filterStatus,
								},
							},
							bson.M{
								"$eq": bson.A{
									"$approvers.action",
									supply_model.ApprovalActionApproved,
								},
							},
						},
					},
					bson.M{
						"$and": bson.A{
							bson.M{
								"$in": bson.A{
									supply_model.ApprovalStatusRejected,
									filterStatus,
								},
							},
							bson.M{
								"$eq": bson.A{
									"$approvers.action",
									supply_model.ApprovalActionRejected,
								},
							},
						},
					},
					bson.M{
						"$and": bson.A{
							bson.M{
								"$in": bson.A{
									supply_model.ApprovalStatusWaitingForApproval,
									filterStatus,
								},
							},
							bson.M{
								"$eq": bson.A{
									"$status",
									supply_model.ApprovalStatusWaitingForApproval,
								},
							},
							bson.M{
								"$eq": bson.A{
									"$approvers.action",
									primitive.Null{},
								},
							},
						},
					},
				},
			},
		}).
		Set(bson.M{
			// Set for root replacement.
			"object._id":      "$_id",
			"object.objectId": "$objectId",
			"object.status": bson.M{
				"$switch": bson.M{
					"branches": bson.A{
						bson.M{
							"case": bson.M{
								"$eq": bson.A{
									"$approvers.action",
									supply_model.ApprovalActionApproved,
								},
								"$then": supply_model.ApprovalStatusApproved,
							},
						},
						bson.M{
							"case": bson.M{
								"$eq": bson.A{
									"$approvers.action",
									supply_model.ApprovalActionRejected,
								},
								"$then": supply_model.ApprovalStatusRejected,
							},
						},
					},
					"default": supply_model.ApprovalStatusWaitingForApproval,
				},
			},
			"object.createdAt": bson.M{
				"$toDate": "$object.createdAt",
			},
		}).
		ReplaceRoot(bson.M{
			"newRoot": "$object",
		})
}
