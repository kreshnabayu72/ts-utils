package new_model

import "time"

type ImportDepreciationGroup struct {
	ID         int    `json:"_id"                  column:"id"`
	Key        int    `json:"key"                 column:"key"`
	Name       string `json:"name"                column:"name"`
	UsefulLife int    `json:"usefulLife"                column:"useful_life"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
