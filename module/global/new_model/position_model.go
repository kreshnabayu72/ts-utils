package new_model

import "time"

type Position struct {
	ID        int       `json:"_id"                  column:"id"`
	Name      string    `json:"name"                column:"name"`
	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
