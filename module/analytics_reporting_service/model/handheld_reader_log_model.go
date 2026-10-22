package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandheldReaderActivateLog records an append-only activation request result.
// RangeType scopes which antenna mode the reader was activated in:
// short_range or long_range.
//
// Result is one of: "Activated" | "Waiting for Approval" | "Failed".
//
// CreatedById and CreatedBySnapshot are empty for system-run events.
type HandheldReaderActivateLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	RangeType  string  `json:"rangeType" bson:"rangeType"`
	DeviceName *string `json:"deviceName" bson:"deviceName"`
	SkuCode    *string `json:"skuCode" bson:"skuCode"`
	Result     string  `json:"result" bson:"result"`

	CreatedById       string `json:"createdById" bson:"createdById"`
	CreatedBySnapshot string `json:"createdBySnapshot" bson:"createdBySnapshot"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

// HandheldReaderUsageLog records an append-only reader usage event.
// ReaderID references handheld_readers.id.
//
// TransactionID is empty for events without a product transaction.
//
// CreatedById and CreatedBySnapshot are empty for system-run events.
type HandheldReaderUsageLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	ReaderID  string `json:"readerId" bson:"readerId"`
	Product   string `json:"product" bson:"product"`
	EventType string `json:"eventType" bson:"eventType"`

	TransactionID *string `json:"transactionId" bson:"transactionId"`

	CreatedById       string `json:"createdById" bson:"createdById"`
	CreatedBySnapshot string `json:"createdBySnapshot" bson:"createdBySnapshot"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
