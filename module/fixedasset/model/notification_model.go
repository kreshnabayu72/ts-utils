package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID         primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	User       primitive.ObjectID     `json:"user" bson:"user"`
	Title      string                 `json:"title"         bson:"title"`
	Detail     string                 `json:"detail"        bson:"detail"`
	IsRead     bool                   `json:"isRead"        bson:"isRead"`
	IsReadOnly bool                   `json:"isReadOnly"    bson:"isReadOnly"`
	IsManager  bool                   `json:"isManager"     bson:"isManager"`
	Module     string                 `json:"module"        bson:"module"`
	Status     string                 `json:"status"        bson:"status"`
	Severity   string                 `json:"severity"      bson:"severity"`
	Category   string                 `json:"category"      bson:"category"`
	Type       string                 `json:"type"          bson:"type"`
	Data       map[string]interface{} `json:"data"          bson:"data"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
