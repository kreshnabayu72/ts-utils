package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CompanyCode string             `json:"companyCode" bson:"companyCode"`
	Email       string             `json:"email" bson:"email"`
}
