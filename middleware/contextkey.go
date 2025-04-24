package middleware

// for using context value, saved in utils so every key come from 1 place
type ContextKey string

const (
	UserKey        = ContextKey("user")
	CompanyCodeKey = ContextKey("companyCode")
	SessionKey     = ContextKey("session")
	Company        = ContextKey("company")
	ReaderId       = ContextKey("readerId")
	FixedAssetPlan = ContextKey("fixedAssetPlan")

	//Only used for GetLicensePlan middleware cause idk how to mock the service in middleware
	IsTest = ContextKey("isTest")
)
