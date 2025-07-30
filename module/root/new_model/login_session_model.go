package new_model

import (
	"time"
)

type LoginSession struct {
	ID              int        `json:"_id,omitempty"        column:"id"`
	UserCompanyCode string     `json:"userCompanyCode"      column:"user_company_code"`
	UserEmail       string     `json:"userEmail"            column:"user_email"`
	DeletedAt       *time.Time `json:"deletedAt" column:"deleted_at"`
	SessionType     string     `json:"sessionType"          column:"session_type"` // Possible Values: "mobile", "web"
	CreatedAt       *time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt       *time.Time `json:"updatedAt"            column:"updated_at"`
}
