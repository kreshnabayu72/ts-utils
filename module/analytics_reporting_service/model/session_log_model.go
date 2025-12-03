package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionLog struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Activity       string             `json:"activity" bson:"activity"`
	Status         *string            `json:"status" bson:"status"`
	Note           string             `json:"note" bson:"note"`
	IPAddress      string             `json:"ipAddress" bson:"ipAddress"`
	UserID         string             `json:"userId" bson:"userId"`
	UserFullName   string             `json:"userFullName" bson:"userFullName"`
	UserEmployeeID *string            `json:"userEmployeeId" bson:"userEmployeeId"`
	UserPosition   *string            `json:"userPosition" bson:"userPosition"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}
