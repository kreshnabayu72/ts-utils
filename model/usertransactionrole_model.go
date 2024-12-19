package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PARTIALLY COMPLETE
type UserTransactionRole struct {
	ID                        primitive.ObjectID   `json:"_id,omitempty"             bson:"_id,omitempty"`
	ManagedGroups             []primitive.ObjectID `json:"managedGroups,omitempty"   bson:"managedGroups,omitempty"`
	Group                     *Group               `json:"group,omitempty"           bson:"group,omitempty"`
	TransactionGroupAttribute string               `json:"transactionGroupAttribute" bson:"transactionGroupAttribute"`
}
