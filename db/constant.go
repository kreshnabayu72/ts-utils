package db

import "fmt"

type DBName string

const (
	RootDBName        DBName = "tagsamurai"
	AdminDBName       DBName = "admin_tagsamurai"
	FixedAssetDBName  DBName = "tagsamurai_fixed_asset"
	SupplyAssetDBName DBName = "tagsamurai_supply_asset"
	GlobalDBName      DBName = "tagsamurai_global"

	PlanningStrategyDBName DBName = "tagsamurai_planning_strategy"
)

func (dbName DBName) CompanyDBName(companyCode string) DBName {
	return DBName(fmt.Sprintf("%s_%s", companyCode, dbName))
}
