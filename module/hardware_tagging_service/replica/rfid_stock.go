package replica

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

func (r *RfidStock) GetTableName() string {
	return "rfid_stocks"
}

func (r *RfidStock) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"short_code": "VARCHAR",
		"code":       "VARCHAR",
		"client_id":  "VARCHAR",
		"usage_type": "VARCHAR",
		"is_deleted": "BOOLEAN",
		"deleted_at": "TIMESTAMPTZ",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
	}
}

func (r *RfidStock) GetColumns() []string {
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
