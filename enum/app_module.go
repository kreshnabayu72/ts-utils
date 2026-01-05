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

var ModuleAccessKey = map[ModuleAccess]string{
	ModuleAccessFixedAsset:     "fixed-asset",
	ModuleAccessGlobalSettings: "global-settings",
	ModuleAccessAdminConsole:   "admin-console",
	ModuleAccessSupplyAsset:    "supply-asset",
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
