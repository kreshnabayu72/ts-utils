package model

import "time"

type NfcStock struct {
	ID        int        `json:"_id"       column:"nfc_stocks.id"`
	ShortCode string     `json:"shortCode" column:"nfc_stocks.short_code"`
	Code      string     `json:"code"      column:"nfc_stocks.code"`
	ClientID  *string    `json:"clientId"  column:"nfc_stocks.client_id"`
	UsageType string     `json:"usageType" column:"nfc_stocks.usage_type"`
	IsDeleted bool       `json:"isDeleted" column:"nfc_stocks.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"nfc_stocks.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"nfc_stocks.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"nfc_stocks.updated_at"`
}

func (r *NfcStock) GetColumns() []string {
	return []string{
		"id",
		"short_code",
		"code",
		"client_id",
		"usage_type",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}
