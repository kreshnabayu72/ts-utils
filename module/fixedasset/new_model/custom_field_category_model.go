package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type CustomFieldCategory struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	CustomFieldId int `json:"customFieldId" column:"custom_field_id" sql:"BIGINT UNIQUE"`
	CategoryId    int `json:"categoryId" column:"category_id" sql:"BIGINT UNIQUE"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (c *CustomFieldCategory) TableName() string {
	return "custom_field_categories"
}

func (c *CustomFieldCategory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(c)

	builder.AddForeignKey(
		"fk_custom_field_categories_custom_field_id",
		"custom_field_id",
		"custom_fields",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_custom_field_categories_category_id",
		"category_id",
		"categories",
		"id",
		"SET NULL",
	)

	builder.AddIndex("idx_cfc_custom_field_active", "custom_field_id", false, "")
	builder.AddIndex("idx_cfc_category_active", "category_id", false, "")
	builder.AddIndex(
		"idx_cfc_custom_field_category_active",
		"custom_field_id,category_id",
		false,
		"",
	)
}
