package model

import (
	"time"

	util_common_model "gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierReturnTransactionStatus string

const (
	SupplierReturnStatusWaitingForApproval SupplierReturnTransactionStatus = "Waiting for Approval"
	SupplierReturnStatusInTransit          SupplierReturnTransactionStatus = "In Transit"
	SupplierReturnStatusCompletionApproval SupplierReturnTransactionStatus = "Completion Approval"
	SupplierReturnStatusCompleted          SupplierReturnTransactionStatus = "Completed"
	SupplierReturnStatusCanceled           SupplierReturnTransactionStatus = "Canceled"
	SupplierReturnStatusReturned           SupplierReturnTransactionStatus = "Returned"
	SupplierReturnStatusApproved           SupplierReturnTransactionStatus = "Approved"
	SupplierReturnStatusRejected           SupplierReturnTransactionStatus = "Rejected"
)


//for return request
type SupplierReturnRequest struct {
	ID             primitive.ObjectID              `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId  primitive.ObjectID              `json:"transactionId" bson:"transactionId"`
	Status         SupplierReturnTransactionStatus `json:"status" bson:"status"`
	ItemName       primitive.ObjectID              `json:"itemName"       bson:"itemName"`
	ItemSKU        primitive.ObjectID              `json:"itemSKU"        bson:"itemSKU"`
	Brand          primitive.ObjectID              `json:"brand"          bson:"brand"`
	Model          primitive.ObjectID              `json:"model"          bson:"model"`
	PIC            primitive.ObjectID              `json:"pic"            bson:"pic"`
	ReturnGroup    primitive.ObjectID              `json:"returnGroup" bson:"returnGroup"`
	Group          primitive.ObjectID              `json:"group" bson:"group"`
	Stocks         *[]primitive.ObjectID           `json:"stocks" bson:"stocks"`
	RequestStocks  *[]SupplierReturnStocks        `json:"requeststocks" bson:"requeststocks"`
	ExchangeStocks *int                            `json:"exchangeStocks" bson:"exchangeStocks"`
	BatchID        primitive.ObjectID              `json:"batchId" bson:"batchId"`
	ReturnStocks   *int                            `json:"returnStocks" bson:"returnStocks"`
	ApprovalNote   *string                         `json:"approvalNote" bson:"approvalNote"`
	TagType        util_common_model.TagType       `json:"tagType"        bson:"tagType"`
	UpdatedAt      time.Time                       `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt      time.Time                       `json:"createdAt"      bson:"createdAt"`
}

type SupplierReturnRequestHistory struct {
	ID             primitive.ObjectID                          `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId  primitive.ObjectID                          `json:"transactionId"  bson:"transactionId"`
	Status         SupplierReturnTransactionStatus             `json:"status"         bson:"status"`
	ItemName       AttributeEmbedded         `json:"itemName"       bson:"itemName"`
	ItemSKU        SKUAttributeEmbedded      `json:"itemSKU"        bson:"itemSKU"`
	Brand          AttributeEmbedded         `json:"brand"          bson:"brand"`
	Model          AttributeEmbedded         `json:"model"          bson:"model"`
	ReturnGroup    TreeAttributeEmbedded     `json:"returnGroup" bson:"returnGroup"`
	Group          TreeAttributeEmbedded     `json:"group" bson:"group"`
	PIC            UserAttributeEmbedded     `json:"pic"            bson:"pic"`
	Stocks         *[]StockAttributeEmbedded `json:"stocks" bson:"stocks"`
	RequestStocks  *[]SupplierReturnStocks                    `json:"requeststocks" bson:"requeststocks"`
	ExchangeStocks *int                                        `json:"exchangeStocks" bson:"exchangeStocks"`
	BatchID        primitive.ObjectID                          `json:"batchId" bson:"batchId"`
	ReturnStocks   *int                                        `json:"returnStocks" bson:"returnStocks"`
	ApprovalNote   *string                                     `json:"approvalNote" bson:"approvalNote"`
	TagType        util_common_model.TagType                   `json:"tagType"        bson:"tagType"`
	UpdatedAt      time.Time                                   `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt      time.Time                                   `json:"createdAt"      bson:"createdAt"`
}

type SupplierReturnStocks struct {
	Stocks primitive.ObjectID `json:"stocks" bson:"stocks"`
	Reason string             `json:"reason" bson:"reason"`
}


//for transaction
type SupplierReturnTransaction struct {
	ID            primitive.ObjectID              `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId string                          `json:"transactionId"  bson:"transactionId"`
	Supplier      primitive.ObjectID              `json:"supplier"       bson:"supplier"`
	Status        SupplierReturnTransactionStatus `json:"status"         bson:"status"`
	ReturnGroup   primitive.ObjectID              `json:"returnGroup" bson:"returnGroup"`
	Note          *string                         `json:"note"           bson:"note"`
	PIC           primitive.ObjectID              `json:"pic"            bson:"pic"`
	TotalItem     int                             `json:"totalItem"      bson:"totalItem"`
	TotalStock    int                             `json:"totalStock"     bson:"totalStock"`
	ApprovalNote  *string                         `json:"approvalNote" bson:"approvalNote"`
	UpdatedAt     time.Time                       `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt     time.Time                       `json:"createdAt"      bson:"createdAt"`
}
type SupplierReturnTransactionHistory struct {
	ID            primitive.ObjectID                      `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId string                                  `json:"transactionId"  bson:"transactionId"`
	Supplier      AttributeEmbedded     `json:"supplier"       bson:"supplier"`
	Status        SupplierReturnTransactionStatus         `json:"status"         bson:"status"`
	ReturnGroup   TreeAttributeEmbedded `json:"returnGroup" bson:"returnGroup"`
	Note          *string                                 `json:"note"           bson:"note"`
	PIC           UserAttributeEmbedded `json:"pic"            bson:"pic"`
	TotalItem     int                                     `json:"totalItem"      bson:"totalItem"`
	TotalStock    int                                     `json:"totalStock"     bson:"totalStock"`
	ApprovalNote  *string                                 `json:"approvalNote" bson:"approvalNote"`
	UpdatedAt     time.Time                               `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt     time.Time                               `json:"createdAt"      bson:"createdAt"`
}

//for adding to prelist
type SupplierReturnPreList struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Supplier    primitive.ObjectID `json:"supplier" bson:"supplier"`
	ReturnGroup primitive.ObjectID `json:"returnGroup" bson:"returnGroup"`
	ReportedBy  primitive.ObjectID `json:"reportedBy" bson:"reportedBy"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type SupplierReturnPreListItem struct {
	ID         primitive.ObjectID        `json:"_id,omitempty" bson:"_id,omitempty"`
	PreListId  primitive.ObjectID        `json:"preListId" bson:"preListId"`
	ItemName   primitive.ObjectID        `json:"itemName"       bson:"itemName"`
	ItemSKU    primitive.ObjectID        `json:"itemSKU"        bson:"itemSKU"`
	Brand      primitive.ObjectID        `json:"brand"          bson:"brand"`
	Model      primitive.ObjectID        `json:"model"          bson:"model"`
	Group      primitive.ObjectID        `json:"group" bson:"group"`
	TagType    util_common_model.TagType `json:"tagType"        bson:"tagType"`
	ReportedBy primitive.ObjectID        `json:"reportedBy" bson:"reportedBy"`
	CreatedAt  time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time                 `json:"updatedAt" bson:"updatedAt"`
}


//make or update completion
type SupplierReturnCompletion struct {
	ID              primitive.ObjectID  `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId   primitive.ObjectID  `json:"transactionId"  bson:"transactionId"`
	ProofOfDelivery string              `json:"proofOfDelivery"       bson:"proofOfDelivery"`
	DeliveryNote    *string             `json:"deliveryNote"           bson:"deliveryNote"`
	ReceivingGroup  *primitive.ObjectID `json:"receivingGroup" bson:"receivingGroup"`
	Document        *string             `json:"document"       bson:"document"`
	ExchangeNote    *string             `json:"exchangeNote"           bson:"exchangeNote"`
	Photo           *string             `json:"photo"       bson:"photo"`
	UpdatedAt       time.Time           `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt       time.Time           `json:"createdAt"      bson:"createdAt"`
}