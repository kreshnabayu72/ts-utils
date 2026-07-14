// Package dto defines the iam-service cross-service activity parameters.
package dto

// DeleteAllUserAuthByClientIDActivityParam holds the client identifier to hard-delete.
// Activity is written in iam-service.
type DeleteAllUserAuthByClientIDActivityParam struct {
	ClientID string
}
