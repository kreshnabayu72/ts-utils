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
		&fixedasset_model.Asset{},
		&fixedasset_model.AssetName{},
		&fixedasset_model.BrandCategory{},
		&fixedasset_model.Brand{},
		&fixedasset_model.CustomField{},
		&fixedasset_model.ImportAccountCode{},
		&fixedasset_model.ImportFixedAsset{},
		&fixedasset_model.ImportAssetName{},
		&fixedasset_model.ImportBrand{},
		&fixedasset_model.ImportCategory{},
		&fixedasset_model.ImportCustomField{},
		&fixedasset_model.ImportDepreciationGroup{},
		&fixedasset_model.ImportGroup{},
		&fixedasset_model.ImportMeasurement{},
		&fixedasset_model.ImportUser{},
		&fixedasset_model.Model{},
		&fixedasset_model.Qr{},
		&fixedasset_model.QrRfidPair{},
		&fixedasset_model.Rfid{},
		&fixedasset_model.UserTransactionRole{},
		&fixedasset_model.AssetCustomField{},
		&fixedasset_model.Attachment{},
		&fixedasset_model.BorrowingApproval{},
		&fixedasset_model.BorrowingPrelist{},
		&fixedasset_model.BorrowingPrelistRequest{},
		&fixedasset_model.BorrowingRequest{},
		&fixedasset_model.BorrowingTransaction{},
		&fixedasset_model.CustomFieldCategory{},
		&fixedasset_model.DamageReport{},
		&fixedasset_model.FileInStorage{},
		&fixedasset_model.HandheldReader{},
		&fixedasset_model.IotAntenna{},
		&fixedasset_model.IotReader{},
		&fixedasset_model.IotReaderGateReceivingGroup{},
		&fixedasset_model.IotReaderPort{},
		&fixedasset_model.LicenseTransaction{},
		&fixedasset_model.MissingReport{},
		&fixedasset_model.ReplaceTagPrelist{},
		&fixedasset_model.UserSystemRole{},
		&fixedasset_model.AssignmentApproval{},
		&fixedasset_model.AssignmentPrelist{},
		&fixedasset_model.AssignmentPrelistRequest{},
		&fixedasset_model.AssignmentRequest{},
		&fixedasset_model.AssignmentTransaction{},
		&fixedasset_model.DisposalApproval{},
		&fixedasset_model.DisposalPrelist{},
		&fixedasset_model.DisposalPrelistRequest{},
		&fixedasset_model.DisposalRequest{},
		&fixedasset_model.DisposalTransaction{},
		&fixedasset_model.TransferApproval{},
		&fixedasset_model.TransferPrelist{},
		&fixedasset_model.TransferPrelistRequest{},
		&fixedasset_model.TransferRequest{},
		&fixedasset_model.TransferTransaction{},
		&fixedasset_model.RoutineSchedule{},
		&fixedasset_model.RoutineScheduleStaff{},
		&fixedasset_model.RoutineScheduledAsset{},
		&fixedasset_model.RoutineScheduleIndicator{},
		&fixedasset_model.RoutineTask{},
		&fixedasset_model.RoutineTaskStaff{},
		&fixedasset_model.RoutineTaskIndicator{},
		&fixedasset_model.RoutineApproval{},
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
	log.Println("Migrate model error:", err)

	viewQueryPath := filepath.Join(baseDir, "migration/fixed_asset_view.sql")
	err = migrator.ExecFile(dbName, viewQueryPath)
	if err != nil {
		return err
	}

	return err
}

func migrateSupplyAssetDB(dbName string) error {
	// TODO: Implement the migration logic for the Supply Asset DB
	return nil
}
