package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

// StockStatus represents possible stock status values
type StockStatus string

const (
	StockStatusAvailable               StockStatus = "Available"
	StockStatusOnCheckoutProcess       StockStatus = "On Checkout Process"
	StockStatusOnTransfer              StockStatus = "On Transfer"
	StockStatusInStockOpname           StockStatus = "In Stock Opname"
	StockStatusDefect                  StockStatus = "Defect"
	StockStatusMissing                 StockStatus = "Missing"
	StockStatusExpired                 StockStatus = "Expired"
	StockStatusUnplaced                StockStatus = "Unplaced"
	StockStatusTAGReported             StockStatus = "TAG Reported"
	StockStatusNotPairedYet            StockStatus = "Not Paired Yet"
	StockStatusWaitingForApproval      StockStatus = "Waiting For Approval"
	StockStatusInSupplierReturnProcess StockStatus = "In Supplier Return Process"
	StockStatusOnDisposalProcess       StockStatus = "On Disposal Process"
	StockStatusDisposed                StockStatus = "Disposed"
	StockStatusCheckedOut              StockStatus = "Checked Out"
	StockStatusExchanged               StockStatus = "Exchanged"
	StockStatusReturned                StockStatus = "Returned"
)

// StockSource represents stock source values.
type StockSource string

const (
	StockSourceReception      StockSource = "Reception"
	StockSourceSupplierReturn StockSource = "Supplier Return"
	StockSourceInternalReturn StockSource = "Internal Return"
)

// Stock represents the schema for a single stock data
type Stock struct {
	ID                           primitive.ObjectID  `json:"_id,omitempty"                bson:"_id,omitempty"`
	ItemName                     primitive.ObjectID  `json:"itemName"                     bson:"itemName"`
	ItemSKU                      primitive.ObjectID  `json:"itemSKU"                      bson:"itemSKU"`
	Brand                        primitive.ObjectID  `json:"brand"                        bson:"brand"`
	Model                        primitive.ObjectID  `json:"model"                        bson:"model"`
	Category                     primitive.ObjectID  `json:"category"                     bson:"category"`
	PlacementAdministrativeGroup *primitive.ObjectID `json:"placementAdministrativeGroup" bson:"placementAdministrativeGroup"`
	ReceivingAdministrativeGroup primitive.ObjectID  `json:"ReceivingAdministrativeGroup" bson:"ReceivingAdministrativeGroup"`
	PlacementGroup               *primitive.ObjectID `json:"placementGroup"               bson:"placementGroup"`
	ReceivingGroup               primitive.ObjectID  `json:"ReceivingGroup"               bson:"receivingGroup"`
	TagType                      model.TagType       `json:"tagType"                      bson:"tagType"`
	Tag                          model.Tag           `json:"tag"                          bson:"tag"`
	NameWithSequence             string              `json:"nameWithSequence"             bson:"nameWithSequence"`
	Sequence                     int                 `json:"sequence"                     bson:"sequence"`
	Status                       []StockStatus       `json:"status"                       bson:"status"`
	Source                       StockSource         `json:"source"                       bson:"source"`
	SourceTransactionLogId       primitive.ObjectID  `json:"sourceTransactionLogId"       bson:"sourceTransactionLogId"`
	ReceptionDate                time.Time           `json:"receptionDate"                bson:"receptionDate"`
	ProductionDate               *time.Time          `json:"productionDate"               bson:"productionDate"`
	ExpiredDate                  *time.Time          `json:"expiredDate"                  bson:"expiredDate"`
	CustomFields                 []CustomField       `json:"customFields"                 bson:"customFields"`
	UpdatedAt                    time.Time           `json:"updatedAt"                    bson:"updatedAt"`
	CreatedAt                    time.Time           `json:"createdAt"                    bson:"createdAt"`
}
