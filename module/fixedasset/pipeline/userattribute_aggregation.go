package pipeline

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
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
