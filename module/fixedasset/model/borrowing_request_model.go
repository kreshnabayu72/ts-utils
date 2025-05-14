package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowingStatus string

const (
	BorrowStatusBorrowed = BorrowingStatus("Borrowed")
	BorrowStatusDamaged  = BorrowingStatus("Reported Damaged")
	BorrowStatusMissing  = BorrowingStatus("Reported Missing")
	BorrowStatusReturned = BorrowingStatus("Returned")
)

type BorrowingRequest struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Transaction      primitive.ObjectID      `json:"transaction" bson:"transaction"`
	TransactionId    string                  `json:"transactionId" bson:"transactionId"`
	Asset            primitive.ObjectID      `json:"asset" bson:"asset"`
	AssetId          string                  `json:"assetId" bson:"assetId"`
	AssetParent      *primitive.ObjectID     `json:"assetParent" bson:"assetParent"`
	AssetImageSmall  *string                 `json:"assetImageSmall" bson:"assetImageSmall"`
	AssetImageMedium *string                 `json:"assetImageMedium" bson:"assetImageMedium"`
	AssetImageBig    *string                 `json:"assetImageBig" bson:"assetImageBig"`
	Status           BorrowingStatus         `json:"status" bson:"status"`
	AssetName        AssetNameEmbedAttr      `json:"assetName" bson:"assetName"`
	AssetCategory    model.AttributeEmbedded `json:"assetCategory" bson:"assetCategory"`
	AssetGroup       model.AttributeEmbedded `json:"assetGroup" bson:"assetGroup"`
	AssetBrand       model.AttributeEmbedded `json:"assetBrand" bson:"assetBrand"`
	AssetModel       model.AttributeEmbedded `json:"assetModel" bson:"assetModel"`

	Manager  model.UserEmbedded `json:"manager" bson:"manager"`
	Borrower model.UserEmbedded `json:"borrower" bson:"borrower"`

	TagType TagType `json:"tagType" bson:"tagType"`

	Notes string `json:"notes" bson:"notes"`

	Duration int `json:"duration" bson:"duration"`

	IsActive         bool      `json:"isActive" bson:"isActive"`
	DueReturnDate    time.Time `json:"dueReturnDate" bson:"dueReturnDate"`
	ActualReturnDate time.Time `json:"actualReturnDate" bson:"actualReturnDate"`
	CreatedAt        time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt" bson:"updatedAt"`
}
