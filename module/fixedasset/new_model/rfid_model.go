package new_model

import "time"

type Rfid struct {
	ID int `json:"_id"         column:"id"`

	Code            string      `json:"code"       column:"code"`
	Status          TagStatus   `json:"status"       column:"status"`
	LastAssetStatus AssetStatus `json:"lastAssetStatus"       column:"last_asset_status"`

	HolderId *int `json:"holderId"               column:"holder_id"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
