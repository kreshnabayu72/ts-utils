package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type CustomFieldCategory struct {
	ID            int       `json:"_id"           column:"custom_field_categories.id"              sql:"BIGINT PRIMARY KEY"`
	CustomFieldId int       `json:"customFieldId" column:"custom_field_categories.custom_field_id" sql:"BIGINT"`
	CategoryId    int       `json:"categoryId"    column:"custom_field_categories.category_id"     sql:"BIGINT"`
	CreatedAt     time.Time `json:"createdAt"     column:"custom_field_categories.created_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time `json:"updatedAt"     column:"custom_field_categories.updated_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

	// Add the unique index for customfield id and category id
	builder.AddIndex(
		"unique_custom_field_category",
		"custom_field_id,category_id",
		true,
		"",
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
