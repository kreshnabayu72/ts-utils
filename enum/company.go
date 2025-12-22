package enum

type CompanyStatus string

const (
	CompanyStatusActive   CompanyStatus = "Active"
	CompanyStatusExpired  CompanyStatus = "Expired"
	CompanyStatusArchived CompanyStatus = "Archived"
)

type CompanyOpenAPIStatus string

const (
	CompanyOpenAPIStatusApproved CompanyOpenAPIStatus = "Approved"
	CompanyOpenAPIStatusWaiting  CompanyOpenAPIStatus = "Waiting for Approval"
)

type CompanyDepreciationMethod string

const (
	CompanyDepreciationMethodStraightLine           CompanyDepreciationMethod = "Straight-Line"
	CompanyDepreciationMethodDoubleDecliningBalance CompanyDepreciationMethod = "Double-Declining Balance"
)

type LicenseType string

const (
	LicenseTypeConcurrentUser            LicenseType = "concurrent_user"
	LicenseTypeFixedAsset                LicenseType = "fixed_asset"
	LicenseTypeFixedAssetAudit           LicenseType = "fixed_asset_audit"
	LicenseTypeFixedAssetAssetControl    LicenseType = "fixed_asset_asset_control"
	LicenseTypeFixedAssetMaintenance     LicenseType = "fixed_asset_maintenance"
	LicenseTypeFixedAssetRepairTicketing LicenseType = "fixed_asset_repair_ticketing"
	LicenseTypeFixedAssetTracking        LicenseType = "fixed_asset_tracking"
	LicenseTypeSupplyAsset               LicenseType = "supply_asset"
	LicenseTypeSupplyAssetAssetControl   LicenseType = "supply_asset_asset_control"
	LicenseTypeSupplyAssetStockOpname    LicenseType = "supply_asset_stock_opname"
	LicenseTypeSupplyAssetTracking       LicenseType = "supply_asset_tracking"
	LicenseTypeInventory                 LicenseType = "inventory"
)

type LicensePurchaseType string

const (
	LicensePurchaseTypePurchase LicensePurchaseType = "Purchase"
	LicensePurchaseTypeTrial    LicensePurchaseType = "Trial"
)
