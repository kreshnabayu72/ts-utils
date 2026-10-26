package middleware

// for using context value, saved in utils so every key come from 1 place
type ContextKey string

const (
	UserKey           ContextKey = "user"
	ClientIdKey       ContextKey = "clientId"
	SessionKey        ContextKey = "session"
	ClientKey         ContextKey = "client"
	ReaderIdKey       ContextKey = "readerId"
	FixedAssetPlanKey ContextKey = "fixedAssetPlan"
	SupplyPlanKey     ContextKey = "supplyPlan"
	UserRolesKey      ContextKey = "userFixedAssetRoles"
	EntitySettingsKey ContextKey = "entityAndSettings"
)
