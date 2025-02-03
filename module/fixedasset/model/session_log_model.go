package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionLog struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Activity   string             `json:"activity"      bson:"activity"`
	User       string             `json:"user"          bson:"user"`
	UserId     primitive.ObjectID `json:"userId"        bson:"userId"`
	Email      string             `json:"email"         bson:"email"`
	EmployeeId string             `json:"employeeId"    bson:"employeeId"`
	Position   string             `json:"position"      bson:"position"`
	IpAddress  string             `json:"ipAddress"     bson:"ipAddress"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
