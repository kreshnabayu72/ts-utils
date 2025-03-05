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

type AuditAttrs struct {
	MaxGracePeriod int `json:"maxGracePeriod" bson:"maxGracePeriod"`
}
type BorrowingAttrs struct {
	MaxBorrowingTime     int `json:"maxBorrowingTime"     bson:"maxBorrowingTime"`
	DefaultBorrowingTime int `json:"defaultBorrowingTime" bson:"defaultBorrowingTime"`
	ConfirmationTimeout  int `json:"confirmationTimeout"  bson:"confirmationTimeout"`
	TransactionTimeout   int `json:"transactionTimeout"   bson:"transactionTimeout"`
}
type TransferAttrs struct {
	TransactionTimeout    int `json:"transactionTimeout"    bson:"transactionTimeout"`
	AssetControlToMissing int `json:"assetControlToMissing" bson:"assetControlToMissing"`
}
type DisposalAttrs struct {
	TransactionTimeout int `json:"transactionTimeout" bson:"transactionTimeout"`
}

// GroupLicense represents the license schema for a group.
type LicenseAttrs struct {
	Managed   *int `bson:"managed"   json:"managed"`
	Unmanaged *int `bson:"unmanaged" json:"unmanaged"`
	Available *int `bson:"available" json:"available"`
	Total     *int `bson:"total"     json:"total"`
	Expired   *int `bson:"expired"   json:"expired"`
}

// License represents the license information for various categories.
type License struct {
	FixedAsset     LicenseAttrs `bson:"fixedAsset"     json:"fixedAsset"`
	Audit          LicenseAttrs `bson:"audit"          json:"audit"`
	TrackingFA     LicenseAttrs `bson:"trackingFA"     json:"trackingFA"`
	TrackingSA     LicenseAttrs `bson:"trackingSA"     json:"trackingSA"`
	Maintenance    LicenseAttrs `bson:"maintenance"    json:"maintenance"`
	AssetControlFA LicenseAttrs `bson:"assetControlFA" json:"assetControlFA"`
	AssetControlSA LicenseAttrs `bson:"assetControlSA" json:"assetControlSA"`
	POS            LicenseAttrs `bson:"pos"            json:"pos"`
	StockOpname    LicenseAttrs `bson:"stockOpname"    json:"stockOpname"`
	SupplyAsset    LicenseAttrs `bson:"supplyAsset"    json:"supplyAsset"`
	Inventory      LicenseAttrs `bson:"inventory"      json:"inventory"`
	Procurement    LicenseAttrs `bson:"procurement"    json:"procurement"`
	ConcurrentUser int          `bson:"concurrentUser" json:"concurrentUser"`
}

type Client struct {
	ID                        primitive.ObjectID  `json:"_id"                       bson:"_id"`
	Name                      string              `json:"name"                      bson:"name"`
	Email                     string              `json:"email"                     bson:"email"`
	Country                   string              `json:"country"                   bson:"country"`
	State                     string              `json:"state"                     bson:"state"`
	City                      string              `json:"city"                      bson:"city"`
	Zip                       string              `json:"zip"                       bson:"zip"`
	PhoneNumber               string              `json:"phoneNumber"               bson:"phoneNumber"`
	Fax                       string              `json:"fax"                       bson:"fax"`
	Address                   string              `json:"address"                   bson:"address"`
	ContactPersonName         string              `json:"contactPersonName"         bson:"contactPersonName"`
	ContactPersonEmail        string              `json:"contactPersonEmail"        bson:"contactPersonEmail"`
	PendingEmailChange        string              `json:"pendingEmailChange"        bson:"pendingEmailChange"`
	Image                     *string             `bson:"image,omitempty"           json:"image,omitempty"`
	OpenAPIStatus             *string             `bson:"openAPIStatus"             json:"openAPIStatus"`
	OpenAPIRequestedAt        *time.Time          `bson:"openAPIRequestedAt"        json:"openAPIRequestedAt"`
	OpenAPIApprovedAt         *time.Time          `bson:"openAPIApprovedAt"         json:"openAPIApprovedAt"`
	ClientId                  string              `json:"clientId"                  bson:"clientId"`
	Partner                   *primitive.ObjectID `json:"partner"                   bson:"partner"`
	Group                     primitive.ObjectID  `json:"group"                     bson:"group"`
	License                   License             `json:"license"                   bson:"license"`
	DepreciationMethod        DepreciationMethod  `json:"depreciationMethod"        bson:"depreciationMethod"`
	PendingDepreciationMethod DepreciationMethod  `json:"pendingDepreciationMethod" bson:"pendingDepreciationMethod"`
	FileStorage               FileStorage         `json:"fileStorage"               bson:"fileStorage"`
	AuditAttrs                AuditAttrs          `json:"auditAttrs"                bson:"auditAttrs"`
	BorrowingAttrs            BorrowingAttrs      `json:"borrowingAttrs"            bson:"borrowingAttrs"`
	TransferAttrs             TransferAttrs       `json:"transferAttrs"             bson:"transferAttrs"`
	DisposalAttrs             DisposalAttrs       `json:"disposalAttrs"             bson:"disposalAttrs"`
	LoginSessionMinutes       int                 `bson:"loginSessionMinutes"       json:"loginSessionMinutes"`
	Pin                       string              `bson:"pin"                       json:"pin"`
	IsEmailConfirmed          bool                `json:"isEmailConfirmed"          bson:"isEmailConfirmed"`
	ModifiedBy                primitive.ObjectID  `json:"modifiedBy"                bson:"modifiedBy"`
	Status                    string              `json:"status"                    bson:"status"`
	ExpiredAt                 *time.Time          `json:"expiredAt"                 bson:"expiredAt"`
	DeleteFromDatabaseAt      *time.Time          `json:"deleteFromDatabaseAt"      bson:"deleteFromDatabaseAt"`
	CreatedAt                 time.Time           `json:"createdAt"                 bson:"createdAt"`
	UpdatedAt                 time.Time           `json:"updatedAt"                 bson:"updatedAt"`

	DateFormat DateFormat `json:"dateFormat" bson:"dateFormat"`
	TimeFormat bool       `json:"timeFormat" bson:"timeFormat"`
	TimeZone   string     `json:"timezone" bson:"timezone"`
	Currency   Currency   `json:"currency" bson:"currency"`
}

type DateFormat struct {
	Locale  string `json:"locale" bson:"locale"`
	Year    string `json:"year" bson:"year"`
	Month   string `json:"month" bson:"month"`
	Day     string `json:"day" bson:"day"`
	Weekday string `json:"weekday" bson:"weekday"`
}
type Currency struct {
	Currency string `json:"currency" bson:"currency"`
	Symbol   string `json:"symbol" bson:"symbol"`
	Locale   string `json:"locale" bson:"locale"`
	Label    string `json:"label" bson:"label"`
}
