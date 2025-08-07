package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AliasCodeData struct {
	ID          int       `column:"acd.id" json:"_id,omitempty" sql:"BIGINT PRIMARY KEY"`
	AliasCodeID int       `column:"acd.alias_code_id" json:"aliasCodeId" sql:"BIGINT NOT NULL"`
	Name        string    `json:"name" column:"acd.name" sql:"VARCHAR NOT NULL CHECK (char_length(name) > 0)"`
	IsChecked   bool      `json:"isChecked" column:"acd.is_checked" sql:"BOOLEAN NOT NULL DEFAULT FALSE CHECK (is_checked IN (TRUE, FALSE))"`
	Value       string    `json:"value" column:"acd.value" sql:"VARCHAR NOT NULL"`
	Order       *int      `json:"order" column:"acd.order" sql:"INTEGER NULL UNIQUE"`
	CreatedAt   time.Time `column:"acd.created_at" json:"createdAt" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt   time.Time `column:"acd.updated_at" json:"updatedAt" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *AliasCodeData) TableName() string {
	return "alias_codes_data"
}

func (m *AliasCodeData) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey("fk_alias_code", "alias_code_id", "alias_codes", "id", "RESTRICT")

	builder.AddIndex("idx_alias_codes_data_alias_code_id", "alias_code_id", false, "")
	builder.AddIndex("idx_alias_codes_data_name", "name", false, "")
	builder.AddIndex("idx_alias_codes_data_is_checked", "is_checked", false, "")
}

type AliasCode struct {
	ID             int             `column:"ac.id" json:"_id,omitempty" sql:"BIGINT PRIMARY KEY"`
	IsActive       bool            `column:"ac.is_active" json:"isActive" sql:"BOOLEAN NOT NULL DEFAULT FALSE CHECK (is_active IN (TRUE, FALSE))"`
	CompanyInitial string          `column:"ac.company_initial" json:"companyInitial" sql:"VARCHAR NOT NULL CHECK (char_length(company_initial) > 0)"`
	Separator      string          `column:"ac.separator" json:"separator" sql:"VARCHAR NOT NULL CHECK (char_length(separator) > 0)"`
	Data           []AliasCodeData `json:"data"`
	CreatedAt      time.Time       `column:"ac.created_at" json:"createdAt" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time       `column:"ac.updated_at" json:"updatedAt" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *AliasCode) TableName() string {
	return "alias_codes"
}

func (m *AliasCode) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("idx_alias_codes_company_initial", "company_initial", false, "")
	builder.AddIndex("idx_alias_codes_separator", "separator", false, "")
}
