package new_model

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
	AttachmentKB      float64 `json:"attachmentKB"            bson:"attachmentKB"`
	DamagedKB         float64 `json:"damagedKB"         bson:"damagedKB"`
	DisposalKB        float64 `json:"disposalKB"        bson:"disposalKB"`
	RepairTicketingKB float64 `json:"repairTicketingKB" bson:"repairTicketingKB"`
	RoutineKB         float64 `json:"routineKB"         bson:"routineKB"`
}

type FileStorageAttribute string

const (
	FileStorageAttributeAttachment = FileStorageAttribute("attachmentKB")
	FileStorageAttributeDamaged    = FileStorageAttribute("damagedKB")
	FileStorageAttributeDisposal   = FileStorageAttribute("disposalKB")
)

type AuditAttrs struct {
	MaxGracePeriod int `json:"maxGracePeriod"            bson:"maxGracePeriod"`
}
type BorrowingAttrs struct {
	MaxBorrowingTime     int `json:"maxBorrowingTime"            bson:"maxBorrowingTime"`
	DefaultBorrowingTime int `json:"defaultBorrowingTime"            bson:"defaultBorrowingTime"`
	ConfirmationTimeout  int `json:"confirmationTimeout"            bson:"confirmationTimeout"`
	TransactionTimeout   int `json:"transactionTimeout"            bson:"transactionTimeout"`
}
type TransferAttrs struct {
	TransactionTimeout    int `json:"transactionTimeout"            bson:"transactionTimeout"`
	AssetControlToMissing int `json:"assetControlToMissing"            bson:"assetControlToMissing"`
}
type DisposalAttrs struct {
	TransactionTimeout int `json:"transactionTimeout"            bson:"transactionTimeout"`
}
type AssetControlAttrs struct {
	Pin int `json:"pin"            bson:"pin"`
}
type LicensePlan struct {
	FixedAsset  PlanTypes `json:"fixedAsset"            bson:"fixedAsset"`
	SupplyAsset PlanTypes `json:"supplyAsset"            bson:"supplyAsset"`
}

type LicenseAmounts struct {
	Total     int `json:"total" bson:"total"`
	Managed   int `json:"managed" bson:"managed"`
	Unmanaged int `json:"unmanaged" bson:"unmanaged"`
	Available int `json:"available" bson:"available"`
}

type LicenseAttrs struct {
	ConcurrentUser int            `json:"concurrentUser" bson:"concurrentUser"`
	FixedAsset     LicenseAmounts `json:"fixedAsset" bson:"fixedAsset"`
}

type Company struct {
	ID                        primitive.ObjectID `json:"_id"                       bson:"_id"`
	DepreciationMethod        DepreciationMethod `json:"depreciationMethod"        bson:"depreciationMethod"`
	PendingDepreciationMethod DepreciationMethod `json:"pendingDepreciationMethod" bson:"pendingDepreciationMethod"`
	FileStorage               FileStorage        `json:"fileStorage"               bson:"fileStorage"`
	AuditAttrs                AuditAttrs         `json:"auditAttrs"                bson:"auditAttrs"`
	BorrowingAttrs            BorrowingAttrs     `json:"borrowingAttrs"            bson:"borrowingAttrs"`
	TransferAttrs             TransferAttrs      `json:"transferAttrs"             bson:"transferAttrs"`
	DisposalAttrs             DisposalAttrs      `json:"disposalAttrs"             bson:"disposalAttrs"`
	AssetControlAttrs         AssetControlAttrs  `json:"assetControlAttrs"         bson:"assetControlAttrs"`
	Plan                      *LicensePlan       `json:"plan"                      bson:"plan"`
	License                   LicenseAttrs       `json:"license"                   bson:"license"`
	CreatedAt                 time.Time          `json:"createdAt"                 bson:"createdAt"`
	UpdatedAt                 time.Time          `json:"updatedAt"                 bson:"updatedAt"`
}
