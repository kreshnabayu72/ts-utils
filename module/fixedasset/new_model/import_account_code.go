package new_model

import (
	"time"
)

type ImportAccountCode struct {
	ID        int       `json:"_id"       column:"import_account_codes.id"`
	Key       int       `json:"key"       column:"import_account_codes.key"        special:"generated"`
	Name      string    `json:"name"      column:"import_account_codes.name"`
	Code      string    `json:"code"      column:"import_account_codes.code"`
	CreatedAt time.Time `json:"createdAt" column:"import_account_codes.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"import_account_codes.updated_at"`
}
