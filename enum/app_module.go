package enum

type ModuleAccess string

const (
	ModuleAccessFixedAsset     = ModuleAccess("Fixed Asset")
	ModuleAccessGlobalSettings = ModuleAccess("Global Settings")
	ModuleAccessAdminConsole   = ModuleAccess("Admin Console")
	ModuleAccessSupplyAsset    = ModuleAccess("Supply Asset")
)

var ModuleAccessEnum = map[string]ModuleAccess{
	"Fixed Asset":     ModuleAccessFixedAsset,
	"Global Settings": ModuleAccessGlobalSettings,
	"Admin Console":   ModuleAccessAdminConsole,
	"Supply Asset":    ModuleAccessSupplyAsset,
}

type PlanTypes string

const (
	PlanBasic      = PlanTypes("Basic")
	PlanStandard   = PlanTypes("Standard")
	PlanEnterprise = PlanTypes("Enterprise")
)

var PlanEnum = map[string]PlanTypes{
	"Basic":      PlanBasic,
	"Standard":   PlanStandard,
	"Enterprise": PlanEnterprise,
}
