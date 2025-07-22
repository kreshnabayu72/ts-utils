package new_model

import "time"

type CustomField struct {
	ID int `json:"_id" column:"id"`

	Name          string              `json:"name"         column:"name"`
	Type          CustomFieldType     `json:"type"         column:"type"`
	DataType      CustomFieldDataType `json:"dataType"    column:"data_type"`
	IsRequired    bool                `json:"isRequired"  column:"is_required"`
	OptionValue   []string            `json:"optionValue" column:"option_value"`
	IsActive      bool                `json:"isActive"  column:"is_active"`
	IsShowOnTable bool                `json:"isShowOnTable"  column:"is_show_on_table"`

	IsDeleted            bool       `json:"isDeleted"            column:"is_deleted"`
	DeleteFromDatabaseAt *time.Time `json:"deleteFromDatabaseAt" column:"delete_from_database_at"`
	CreatedAt            time.Time  `json:"createdAt"            column:"created_at"`
	UpdatedAt            time.Time  `json:"updatedAt"            column:"updated_at"`
}
