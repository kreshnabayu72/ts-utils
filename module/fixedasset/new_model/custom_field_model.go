package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type CustomField struct {
	ID            int                 `json:"_id"           column:"custom_fields.id"               sql:"BIGINT PRIMARY KEY"`
	Name          string              `json:"name"          column:"custom_fields.name"             sql:"VARCHAR NOT NULL"`
	Type          CustomFieldType     `json:"type"          column:"custom_fields.type"             sql:"custom_field_type_enum NOT NULL"`
	DataType      CustomFieldDataType `json:"dataType"      column:"custom_fields.data_type"        sql:"custom_field_data_type_enum NOT NULL"`
	IsRequired    bool                `json:"isRequired"    column:"custom_fields.is_required"      sql:"BOOLEAN NOT NULL"`
	OptionValue   []string            `json:"optionValue"   column:"custom_fields.option_value"     sql:"TEXT[] NULL"`
	IsActive      bool                `json:"isActive"      column:"custom_fields.is_active"        sql:"BOOLEAN DEFAULT TRUE"`
	IsShowOnTable bool                `json:"isShowOnTable" column:"custom_fields.is_show_on_table" sql:"BOOLEAN DEFAULT TRUE"`
	IsDeleted     bool                `json:"isDeleted"     column:"custom_fields.is_deleted"       sql:"BOOLEAN DEFAULT FALSE"`
	DeletedAt     *time.Time          `json:"deletedAt"     column:"custom_fields.deleted_at"       sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt     time.Time           `json:"createdAt"     column:"custom_fields.created_at"       sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time           `json:"updatedAt"     column:"custom_fields.updated_at"       sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *CustomField) TableName() string {
	return "custom_fields"
}

func (m *CustomField) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("custom_fields_name_lower_idx", "LOWER(name)", true, "is_deleted = FALSE")
	builder.AddIndex("idx_cf_active", "is_active", false, "is_deleted = FALSE")
	builder.AddIndex("idx_cf_type_active", "type", false, "is_deleted = FALSE")
	builder.AddIndex("idx_cf_datatype_active", "data_type", false, "is_deleted = FALSE")
	builder.AddIndex("idx_cf_created_at_active", "created_at DESC", false, "is_deleted = FALSE")
}
