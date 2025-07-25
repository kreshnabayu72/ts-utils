package new_model

import (
	"time"
)

type AccountCode struct {
	ID                   int        `json:"_id"                  column:"id"`
	Name                 string     `json:"name"                 column:"name"`
	Code                 string     `json:"code"                 column:"code"`
	IsDeleted            bool       `json:"isDeleted"            column:"is_deleted"`
	DeleteFromDatabaseAt *time.Time `json:"deleteFromDatabaseAt" column:"delete_from_database_at"`
	CreatedAt            time.Time  `json:"createdAt"            column:"created_at"`
	UpdatedAt            time.Time  `json:"updatedAt"            column:"updated_at"`
}
