package db

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	admin_model "gitea.qwertysystem.net/BETS/ts-utils/module/admin/new_model"
	fixedasset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	global_model "gitea.qwertysystem.net/BETS/ts-utils/module/global/new_model"
	root_model "gitea.qwertysystem.net/BETS/ts-utils/module/root/new_model"
)

func MigrateDB(dbName DBName, host, username, password, sslModel string) error {
	strDBName := string(dbName)

	migrator := migration.MakeMigrator(host, username, password, sslModel)
	if strDBName == string(RootDBName) {
		log.Println("Migrating Root DB")
		return migrateRootDB(strDBName, migrator)
	}

	if strDBName == string(AdminDBName) {
		log.Println("Migrating Admin DB")
		return migrateAdminDB(strDBName, migrator)
	}

	if strings.Contains(strDBName, string(GlobalDBName)) {
		log.Println("Migrating Global DB")
		return migrateGlobalDB(strDBName, migrator)
	}

	if strings.Contains(strDBName, string(FixedAssetDBName)) {
		log.Println("Migrating Fixed Asset DB")
		return migrateFixedAssetDB(strDBName, migrator)
	}

	if strings.Contains(strDBName, string(SupplyAssetDBName)) {
		log.Println("Migrating Supply Asset DB")
		return migrateSupplyAssetDB(strDBName)
	}

	return nil
}

func migrateRootDB(dbName string, migrator *migration.Migrator) error {
	models := []migration.Model{
		&root_model.UserToken{},
		&root_model.LoginSession{},
		&root_model.ResetPasswordToken{},
		&root_model.UserAuth{},
	}

	isDBExist := migrator.IsDBExist(dbName)
	if !isDBExist {
		err := migrator.MigrateDB(dbName)
		if err != nil {
			return err
		}
	}

	err := migrator.MigrateModels(dbName, models...)

	return err
}

func migrateAdminDB(dbName string, migrator *migration.Migrator) error {
	models := []migration.Model{
		&admin_model.Client{},
		&admin_model.User{},
	}

	isDBExist := migrator.IsDBExist(dbName)
	if !isDBExist {
		err := migrator.MigrateDB(dbName)
		if err != nil {
			return err
		}
	}

	err := migrator.MigrateModels(dbName, models...)

	return err
}

func migrateGlobalDB(dbName string, migrator *migration.Migrator) error {
	models := []migration.Model{
		&global_model.User{},
		&global_model.Position{},
		&global_model.Division{},
	}

	isDBExist := migrator.IsDBExist(dbName)
	if !isDBExist {
		err := migrator.MigrateDB(dbName)
		if err != nil {
			return err
		}
	}

	err := migrator.MigrateModels(dbName, models...)

	return err
}

func migrateFixedAssetDB(dbName string, migrator *migration.Migrator) error {
	models := []migration.Model{
		&fixedasset_model.Measurement{},
		&fixedasset_model.User{},
		&fixedasset_model.AliasCode{},
		&fixedasset_model.AliasCodeData{},
		&fixedasset_model.AccountCode{},
		&fixedasset_model.DepreciationGroup{},
		&fixedasset_model.Group{},
		&fixedasset_model.GroupApprovalType{},
		&fixedasset_model.GroupQuota{},
		&fixedasset_model.Category{},
	}

	isDBExist := migrator.IsDBExist(dbName)
	if !isDBExist {
		err := migrator.MigrateDB(dbName)
		if err != nil {
			return err
		}
	}

	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	functionQueryPath := filepath.Join(baseDir, "migration/fixed_asset_function.sql")
	err := migrator.ExecFile(dbName, functionQueryPath)
	if err != nil {
		return err
	}

	enumQueryPath := filepath.Join(baseDir, "migration/fixed_asset_enum.sql")
	err = migrator.ExecFile(dbName, enumQueryPath)
	if err != nil {
		return err
	}

	err = migrator.MigrateModels(dbName, models...)

	return err
}

func migrateSupplyAssetDB(dbName string) error {
	// TODO: Implement the migration logic for the Supply Asset DB
	return nil
}
