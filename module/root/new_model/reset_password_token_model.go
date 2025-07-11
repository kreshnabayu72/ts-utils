package new_model

import (
	"time"
)

type ResetPasswordToken struct {
	ID int `json:"_id"             column:"id"`

	Token       string `json:"token"           column:"token"`
	Email       string `json:"email"       column:"email"`
	CompanyCode string `json:"companyCode" column:"company_code"`

	CreatedAt *time.Time `json:"createdAt"       column:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at"`
}
