package new_model

import (
	"time"
)

type UserToken struct {
	ID          int        `json:"_id"         column:"user_tokens.id"`
	Token       string     `json:"token"       column:"user_tokens.token"`
	Email       string     `json:"email"       column:"user_tokens.email"`
	CompanyCode string     `json:"companyCode" column:"user_tokens.company_code"`
	DeletedAt   *time.Time `json:"deletedAt"   column:"user_tokens.deleted_at"`
	CreatedAt   *time.Time `json:"createdAt"   column:"user_tokens.created_at"`
	UpdatedAt   *time.Time `json:"updatedAt"   column:"user_tokens.updated_at"`
}
