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
)
