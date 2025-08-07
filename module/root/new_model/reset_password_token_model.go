package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ResetPasswordToken struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Token       string `json:"token" column:"token" sql:"VARCHAR UNIQUE NOT NULL"`
	Email       string `json:"email" column:"email" sql:"VARCHAR NOT NULL"`
	CompanyCode string `json:"companyCode" column:"company_code" sql:"VARCHAR NOT NULL"`

	DeletedAt *time.Time `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt *time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *ResetPasswordToken) TableName() string {
	return "reset_password_tokens"
}

func (m *ResetPasswordToken) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
