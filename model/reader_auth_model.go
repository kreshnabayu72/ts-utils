package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReaderAuth struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyCode string             `json:"companyCode" bson:"companyCode"`
	ReaderId    string             `json:"readerId" bson:"readerId"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}
