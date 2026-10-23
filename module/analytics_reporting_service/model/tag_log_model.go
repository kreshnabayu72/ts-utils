package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TagLog records one append-only TAG Log event of a TAG.
type TagLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	// TagID references the tenant tags.id.
	TagID       string              `json:"tagId" bson:"tagId"`
	Event       string              `json:"event" bson:"event"`
	Description []TagLogDescription `json:"description" bson:"description"`

	// CreatedBy* are empty for system events; names are fixed at log time.
	CreatedByID                *string `json:"createdById" bson:"createdById"`
	CreatedBySnapshot          *string `json:"createdBySnapshot" bson:"createdBySnapshot"`
	CreatedByFirstNameSnapshot *string `json:"createdByFirstNameSnapshot" bson:"createdByFirstNameSnapshot"`
	CreatedByLastNameSnapshot  *string `json:"createdByLastNameSnapshot" bson:"createdByLastNameSnapshot"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// TagLogDescription is one label/value row in a TAG Log entry.
type TagLogDescription struct {
	Label string `json:"label" bson:"label"`
	Value string `json:"value" bson:"value"`
}
