package model

import "time"

type RfidStock struct {
	ID        int        `json:"_id"       column:"rfid_stocks.id"`
	ShortCode string     `json:"shortCode" column:"rfid_stocks.short_code"`
	Code      string     `json:"code"      column:"rfid_stocks.code"`
	ClientID  *string    `json:"clientId"  column:"rfid_stocks.client_id"`
	UsageType string     `json:"usageType" column:"rfid_stocks.usage_type"`
	IsDeleted bool       `json:"isDeleted" column:"rfid_stocks.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"rfid_stocks.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"rfid_stocks.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"rfid_stocks.updated_at"`
}

