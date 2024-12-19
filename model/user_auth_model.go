package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID          primitive.ObjectID `json:"id,omitempty"         bson:"_id,omitempty"`
	Email       string             `json:"email"                  bson:"email"`
	Password    string             `json:"password"            bson:"password"`
	IsAdmin     bool               `json:"isAdmin"             bson:"isAdmin"`
	CompanyCode string             `json:"companyCode"            bson:"companyCode"`
	UserDetail  primitive.ObjectID `json:"userDetail"         bson:"userDetail"`
	CreatedAt   time.Time          `json:"createdAt"            bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"            bson:"updatedAt"`
}
