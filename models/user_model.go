package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEmbedded struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FullName *string             `json:"fullName" bson:"fullName"`
	Key      *int32              `json:"key" bson:"key"`
}
type User struct {
	ID                   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName            string             `json:"firstName" bson:"firstName"`
	LastName             string             `json:"lastName" bson:"lastName"`
	Key                  int32              `json:"key" bson:"key"`
	ProfilePictureBig    string             `json:"profilePictureBig" bson:"profilePictureBig"`
	ProfilePictureMedium string             `json:"profilePictureMedium" bson:"profilePictureMedium"`
	ProfilePictureSmall  string             `json:"profilePictureSmall" bson:"profilePictureSmall"`
	CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt" bson:"updatedAt"`
}
