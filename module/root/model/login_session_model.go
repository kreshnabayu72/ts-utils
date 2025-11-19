package model

import (
	"time"
)

type LoginSession struct {
	ID          int        `json:"_id,omitempty" column:"login_sessions.id"`
	CompanyCode string     `json:"companyCode"   column:"login_sessions.company_code"`
	Email       string     `json:"email"         column:"login_sessions.email"`
	DeletedAt   *time.Time `json:"deletedAt"     column:"login_sessions.deleted_at"`
	SessionType string     `json:"sessionType"   column:"login_sessions.session_type"` // Possible Values: "mobile", "web"
	CreatedAt   *time.Time `json:"createdAt"     column:"login_sessions.created_at"`
	UpdatedAt   *time.Time `json:"updatedAt"     column:"login_sessions.updated_at"`
}
