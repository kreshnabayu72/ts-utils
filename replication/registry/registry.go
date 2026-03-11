package registry

import (
	"fmt"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/file_manager_service/replica"
	hardware_tagging_replica "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/hardware_tagging_service/replica"
	iam_replica "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/iam_service/replica"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/api"
)

// ReplicatedTables is the central registry that declares which tables each service
// needs to replicate (consume) from other services (publishers).
var ReplicatedTables = map[string]map[string][]api.ReplicatedTable{
	constant.IAM_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", api.ReplicationFlowTenantToTenant, db.HardwareTaggingDBName): { // Publisher
			&iam_replica.Rfid{},
		},
	},
	constant.HARDWARE_TAGGING_SERVICE_NAME: { // Listener/Consumer
		fmt.Sprintf("%v.%s", api.ReplicationFlowNonTenantToTenant, db.HardwareTaggingDBName): { // Publisher
			&hardware_tagging_replica.DeviceName{},
			&hardware_tagging_replica.HandheldStock{},
			&hardware_tagging_replica.Sku{},
			&hardware_tagging_replica.TabletStock{},
			&hardware_tagging_replica.SmartFixedStock{},
			&hardware_tagging_replica.RfidStock{},
			&hardware_tagging_replica.NfcStock{},
		},
		fmt.Sprintf("%v.%s", api.ReplicationFlowTenantToTenant, db.IAMDBName): { // Publisher
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
		fmt.Sprintf("%v.%s", api.ReplicationFlowTenantToTenant, db.MasterAssetRegistryDBName): { // Publisher
			&replica.AssetCustomField{},
			&replica.Asset{},
			&replica.CustomField{},
		},
		fmt.Sprintf("%v.%s", api.ReplicationFlowTenantToTenant, db.AssetTransactionDBName): { // Publisher
			&replica.FixedAssetDisposalTransaction{},
		},
		fmt.Sprintf("%v.%s", api.ReplicationFlowTenantToTenant, db.IAMDBName): { // Publisher
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
