package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Email       string             `json:"email" bson:"email"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
	CompanyCode string             `json:"companyCode" bson:"companyCode"`
	Password    string             `json:"password" bson:"password"`
	IsDeleted   bool               `json:"isDeleted" bson:"isDeleted"`
	CreatedAt   time.Time          `json:"createdAt"     bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"     bson:"updatedAt"`
}
