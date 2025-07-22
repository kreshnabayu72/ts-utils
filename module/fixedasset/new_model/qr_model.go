package new_model

import "time"

type Qr struct {
	ID int `json:"_id"         column:"id"`

	Code   string    `json:"code"       column:"code"`
	Status TagStatus `json:"status"       column:"status"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
