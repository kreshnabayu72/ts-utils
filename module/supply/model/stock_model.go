package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

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
	StockStatusWaitingForApproval      StockStatus = "Waiting for Approval"
	StockStatusInSupplierReturnProcess StockStatus = "In Supplier Return Process"
	StockStatusOnDisposalProcess       StockStatus = "On Disposal Process"
	StockStatusDisposed                StockStatus = "Disposed"
	StockStatusCheckedOut              StockStatus = "Checked Out"
	StockStatusExchanged               StockStatus = "Exchanged"
	StockStatusReturned                StockStatus = "Returned"
)

// AddStockStatus adds a new stock status to an existing stock statuses,
// validate the transition according to the rules and returns the result
func AddStockStatus(s []StockStatus, statuses ...StockStatus) ([]StockStatus, error) {
	// TODO(refan): Programmatically adds status based on the transition rules
	s = append(s, s...)
	return s, nil
}

// SetStockStatus sets a new stock status,
// validate the transition according to the rules based from the existing stock statuses and returns the result
func SetStockStatus(s []StockStatus, statuses ...StockStatus) ([]StockStatus, error) {
	// TODO(refan): Programmatically sets status ONLY if the transition rules allows it
	return statuses, nil
}

// RemoveStockStatus removes a stock status,
// validate the transition according to the rules based from the existing stock statuses and returns the result
func RemoveStockStatus(s []StockStatus, statuses ...StockStatus) ([]StockStatus, error) {
	// TODO(refan): Programmatically removes status based on the transition rules
	return s, nil
}

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
	Supplier                     *primitive.ObjectID `json:"supplier"                     bson:"supplier"`
	PlacementAdministrativeGroup *primitive.ObjectID `json:"placementAdministrativeGroup" bson:"placementAdministrativeGroup"`
	ReceivingAdministrativeGroup primitive.ObjectID  `json:"receivingAdministrativeGroup" bson:"receivingAdministrativeGroup"`
	PlacementGroup               *primitive.ObjectID `json:"placementGroup"               bson:"placementGroup"`
	ReceivingGroup               primitive.ObjectID  `json:"receivingGroup"               bson:"receivingGroup"`
	TagType                      model.TagType       `json:"tagType"                      bson:"tagType"`
	Tag                          model.Tag           `json:"tag"                          bson:"tag"`
	NameWithSequence             string              `json:"nameWithSequence"             bson:"nameWithSequence"`
	AliasCode                    string              `json:"aliasCode"                    bson:"aliasCode"`
	Sequence                     int                 `json:"sequence"                     bson:"sequence"`
	Status                       []StockStatus       `json:"status"                       bson:"status"`
	Source                       *StockSource        `json:"source"                       bson:"source"`
	ModifiedBy                   primitive.ObjectID  `json:"modifiedBy"                   bson:"modifiedBy"`
	ReceptionDate                time.Time           `json:"receptionDate"                bson:"receptionDate"`
	PlacementDate                *time.Time          `json:"placementDate"                bson:"placementDate"`
	ProductionDate               *time.Time          `json:"productionDate"               bson:"productionDate"`
	ExpiredDate                  *time.Time          `json:"expiredDate"                  bson:"expiredDate"`
	CustomFields                 []CustomField       `json:"customFields"                 bson:"customFields"`
	UpdatedAt                    time.Time           `json:"updatedAt"                    bson:"updatedAt"`
	CreatedAt                    time.Time           `json:"createdAt"                    bson:"createdAt"`
}
