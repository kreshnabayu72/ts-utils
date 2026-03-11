package replica

import "time"

type Rfid struct {
	ID        int       `json:"_id"       column:"rfids.id"`
	ShortCode string    `json:"shortCode" column:"rfids.short_code"`
	Code      string    `json:"code"      column:"rfids.code"`
	CreatedAt time.Time `json:"createdAt" column:"rfids.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"rfids.updated_at"`
}

func (r *Rfid) GetTableName() string {
	return "rfids"
}

func (r *Rfid) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"short_code": "VARCHAR",
		"code":       "VARCHAR",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
	}
}

func (r *Rfid) GetColumns() []string {
	return []string{
		"id",
		"short_code",
		"code",
		"created_at",
		"updated_at",
	}
}
