package middleware

// for using context value, saved in utils so every key come from 1 place
type ContextKey string

const (
	ErrorLog           = ContextKey("errorLog")
	UserKey            = ContextKey("user")
	CompanyCodeKey     = ContextKey("companyCode")
	SessionKey         = ContextKey("session")
	Company            = ContextKey("company")
	ReaderId           = ContextKey("readerId")
	FixedAssetPlan     = ContextKey("fixedAssetPlan")
	SupplyPlan         = ContextKey("supplyPlan")
	FileSize           = ContextKey("fileSize")
	FileSizeString     = ContextKey("fileSizeString")
	FileNameWithoutExt = ContextKey("fileNameWithoutExt")
	FileExt            = ContextKey("fileExt")
)
