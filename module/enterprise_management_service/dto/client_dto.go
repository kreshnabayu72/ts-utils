package dto

type GetClientParam struct {
	ClientID string
}

type UpdateClientActivityParam struct {
	ClientID string
	Email    string
}

type RollbackUpdateClientActivityParam struct {
	ClientID string
	Email    string
}
