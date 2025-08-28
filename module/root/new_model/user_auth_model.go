package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	root_migration "gitea.qwertysystem.net/BETS/ts-utils/module/root/new_model/migration"
)

type UserAuth struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Email       string `json:"email" column:"email" sql:"VARCHAR NOT NULL"`
	CompanyCode string `json:"companyCode" column:"company_code" sql:"VARCHAR NOT NULL"`
	Password    string `json:"password" column:"password" sql:"VARCHAR NOT NULL"`

	CreatedAt *time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *UserAuth) TableName() string {
	return "user_auths"
}

func (m *UserAuth) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("user_auths_email_uindex", "email", true, "")
	builder.AddIndex("user_auths_company_code_uindex", "company_code", true, "")

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(root_migration.UserAuthPostQuery)
}
