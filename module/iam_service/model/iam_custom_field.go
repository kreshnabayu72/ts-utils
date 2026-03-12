package model

import "time"

type IAMCustomField struct {
	ID            int        `json:"_id"           column:"iam_custom_fields.id"`
	Name          string     `json:"name"          column:"iam_custom_fields.name"`
	DataType      string     `json:"dataType"      column:"iam_custom_fields.data_type"`
	OptionValue   []string   `json:"optionValue"   column:"iam_custom_fields.option_value"`
	DecimalPlaces *int       `json:"decimalPlaces" column:"iam_custom_fields.decimal_places"` // ✅ Pointer = nullable
	IsRequired    bool       `json:"isRequired"    column:"iam_custom_fields.is_required"`
	IsActive      bool       `json:"isActive"      column:"iam_custom_fields.is_active"`
	IsDeleted     bool       `json:"isDeleted"     column:"iam_custom_fields.is_deleted"`
	DeletedAt     *time.Time `json:"deletedAt"     column:"iam_custom_fields.deleted_at"`
	CreatedAt     time.Time  `json:"createdAt"     column:"iam_custom_fields.created_at"`
	UpdatedAt     time.Time  `json:"updatedAt"     column:"iam_custom_fields.updated_at"`
}

func (r *IAMCustomField) GetColumns() []string {
	return []string{
		"id",
		"name",
		"data_type",
		"option_value",
		"decimal_places",
		"is_required",
		"is_active",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}
