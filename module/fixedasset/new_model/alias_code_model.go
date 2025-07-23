package new_model

import (
	"time"
)

type AliasCodeData struct {
	Name      string `json:"name" column:"name"`
	IsChecked bool   `json:"isChecked" column:"is_checked"`
	Value     string `json:"value" column:"value"`
}

type AliasCode struct {
	ID             int             `column:"id" json:"_id,omitempty"`
	IsActive       bool            `column:"is_active" json:"isActive"`
	CompanyInitial string          `column:"company_initial" json:"companyInitial"`
	Separator      string          `column:"separator" json:"separator"`
	Data           []AliasCodeData `json:"data"`
	CreatedAt      time.Time       `column:"created_at" json:"createdAt"`
	UpdatedAt      time.Time       `column:"updated_at" json:"updatedAt"`
}
