package new_model

import "time"

type ImportGroup struct {
	ID           int     `json:"_id"                  column:"id"`
	Name         string  `json:"name"                column:"name"`
	GroupCode    string  `json:"groupCode"          column:"group_code"`
	IsDisposable *string `json:"isDisposable"       column:"is_disposable"`
	Key          int     `json:"key"                 column:"key"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
