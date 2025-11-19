package model

import (
	"time"
)

type UserToken struct {
	ID          int           `json:"_id"         column:"user_tokens.id"`
	Type        UserTokenType `json:"type"       column:"user_tokens.type"`
	Token       string        `json:"token"       column:"user_tokens.token"`
	Email       string        `json:"email"       column:"user_tokens.email"`
	CompanyCode string        `json:"companyCode" column:"user_tokens.company_code"`
	DeletedAt   *time.Time    `json:"deletedAt"   column:"user_tokens.deleted_at"`
	CreatedAt   *time.Time    `json:"createdAt"   column:"user_tokens.created_at"`
	UpdatedAt   *time.Time    `json:"updatedAt"   column:"user_tokens.updated_at"`
}

type UserTokenType string

const (
	UserTokenTypeEmailChange      = UserTokenType("email_change")
	UserTokenTypeOtp              = UserTokenType("otp")
	UserTokenTypePasswordChange   = UserTokenType("password_change")
	UserTokenTypeUserConfirmation = UserTokenType("user_confirmation")
)
