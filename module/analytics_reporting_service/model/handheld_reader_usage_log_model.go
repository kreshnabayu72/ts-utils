package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HandheldReaderUsageLog struct {
	ID            primitive.ObjectID `json:"_id"              bson:"_id"`
	ReaderID      string             `json:"readerId"         bson:"readerId"`
	Activity      string             `json:"activity"         bson:"activity"`
	PIC           string             `json:"pic"              bson:"pic"`
	User          string             `json:"user"          bson:"user"`
	Manager       string             `json:"manager"          bson:"manager"`
	Transaction   string             `json:"transaction"             bson:"transaction"`
	TransactionID string             `json:"transactionId"             bson:"transactionId"`
	Borrower      string             `json:"borrower" bson:"borrower"`
	Origin        string             `json:"origin" bson:"origin"`
	Destination   string             `json:"destination" bson:"nadestination"`
	ReportedBy    string             `json:"reportedBy" bson:"reportedBy"`
	Group         string             `json:"group" bson:"group"`
	AuditType     string             `json:"auditType"        bson:"auditType"`
	Condition     string             `json:"condition"        bson:"condition"`
	Staff         string             `json:"staff"        bson:"staff"`
	Stock         string             `json:"stock"            bson:"stock"`
	SKU           string             `json:"sku"              bson:"sku"`
	Supplier      string             `json:"supplier"              bson:"supplier"`
	Requester     string             `json:"requester"              bson:"requester"`
	CreatedAt     time.Time          `json:"createdAt"        bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"        bson:"updatedAt"`
}
