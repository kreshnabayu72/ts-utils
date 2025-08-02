package new_model

type CompanyStatus string

const (
	CompanyStatusActive   CompanyStatus = "Active"
	CompanyStatusExpired  CompanyStatus = "Expired"
	CompanyStatusArchived CompanyStatus = "Archived"
)

type CompanyOpenAPIStatus string

const (
	CompanyOpenAPIStatusApproved CompanyOpenAPIStatus = "Approved"
	CompanyOpenAPIStatusWaiting  CompanyOpenAPIStatus = "Waiting for Approval"
)

type CompanyDepreciationMethod string

const (
	CompanyDepreciationMethodStraightLine           CompanyDepreciationMethod = "Straight Line"
	CompanyDepreciationMethodDoubleDecliningBalance CompanyDepreciationMethod = "Double-Declining Balance"
)
