package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRFID struct {
	Code      string `json:"code"      bson:"code"`
	ShortCode string `json:"shortCode" bson:"shortCode"`
}

type User struct {
	ID                   primitive.ObjectID `json:"id,omitempty"         bson:"_id,omitempty"`
	Key                  int                `json:"key"                  bson:"key"`
	FirstName            string             `json:"firstName"            bson:"firstName"`
	LastName             string             `json:"lastName"             bson:"lastName"`
	Division             string             `json:"division"             bson:"division"`
	Position             *string            `json:"position"             bson:"position"`
	EmployeeID           *string            `json:"employeeId"           bson:"employeeId"`
	PhoneNumber          string             `json:"phoneNumber"          bson:"phoneNumber"`
	Email                string             `json:"email"                bson:"email"`
	RFID                 *UserRFID          `json:"rfid"                 bson:"rfid"`
	IsActive             bool               `json:"isActive"             bson:"isActive"`
	IsTemporary          bool               `json:"isTemporary"          bson:"isTemporary"`
	ExpiryDate           *time.Time         `json:"expiryDate"           bson:"expiryDate"`
	PendingEmailChange   string             `json:"pendingEmailChange"   bson:"pendingEmailChange"`
	ModifiedBy           primitive.ObjectID `json:"modifiedBy"           bson:"modifiedBy"`
	ProfilePictureSmall  *string            `json:"profilePictureSmall"  bson:"profilePictureSmall"`
	ProfilePictureMedium *string            `json:"profilePictureMedium" bson:"profilePictureMedium"`
	ProfilePictureBig    *string            `json:"profilePictureBig"    bson:"profilePictureBig"`
	IsConfirmed          bool               `json:"isConfirmed"          bson:"isConfirmed"`
	IsDefault            bool               `json:"isDefault"            bson:"isDefault"`
	CreatedAt            time.Time          `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt"            bson:"updatedAt"`
}
