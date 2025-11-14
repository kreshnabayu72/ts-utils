package new_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContinuousImprovementLog struct {
	ID     primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Action string                 `json:"action" bson:"action"`
	User   string                 `json:"user" bson:"user"`
	Detail map[string]interface{} `json:"detail" bson:"detail"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
