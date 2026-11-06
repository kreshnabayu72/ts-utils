// Package dto defines the hardware-tagging-service activity parameters for the
// admin-console Hardware Stock module.
package dto

import "time"

// ReceiptUnit is one unit of an Add Stock batch as captured by the client or the
// Handheld Reader driver.
type ReceiptUnit struct {
	SerialNumber    string `json:"serialNumber"`
	MinPowerDbm     *int   `json:"minPowerDbm,omitempty"`
	MaxPowerDbm     *int   `json:"maxPowerDbm,omitempty"`
	CurrentPowerDbm *int   `json:"currentPowerDbm,omitempty"`
}

// CreateHardwareStockReceiptParam is the CreateHardwareStockReceiptActivity input.
// SerialNumbers is used for manual batches; Units for Handheld Reader Sync batches.
type CreateHardwareStockReceiptParam struct {
	ClientID      string
	EntityType    string
	EntityID      *int64
	EntityName    string
	UserID        int64
	UserFirstName string
	UserLastName  string
	UserFullName  string

	SkuID        int64
	SkuCode      string
	DeviceName   string
	HardwareType string
	RangeType    string

	InputSource    string
	TransactionID  string
	IsTotalControl bool

	SerialNumbers []string
	Units         []ReceiptUnit
}

// DuplicateRef is one batch unit whose serial matched an existing unit; the
// enterprise workflow turns it into a Register Duplicate approval subject.
type DuplicateRef struct {
	UnitID               int64
	SerialNumber         string
	ExistingRecordStatus string
	ExistingRecordID     *int64
	MatchedRecordType    string
	MatchedRecordID      *int64
	MatchedIdentifier    string
	RegistrantEntityName string
	RegistrantUserName   string
	RegisteredAt         time.Time
}

// CreateHardwareStockReceiptResult is the CreateHardwareStockReceiptActivity output.
type CreateHardwareStockReceiptResult struct {
	ReceiptID            int64
	RegisteredCount      int
	PendingApprovalCount int
	Duplicates           []DuplicateRef
}

// CompensateHardwareStockReceiptParam deletes the receipt and units created by one
// HardwareStockReceiptWorkflow run.
type CompensateHardwareStockReceiptParam struct {
	ClientID  string
	ReceiptID int64
}

// ApplyHardwareUnitDecisionParam transitions one approval subject unit.
type ApplyHardwareUnitDecisionParam struct {
	ClientID string
	UnitID   int64
	Decision string // "approved" | "rejected"
	Note     string
}

// CallerScope is the entity context of a tenant client, resolved by
// enterprise-management-service for HTS and AR.
type CallerScope struct {
	EntityID                 *int64
	EntityType               string
	DistributorID            *int64
	LocalRegistrationEnabled bool
	Suspended                bool
}
