package model

import "time"

type UserAuth struct {
	ID        int        `json:"_id"       column:"user_auths.id"`
	Email     string     `json:"email"     column:"user_auths.email"`
	ClientId  string     `json:"clientId"  column:"user_auths.client_id"`
	Password  *string    `json:"password"  column:"user_auths.password"`
	IsDeleted bool       `json:"isDeleted" column:"user_auths.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"user_auths.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"user_auths.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"user_auths.updated_at"`
}

func (r *UserAuth) GetColumns() []string {
	return []string{
		"id",
		"email",
		"client_id",
		"password",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}
