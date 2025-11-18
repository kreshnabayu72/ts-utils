package entity

// UseCase is the component of the core application's business logic.
type UseCase[P any, R any] interface {
	Invoke(param P) (R, error)
	InitService(companyCode string)
}
