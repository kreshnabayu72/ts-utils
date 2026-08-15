package enum

// Aliases of the shared key/name types — see system_roles_category.go. Global Settings
// grants are always category Global: the module has no groups, so no scope FK applies.
type GlobalSettingsSystemRoleKey = SystemRoleKey

const (
	GlobalSettingsSystemRoleKeyManageUserAndRole     GlobalSettingsSystemRoleKey = "user_and_role"
	GlobalSettingsSystemRoleKeyManageMasterReference GlobalSettingsSystemRoleKey = "master_reference"
	GlobalSettingsSystemRoleKeyManageHardware        GlobalSettingsSystemRoleKey = "hardware"
	GlobalSettingsSystemRoleKeyImportData            GlobalSettingsSystemRoleKey = "import_data"
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
