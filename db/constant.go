package db

import "fmt"

type DBName string

const (
	RootDBName        DBName = "tagsamurai"
	FixedAssetDBName  DBName = "tagsamurai_fixed_asset"
	SupplyAssetDBName DBName = "tagsamurai_supply_asset"
	GlobalDBName      DBName = "tagsamurai_global"
)

func (dbName DBName) CompanyDBName(companyCode string) DBName {
	return DBName(fmt.Sprintf("%s_%s", companyCode, dbName))
}
