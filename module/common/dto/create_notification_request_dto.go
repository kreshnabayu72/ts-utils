package dto

// CreateNotificationRequest is the body of POST /internal/notifications.
// Will be called by all notification senders
type CreateNotificationRequest struct {
	UserID string         `json:"userId" validate:"required"`
	Title  string         `json:"title"  validate:"required"`
	Detail string         `json:"detail"`
	Module string         `json:"module"`
	Data   map[string]any `json:"data"`
}
