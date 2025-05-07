package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterToken struct {
	ID              primitive.ObjectID `json:"_id"             bson:"_id"`
	Token           string             `json:"token"           bson:"token"`
	UserIsAdmin     bool               `json:"userIsAdmin"     bson:"userIsAdmin"`
	UserEmail       string             `json:"userEmail"       bson:"userEmail"`
	UserCompanyCode string             `json:"userCompanyCode" bson:"userCompanyCode"`
	UserPhoneNumber string             `json:"userPhoneNumber" bson:"userPhoneNumber"`
	CreatedAt       time.Time          `json:"createdAt"       bson:"createdAt"`
}
