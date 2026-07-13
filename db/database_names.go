// Package db provides database and table name constants, Valkey service
// constructors, and the shared PostgreSQL service interface.
package db

import "fmt"

type DBName string

const (
	// Potentially DEPRECATED!
	RootDBName        DBName = "tagsamurai"
	AdminDBName       DBName = "admin_tagsamurai"
	FixedAssetDBName  DBName = "tagsamurai_fixed_asset"
	SupplyAssetDBName DBName = "tagsamurai_supply_asset"
	GlobalDBName      DBName = "tagsamurai_global"

	// Used for v2
	AnalyticsReportingDBName   DBName = "tagsamurai_analytics_reporting"
	EnterpriseManagementDBName DBName = "tagsamurai_enterprise_management"
	FileManagerDBName          DBName = "tagsamurai_file_manager"
	HardwareTaggingDBName      DBName = "tagsamurai_hardware_tagging"
	IAMDBName                  DBName = "tagsamurai_iam"
	ImportDBName               DBName = "tagsamurai_import"
	LanguageDBName             DBName = "tagsamurai_language"
	NotificationDBName         DBName = "tagsamurai_notification"
	PlanningStrategyDBName     DBName = "tagsamurai_planning_strategy"
	ResourceAllocationDBName   DBName = "tagsamurai_resource_allocation"
	MasterAssetRegistryDBName  DBName = "tagsamurai_master_asset_registry"
	AssetTransactionDBName     DBName = "tagsamurai_asset_transaction"
	ApprovalDBName             DBName = "tagsamurai_approval"
	AssetAddonDBName           DBName = "tagsamurai_asset_addon"
)

// CompanyDBName returns the database name with company code prefix.
// Resulting format: <company_code>:<database_name>
func (dbName DBName) CompanyDBName(companyCode string) DBName {
	// Since now company code no longer used in database name, we use it as a prefix for easy split later for schema name
	return DBName(fmt.Sprintf("%s:%s", companyCode, dbName))
}
