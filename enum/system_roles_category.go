package enum

type ScopeLevel string

const (
	ScopeLevelGlobal ScopeLevel = "global"
	ScopeLevelGroup  ScopeLevel = "group"
)

type SystemRoleKey string

const (
	//Common
	SystemRoleKeyManageUserAndRole SystemRoleKey = "manage_user_and_role"
	SystemRoleKeyManageGroup       SystemRoleKey = "manage_group"

	//Global settings
	SystemRoleKeyManageMasterReference SystemRoleKey = "manage_master_reference"
	SystemRoleKeyManageHardware        SystemRoleKey = "manage_hardware"
	SystemRoleKeyImportData            SystemRoleKey = "import_data"
	SystemRoleKeyManageFiles           SystemRoleKey = "manage_files"
	SystemRoleKeyManageTAG             SystemRoleKey = "manage_tag"

	//Admin console
	SystemRoleKeyManageDeviceMasterData  SystemRoleKey = "manage_device_master_data"
	SystemRoleKeyManageTAGStock          SystemRoleKey = "manage_tag_stock"
	SystemRoleKeyManageHardwareStock     SystemRoleKey = "manage_hardware_stock"
	SystemRoleKeyManageLanguage          SystemRoleKey = "manage_language"
	SystemRoleKeyManagePresetup          SystemRoleKey = "manage_presetup"
	SystemRoleKeyManagePreload           SystemRoleKey = "manage_preload"
	SystemRoleKeyManagePlansAndLicensing SystemRoleKey = "manage_plans_and_licensing"
	SystemRoleKeyManagePrivacyPolicy     SystemRoleKey = "manage_privacy_policy"
	SystemRoleKeyManageDistributor       SystemRoleKey = "manage_distributor"
	SystemRoleKeyManagePartner           SystemRoleKey = "manage_partner"
	SystemRoleKeyManageActiveClient      SystemRoleKey = "manage_active_client"
	SystemRoleKeyManageArchivedClient    SystemRoleKey = "manage_archived_client"
	SystemRoleKeyManageAllocation        SystemRoleKey = "manage_allocation"
	SystemRoleKeyManageWalletAndPricing  SystemRoleKey = "manage_wallet_and_pricing"

	//Fixed asset
	SystemRoleKeyManageAssetMasterData          SystemRoleKey = "manage_asset_master_data"
	SystemRoleKeyManageAssetTransactionalPolicy SystemRoleKey = "manage_asset_transactional_policy"
	SystemRoleKeyManageAccounting               SystemRoleKey = "manage_accounting"
	SystemRoleKeyManageTransactionSettings      SystemRoleKey = "manage_transaction_settings"
	SystemRoleKeyManageOpenAPI                  SystemRoleKey = "manage_open_api"
	SystemRoleKeyManageAssetIDConfiguration     SystemRoleKey = "manage_asset_id_configuration"
	SystemRoleKeyManageGlobalLicense            SystemRoleKey = "manage_global_license"
	SystemRoleKeyManageGlobalUserAndRole        SystemRoleKey = "manage_global_user_and_role"
	SystemRoleKeyManageAssetAndTAG              SystemRoleKey = "manage_asset_and_tag"
	SystemRoleKeyManageAssetLicense             SystemRoleKey = "manage_asset_license"
	SystemRoleKeyManagePurchase                 SystemRoleKey = "manage_purchase"
	SystemRoleKeyViewAssetValue                 SystemRoleKey = "view_asset_value"
)

type SystemRoleName string

const (
	//Common
	SystemRoleNameManageUserAndRole SystemRoleName = "Manage user and role"
	SystemRoleNameManageGroup       SystemRoleName = "Manage group"

	//Global settings
	SystemRoleNameManageMasterReference SystemRoleName = "Manage master reference"
	SystemRoleNameManageHardware        SystemRoleName = "Manage hardware"
	SystemRoleNameImportData            SystemRoleName = "Import data"
	SystemRoleNameManageFiles           SystemRoleName = "Manage files"
	SystemRoleNameManageTAG             SystemRoleName = "Manage TAG"

	//Admin console
	SystemRoleNameManageDeviceMasterData  SystemRoleName = "Manage device master data"
	SystemRoleNameManageTAGStock          SystemRoleName = "Manage TAG stock"
	SystemRoleNameManageHardwareStock     SystemRoleName = "Manage hardware stock"
	SystemRoleNameManageLanguage          SystemRoleName = "Manage language"
	SystemRoleNameManagePresetup          SystemRoleName = "Manage presetup"
	SystemRoleNameManagePreload           SystemRoleName = "Manage preload"
	SystemRoleNameManagePlansAndLicensing SystemRoleName = "Manage plans & licensing"
	SystemRoleNameManagePrivacyPolicy     SystemRoleName = "Manage privacy policy"
	SystemRoleNameManageDistributor       SystemRoleName = "Manage distributor"
	SystemRoleNameManagePartner           SystemRoleName = "Manage partner"
	SystemRoleNameManageActiveClient      SystemRoleName = "Manage active client"
	SystemRoleNameManageArchivedClient    SystemRoleName = "Manage archived client"
	SystemRoleNameManageAllocation        SystemRoleName = "Manage allocation"
	SystemRoleNameManageWalletAndPricing  SystemRoleName = "Manage wallet & pricing"

	//Fixed asset
	SystemRoleNameManageAssetMasterData          SystemRoleName = "Manage asset master data"
	SystemRoleNameManageAssetTransactionalPolicy SystemRoleName = "Manage asset transactional policy"
	SystemRoleNameManageAccounting               SystemRoleName = "Manage accounting"
	SystemRoleNameManageTransactionSettings      SystemRoleName = "Manage transaction settings"
	SystemRoleNameManageOpenAPI                  SystemRoleName = "Manage Open API"
	SystemRoleNameManageAssetIDConfiguration     SystemRoleName = "Manage asset ID configuration"
	SystemRoleNameManageGlobalLicense            SystemRoleName = "Manage global license"
	SystemRoleNameManageGlobalUserAndRole        SystemRoleName = "Manage global user & role"
	SystemRoleNameManageAssetAndTAG              SystemRoleName = "Manage asset & TAG"
	SystemRoleNameManageAssetLicense             SystemRoleName = "Manage asset license"
	SystemRoleNameManagePurchase                 SystemRoleName = "Manage purchase"
	SystemRoleNameViewAssetValue                 SystemRoleName = "View asset value"
)
