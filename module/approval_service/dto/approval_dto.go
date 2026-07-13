// Package dto defines the approval-service activity parameters.
package dto

type GetUsersPendingApprovalCountParam struct {
	ClientID string
	UserID   string
	// Plan        string
}
