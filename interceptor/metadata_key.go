package interceptor

// MetadataKey used for context value, saved in utils so every key come from 1 place
type MetadataKey string

const (
	UserKey        = MetadataKey("user")
	CompanyCodeKey = MetadataKey("company-code")
	SessionKey     = MetadataKey("session")
	Company        = MetadataKey("company")
)
