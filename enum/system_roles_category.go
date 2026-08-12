package enum

// SystemRoleKey and SystemRoleName are the key/name types of the merged
// user_system_roles table. Every module's grants live in that one table, so the
// column holds keys from all of them and one Go type has to cover them all.
//
// A grant is identified by (user_id, module, key, group id),
// so always constrain Module before matching a Key.
type SystemRoleKey string

const (
	SystemRoleKeyManageUserAndRole     SystemRoleKey = "manageUserAndRole"
	SystemRoleKeyManageMasterReference SystemRoleKey = "manageMasterReference"
	SystemRoleKeyManageHardware        SystemRoleKey = "manageHardware"
	SystemRoleKeyImportData            SystemRoleKey = "importData"
	SystemRoleKeyManageFiles           SystemRoleKey = "manageFiles"
	SystemRoleKeyManageTAG             SystemRoleKey = "manageTAG"
)

type SystemRoleName string

const (
	SystemRoleNameManageUserAndRole      SystemRoleName = "Manage user and role"
	SystemRoleNameManageMasterReference  SystemRoleName = "Manage master reference"
	SystemRoleNameManageHardware         SystemRoleName = "Manage hardware"
	SystemRoleNameImportData             SystemRoleName = "Import data"
	SystemRoleNameManageFiles            SystemRoleName = "Manage files"
	SystemRoleNameManageTAG              SystemRoleName = "Manage TAG"
)

// SystemRolesCategory reports which arm of the scope arc a grant uses: Group means a
// scope FK is set, Global means every scope FK is NULL and the grant is tenant-wide
// within its module.
type SystemRolesCategory string

const (
	SystemRolesCategoryGroup  SystemRolesCategory = "Group"
	SystemRolesCategoryGlobal SystemRolesCategory = "Global"
)
