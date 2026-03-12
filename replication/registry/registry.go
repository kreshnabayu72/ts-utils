package registry

import (
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/asset_transaction_service/model"
	ems_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/enterprise_management_service/model"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/file_manager_service/replica"
	hardware_tagging_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/hardware_tagging_service/model"
	hardware_tagging_replica "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/hardware_tagging_service/replica"
	iam_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/iam_service/model"
	iam_replica "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/iam_service/replica"
	master_asset_registry_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/master_asset_registry_service/model"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/api"
)

// MainModels defines the source of truth for the full schema of each domain table.
// This is used to calculate which columns to EXCLUDE in PeerDB CDC flows by taking
// the set difference: Source (Main Model) - Consumer (Union of Replica Models).
var MainModels = map[db.DBName]map[string]api.GetColumnsProvider{
	db.IAMDBName: {
		"divisions":                          &iam_model.Division{},
		"fixed_asset_groups":                 &iam_model.FixedAssetGroup{},
		"fixed_asset_group_approval_types":   &iam_model.FixedAssetGroupApprovalType{},
		"fixed_asset_user_system_roles":      &iam_model.FixedAssetUserSystemRole{},
		"fixed_asset_user_transaction_roles": &iam_model.FixedAssetUserTransactionRole{},
		"iam_custom_fields":                  &iam_model.IAMCustomField{},
		"positions":                          &iam_model.Position{},
		"supply_asset_groups":                &iam_model.SupplyAssetGroup{},
		"users":                              &iam_model.User{},
		"user_auths":                         &iam_model.UserAuth{},
		"user_custom_fields":                 &iam_model.UserCustomField{},
	},
	db.HardwareTaggingDBName: {
		"device_names":       &hardware_tagging_model.DeviceName{},
		"handheld_readers":   &hardware_tagging_model.HandheldReader{},
		"handheld_stocks":    &hardware_tagging_model.HandheldStock{},
		"nfcs":               &hardware_tagging_model.NFC{},
		"nfc_stocks":         &hardware_tagging_model.NfcStock{},
		"qrs":                &hardware_tagging_model.QR{},
		"rfids":              &hardware_tagging_model.RFID{},
		"rfid_stocks":        &hardware_tagging_model.RfidStock{},
		"skus":               &hardware_tagging_model.Sku{},
		"smart_fixed_stocks": &hardware_tagging_model.SmartFixedStock{},
		"tablets":            &hardware_tagging_model.Tablet{},
		"tablet_stocks":      &hardware_tagging_model.TabletStock{},
		"tag_pairs":          &hardware_tagging_model.TagPair{},
	},
	db.MasterAssetRegistryDBName: {
		"assets":              &master_asset_registry_model.Asset{},
		"asset_custom_fields": &master_asset_registry_model.AssetCustomField{},
		"custom_fields":       &master_asset_registry_model.CustomField{},
	},
	db.AssetTransactionDBName: {
		"fixed_asset_disposal_transactions": &model.FixedAssetDisposalTransaction{},
	},
	db.EnterpriseManagementDBName: {
		"clients": &ems_model.Client{},
	},
}

// ReplicatedTables is the central registry that declares which tables each service
// needs to replicate (consume) from other services (publishers).
var ReplicatedTables = map[string]map[api.ReplicationKey][]api.ReplicatedTable{
	constant.IAM_SERVICE_NAME: { // Listener/Consumer
		api.ReplicationKey{Flow: api.ReplicationFlowTenantToTenant, SourceDBName: db.HardwareTaggingDBName}: { // Publisher
			&iam_replica.Rfid{},
		},
	},
	constant.HARDWARE_TAGGING_SERVICE_NAME: { // Listener/Consumer
		api.ReplicationKey{Flow: api.ReplicationFlowNonTenantToTenant, SourceDBName: db.HardwareTaggingDBName}: { // Publisher
			&hardware_tagging_replica.DeviceName{},
			&hardware_tagging_replica.HandheldStock{},
			&hardware_tagging_replica.Sku{},
			&hardware_tagging_replica.TabletStock{},
			&hardware_tagging_replica.SmartFixedStock{},
			&hardware_tagging_replica.RfidStock{},
			&hardware_tagging_replica.NfcStock{},
		},
		api.ReplicationKey{Flow: api.ReplicationFlowTenantToTenant, SourceDBName: db.IAMDBName}: { // Publisher
			&hardware_tagging_replica.FixedAssetGroup{},
			&hardware_tagging_replica.SupplyAssetGroup{},
			&hardware_tagging_replica.User{},
		},
	},
	constant.ENTERPRISE_MANAGEMENT_SERVICE_NAME: nil,
	constant.MASTER_ASSET_REGISTRY_SERVICE_NAME: nil,
	constant.ANALYTICS_REPORTING_SERVICE_NAME:   nil,
	constant.APPROVAL_SERVICE_NAME:              nil,
	constant.FILE_MANAGER_SERVICE_NAME: { // Listener/Consumer
		api.ReplicationKey{Flow: api.ReplicationFlowTenantToTenant, SourceDBName: db.MasterAssetRegistryDBName}: { // Publisher
			&replica.AssetCustomField{},
			&replica.Asset{},
			&replica.CustomField{},
		},
		api.ReplicationKey{Flow: api.ReplicationFlowTenantToTenant, SourceDBName: db.AssetTransactionDBName}: { // Publisher
			&replica.FixedAssetDisposalTransaction{},
		},
		api.ReplicationKey{Flow: api.ReplicationFlowTenantToTenant, SourceDBName: db.IAMDBName}: { // Publisher
			&replica.FixedAssetGroup{},
			&replica.User{},
		},
	},
	constant.IMPORT_SERVICE_NAME:               nil,
	constant.LANGUAGE_TRANSLATION_SERVICE_NAME: nil,
	constant.NOTIFICATION_SERVICE_NAME:         nil,
	constant.PLANNING_STRATEGY_SERVICE_NAME:    nil,
	constant.RESOURCE_ALLOCATION_SERVICE_NAME:  nil,
	constant.ASSET_ADDON_SERVICE_NAME:          nil,
	constant.ASSET_TRANSACTION_SERVICE_NAME:    nil,
}
