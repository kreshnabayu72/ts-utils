package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginSession struct {
	ID                   primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`
	UserCompanyCode      string             `json:"userCompanyCode"      bson:"userCompanyCode"`
	UserEmail            string             `json:"userEmail"            bson:"userEmail"`
	DeleteFromDatabaseAt time.Time          `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	SessionType          string             `json:"sessionType"          bson:"sessionType"` // Possible Values: "mobile", "web"
	CreatedAt            time.Time          `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt"            bson:"updatedAt"`
}
