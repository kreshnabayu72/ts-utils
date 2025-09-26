package new_model

import (
	"time"
)

type UserAuth struct {
	ID          int        `json:"_id"         column:"user_auths.id"`
	Email       string     `json:"email"       column:"user_auths.email"`
	CompanyCode string     `json:"companyCode" column:"user_auths.company_code"`
	Password    string     `json:"password"    column:"user_auths.password"`
	CreatedAt   *time.Time `json:"createdAt"   column:"user_auths.created_at"`
	UpdatedAt   *time.Time `json:"updatedAt"   column:"user_auths.updated_at"`
}
