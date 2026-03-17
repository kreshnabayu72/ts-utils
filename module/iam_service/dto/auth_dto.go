package dto

// DeleteAllUserAuthByClientIDActivityParam holds the client identifier to hard-delete.
// Activity is written in iam-service.
type DeleteAllUserAuthByClientIDActivityParam struct {
	ClientID string
}
