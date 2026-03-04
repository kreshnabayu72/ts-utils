package dto

type GetClientParam struct {
	ClientID string
}

type UpdateClientActivityParam struct {
	ClientID string
	Email    string
}

type UndoUpdateClientActivityParam struct {
	ClientID string
	Email    string
}
