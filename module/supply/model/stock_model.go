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
	StockStatusWaitingForApproval      StockStatus = "Waiting for Approval"
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
	StockSourceInitialStock   StockSource = "Initial Stock"
)

// DefectSource defines possible sources for stock defect status.
type DefectSource string

const (
	DefectSourceStockOpname    DefectSource = "Stock Opname"
	DefectSourceTransfer       DefectSource = "Transfer"
	DefectSourceInternalReturn DefectSource = "Internal Return"
	DefectSourceTracking       DefectSource = "Tracking"
	DefectSourceBatchNumber    DefectSource = "Batch Number"
)

// Stock represents the schema for single stock data
type Stock struct {
	ID                           	primitive.ObjectID  `json:"_id,omitempty"                   bson:"_id,omitempty"`
	Key                          	int                 `json:"key"                             bson:"key"`
	ItemName                     	primitive.ObjectID  `json:"itemName"                        bson:"itemName"`
	ItemSKU                      	primitive.ObjectID  `json:"itemSKU"                         bson:"itemSKU"`
	Brand                        	primitive.ObjectID  `json:"brand"                           bson:"brand"`
	Model                        	primitive.ObjectID  `json:"model"                           bson:"model"`
	Category                     	primitive.ObjectID  `json:"category"                        bson:"category"`
	Supplier                     	*primitive.ObjectID `json:"supplier"                        bson:"supplier"`
	OldPlacementAdministrativeGroup *primitive.ObjectID `json:"oldPlacementAdministrativeGroup" bson:"oldPlacementAdministrativeGroup"`
	PlacementAdministrativeGroup 	*primitive.ObjectID `json:"placementAdministrativeGroup"    bson:"placementAdministrativeGroup"`
	ReceivingAdministrativeGroup 	primitive.ObjectID  `json:"receivingAdministrativeGroup"    bson:"receivingAdministrativeGroup"`
	OldPlacementGroup            	*primitive.ObjectID `json:"oldPlacementGroup"               bson:"oldPlacementGroup"`
	PlacementGroup               	*primitive.ObjectID `json:"placementGroup"                  bson:"placementGroup"`
	ReceivingGroup               	primitive.ObjectID  `json:"receivingGroup"                  bson:"receivingGroup"`
	TagType                      	model.TagType       `json:"tagType"                         bson:"tagType"`
	Tag                          	model.Tag           `json:"tag"                             bson:"tag"`
	NameWithSequence             	string              `json:"nameWithSequence"                bson:"nameWithSequence"`
	AliasCode                    	string              `json:"aliasCode"                       bson:"aliasCode"`
	Sequence                     	int                 `json:"sequence"                        bson:"sequence"`
	Status                       	[]StockStatus       `json:"status"                          bson:"status"`
	Source                       	*StockSource        `json:"source"                          bson:"source"`
	BatchNumber                  	*primitive.ObjectID `json:"batchNumber"                     bson:"batchNumber"`
	ModifiedBy                   	primitive.ObjectID  `json:"modifiedBy"                      bson:"modifiedBy"`
	OnHold                       	bool                `json:"onHold"                          bson:"onHold"`
	DefectSource                 	*DefectSource       `json:"defectSource"                    bson:"defectSource"`
	ReceptionDate                	time.Time           `json:"receptionDate"                   bson:"receptionDate"`
	PlacementDate                	*time.Time          `json:"placementDate"                   bson:"placementDate"`
	ProductionDate               	*time.Time          `json:"productionDate"                  bson:"productionDate"`
	ExpiredDate                  	*time.Time          `json:"expiredDate"                     bson:"expiredDate"`
	ExpiredReportedDate          	*time.Time          `json:"expiredReportedDate"             bson:"expiredReportedDate"`
	MissingReportedDate          	*time.Time          `json:"missingReportedDate"             bson:"missingReportedDate"`
	DefectReportedDate          	*time.Time          `json:"defectReportedDate"              bson:"defectReportedDate"`
	CustomFields                 	[]CustomField       `json:"customFields"                    bson:"customFields"`
	UpdatedAt                    	time.Time           `json:"updatedAt"                       bson:"updatedAt"`
	CreatedAt                    	time.Time           `json:"createdAt"                       bson:"createdAt"`
}
