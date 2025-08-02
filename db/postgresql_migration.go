package db

import (
	"strings"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	global_model "gitea.qwertysystem.net/BETS/ts-utils/module/global/new_model"
	root_model "gitea.qwertysystem.net/BETS/ts-utils/module/root/new_model"
)

func MigrateDB(dbName DBName, host, username, password, sslModel string) error {
	strDBName := string(dbName)

	migrator := migration.MakeMigrator(host, username, password, sslModel)
	if strDBName == string(RootDBName) {
		return migrateRootDB(strDBName, migrator)
	}

	if strDBName == string(AdminDBName) {
		return migrateAdminDB(strDBName)
	}

	if strings.Contains(strDBName, string(GlobalDBName)) {
		return migrateGlobalDB(strDBName, migrator)
	}

	if strings.Contains(strDBName, string(FixedAssetDBName)) {
		return migrateFixedAssetDB(strDBName)
	}

	if strings.Contains(strDBName, string(SupplyAssetDBName)) {
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

func migrateAdminDB(dbName string) error {
	// TODO: Implement the migration logic for the Admin DB
	return nil
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

func migrateFixedAssetDB(dbName string) error {
	// TODO: Implement the migration logic for the Fixed Asset DB
	return nil
}

func migrateSupplyAssetDB(dbName string) error {
	// TODO: Implement the migration logic for the Supply Asset DB
	return nil
}
