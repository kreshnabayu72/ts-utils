// Package model defines the shared analytics-reporting-service wire models for the
// admin-console Hardware Stock module.
package model

import "time"

// HardwareLog records one append-only Hardware Log event of one hardware unit.
type HardwareLog struct {
	// ID is a decimal Snowflake string generated deterministically by the writer.
	ID string `json:"_id" bson:"_id"`

	UnitID string           `json:"unitId" bson:"unitId"`
	SkuID  string           `json:"skuId" bson:"skuId"`
	Event  HardwareLogEvent `json:"event" bson:"event"`
	Title  string           `json:"title" bson:"title"`
	// Description holds one "Title: value" line per catalog row, joined by "\n".
	Description *string `json:"description" bson:"description"`

	// ActorID is nil for system events; names are fixed at log time.
	ActorID        *string              `json:"actorId" bson:"actorId"`
	ActorName      string               `json:"actorName" bson:"actorName"`
	ActorFirstName *string              `json:"actorFirstName" bson:"actorFirstName"`
	ActorLastName  *string              `json:"actorLastName" bson:"actorLastName"`
	ActorFullName  string               `json:"actorFullName" bson:"actorFullName"`
	ActorType      HardwareLogActorType `json:"actorType" bson:"actorType"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type HardwareLogEvent string

const (
	HardwareLogEventRegistered         HardwareLogEvent = "registered"
	HardwareLogEventWaitingForApproval HardwareLogEvent = "waiting_for_approval"
	HardwareLogEventApproved           HardwareLogEvent = "approved"
	HardwareLogEventActivated          HardwareLogEvent = "activated"
	HardwareLogEventReturnedToStock    HardwareLogEvent = "returned_to_stock"
)

type HardwareLogActorType string

const (
	HardwareLogActorTypeUser   HardwareLogActorType = "user"
	HardwareLogActorTypeSystem HardwareLogActorType = "system"
)
