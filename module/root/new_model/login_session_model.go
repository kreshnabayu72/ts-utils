package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type LoginSession struct {
	ID              int        `json:"_id,omitempty" column:"id" sql:"BIGINT PRIMARY KEY"`
	UserCompanyCode string     `json:"userCompanyCode" column:"user_company_code" sql:"VARCHAR NOT NULL"`
	UserEmail       string     `json:"userEmail" column:"user_email" sql:"VARCHAR NOT NULL"`
	DeletedAt       *time.Time `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMP NULL DEFAULT NULL"`
	SessionType     string     `json:"sessionType" column:"session_type" sql:"session_type_enum NULL"` // Possible Values: "mobile", "web"
	CreatedAt       *time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt       *time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (m *LoginSession) TableName() string {
	return "login_sessions"
}

func (m *LoginSession) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddEnum("session_type_enum", []string{"Mobile", "Web"})
}
