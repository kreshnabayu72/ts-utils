package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEmbeddedRole struct {
	IsActive   bool `json:"isActive"   bson:"isActive"`
	IsAssigned bool `json:"isAssigned" bson:"isAssigned"`
}

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty"         bson:"_id,omitempty"`
	Key         int                `json:"key"                  bson:"key"`
	FirstName   string             `json:"firstName"            bson:"firstName"`
	LastName    string             `json:"lastName"             bson:"lastName"`
	Division    string             `json:"division"             bson:"division"`
	Position    *string            `json:"position"             bson:"position"`
	EmployeeID  *string            `json:"employeeId"           bson:"employeeId"`
	PhoneNumber string             `json:"phoneNumber"          bson:"phoneNumber"`
	Email       string             `json:"email"                bson:"email"`
	Password    *string            `json:"password"             bson:"password"`
	RFID        *struct {
		Code      string `json:"code" bson:"code"`
		ShortCode string `json:"shortCode" bson:"shortCode"`
	} `json:"rfid"                 bson:"rfid"`
	DefaultGroup         *model.AttributeEmbedded `json:"defaultGroup"         bson:"defaultGroup"`
	IsActive             bool                     `json:"isActive"             bson:"isActive"`
	IsTemporary          bool                     `json:"isTemporary"          bson:"isTemporary"`
	ExpiryDate           *time.Time               `json:"expiryDate"           bson:"expiryDate"`
	IsAdmin              bool                     `json:"isAdmin"              bson:"isAdmin"`
	PendingEmailChange   string                   `json:"pendingEmailChange"   bson:"pendingEmailChange"`
	ModifiedBy           *model.UserEmbedded      `json:"modifiedBy"           bson:"modifiedBy"`
	Parent               *primitive.ObjectID      `json:"parent"               bson:"parent"`
	ProfilePictureSmall  *string                  `json:"profilePictureSmall"  bson:"profilePictureSmall"`
	ProfilePictureMedium *string                  `json:"profilePictureMedium" bson:"profilePictureMedium"`
	ProfilePictureBig    *string                  `json:"profilePictureBig"  bson:"profilePictureBig"`
	ProfilePicture       *string                  `json:"profilePicture"  bson:"profilePicture"`
	IsDeleted            bool                     `json:"isDeleted"            bson:"isDeleted"`
	IsSubUser            bool                     `json:"isSubUser"            bson:"isSubUser"`
	IsConfirmed          bool                     `json:"isConfirmed"          bson:"isConfirmed"`
	IsDefault            bool                     `json:"isDefault"            bson:"isDefault"`
	SubUserID            *string                  `json:"subUserId"            bson:"subUserId"`
	ParentName           *string                  `json:"parentName"           bson:"parentName"`
	TotalControl         *UserEmbeddedRole        `json:"totalControl"         bson:"totalControl"`
	ReadOnly             *UserEmbeddedRole        `json:"readOnly"             bson:"readOnly"`
	Access               []string                 `json:"access" bson:"access"`
	FCMToken             string                   `json:"fcmToken"    bson:"fcmToken"`
	CreatedAt            time.Time                `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time                `json:"updatedAt"            bson:"updatedAt"`
}
