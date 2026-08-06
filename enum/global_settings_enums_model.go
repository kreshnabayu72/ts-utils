package enum

// Aliases of the shared key/name types — see FixedAssetSystemRoleKey. Global Settings
// grants are always category Global: the module has no groups, so no scope FK applies.
type GlobalSettingsSystemRoleKey = SystemRoleKey

const (
	GlobalSettingsSystemRoleKeyManageUserAndRole     GlobalSettingsSystemRoleKey = "userAndRole"
	GlobalSettingsSystemRoleKeyManageMasterReference GlobalSettingsSystemRoleKey = "masterReference"
	GlobalSettingsSystemRoleKeyManageHardware        GlobalSettingsSystemRoleKey = "hardware"
	GlobalSettingsSystemRoleKeyImportData            GlobalSettingsSystemRoleKey = "importData"
	GlobalSettingsSystemRoleKeyManageFiles           GlobalSettingsSystemRoleKey = "files"
	GlobalSettingsSystemRoleKeyManageTAG             GlobalSettingsSystemRoleKey = "tag"
)

type GlobalSettingsSystemRoleName = SystemRoleName

const (
	GlobalSettingsSystemRoleNameManageUserAndRole     GlobalSettingsSystemRoleName = "Manage user and role"
	GlobalSettingsSystemRoleNameManageMasterReference GlobalSettingsSystemRoleName = "Manage master reference"
	GlobalSettingsSystemRoleNameManageHardware        GlobalSettingsSystemRoleName = "Manage hardware"
	GlobalSettingsSystemRoleNameImportData            GlobalSettingsSystemRoleName = "Import data"
	GlobalSettingsSystemRoleNameManageFiles           GlobalSettingsSystemRoleName = "Manage files"
	GlobalSettingsSystemRoleNameManageTAG             GlobalSettingsSystemRoleName = "Manage TAG"
)
