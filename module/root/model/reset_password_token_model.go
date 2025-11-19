package model

import (
	"time"
)

type ResetPasswordToken struct {
	ID          int        `json:"_id"         column:"reset_password_tokens.id"`
	Token       string     `json:"token"       column:"reset_password_tokens.token"`
	Email       string     `json:"email"       column:"reset_password_tokens.email"`
	CompanyCode string     `json:"companyCode" column:"reset_password_tokens.company_code"`
	DeletedAt   *time.Time `json:"deletedAt"   column:"reset_password_tokens.deleted_at"`
	CreatedAt   *time.Time `json:"createdAt"   column:"reset_password_tokens.created_at"`
	UpdatedAt   *time.Time `json:"updatedAt"   column:"reset_password_tokens.updated_at"`
}
