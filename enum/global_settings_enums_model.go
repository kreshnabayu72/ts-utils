package enum

type GlobalSettingsSystemRoleKey string

const (
	GlobalSettingsSystemRoleKeyManageUserAndRole     GlobalSettingsSystemRoleKey = "userAndRole"
	GlobalSettingsSystemRoleKeyManageMasterReference GlobalSettingsSystemRoleKey = "masterReference"
	GlobalSettingsSystemRoleKeyManageHardware        GlobalSettingsSystemRoleKey = "hardware"
	GlobalSettingsSystemRoleKeyImportData            GlobalSettingsSystemRoleKey = "importData"
	GlobalSettingsSystemRoleKeyManageFiles           GlobalSettingsSystemRoleKey = "files"
	GlobalSettingsSystemRoleKeyManageTAG             GlobalSettingsSystemRoleKey = "tag"
)

type GlobalSettingsSystemRoleName string

const (
	GlobalSettingsSystemRoleNameManageUserAndRole     GlobalSettingsSystemRoleName = "Manage user and role"
	GlobalSettingsSystemRoleNameManageMasterReference GlobalSettingsSystemRoleName = "Manage master reference"
	GlobalSettingsSystemRoleNameManageHardware        GlobalSettingsSystemRoleName = "Manage hardware"
	GlobalSettingsSystemRoleNameImportData            GlobalSettingsSystemRoleName = "Import data"
	GlobalSettingsSystemRoleNameManageFiles           GlobalSettingsSystemRoleName = "Manage files"
	GlobalSettingsSystemRoleNameManageTAG             GlobalSettingsSystemRoleName = "Manage TAG"
)
