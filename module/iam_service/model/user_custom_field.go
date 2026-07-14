package model

import "time"

type UserCustomField struct {
	ID            int        `json:"_id"           column:"user_custom_fields.id"`
	UserID        int        `json:"userId"        column:"user_custom_fields.user_id"`
	CustomFieldID int        `json:"customFieldId" column:"user_custom_fields.custom_field_id"`
	Value         *any       `json:"value"         column:"user_custom_fields.value"`
	IsDeleted     bool       `json:"isDeleted"     column:"user_custom_fields.is_deleted"`
	DeletedAt     *time.Time `json:"deletedAt"     column:"user_custom_fields.deleted_at"`
	CreatedAt     time.Time  `json:"createdAt"     column:"user_custom_fields.created_at"`
	UpdatedAt     time.Time  `json:"updatedAt"     column:"user_custom_fields.updated_at"`
}

