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
)

func (dbName DBName) CompanyDBName(companyCode string) DBName {
	return DBName(fmt.Sprintf("%s_%s", companyCode, dbName))
}
