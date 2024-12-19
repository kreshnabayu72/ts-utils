package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Possible values: "Straight-Line" | "Double-Declining Balance"
type DepreciationMethod string

type FileStorage struct {
	UsedKB            float64 `json:"usedKB"            bson:"usedKB"`
	AvailableKB       float64 `json:"availableKB"       bson:"availableKB"`
	TotalKB           float64 `json:"totalKB"           bson:"totalKB"`
	DamagedKB         float64 `json:"damagedKB"         bson:"damagedKB"`
	DisposalKB        float64 `json:"disposalKB"        bson:"disposalKB"`
	RepairTicketingKB float64 `json:"repairTicketingKB" bson:"repairTicketingKB"`
	RoutineKB         float64 `json:"routineKB"         bson:"routineKB"`
}

type Company struct {
	ID                        primitive.ObjectID `json:"_id"                       bson:"_id"`
	DepreciationMethod        DepreciationMethod `json:"depreciationMethod"        bson:"depreciationMethod"`
	PendingDepreciationMethod DepreciationMethod `json:"pendingDepreciationMethod" bson:"pendingDepreciationMethod"`
	FileStorage               FileStorage        `json:"fileStorage"               bson:"fileStorage"`
	CreatedAt                 time.Time          `json:"createdAt"                 bson:"createdAt"`
	UpdatedAt                 time.Time          `json:"updatedAt"                 bson:"updatedAt"`
}
