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
	PlanningStrategyDBName DBName = "tagsamurai_planning_strategy"
	IAMDBName              DBName = "tagsamurai_iam"
)

func (dbName DBName) CompanyDBName(companyCode string) DBName {
	return DBName(fmt.Sprintf("%s_%s", companyCode, dbName))
}
