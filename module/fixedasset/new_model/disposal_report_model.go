package new_model

import "time"

type DisposalReport struct {
	ID            int                   `json:"id" column:"id"`
	AssetId       int                   `json:"assetId" column:"asset_id"`
	ReportedById  int                   `json:"reportedById" column:"reported_by_id"`
	LastCondition DisposalLastCondition `json:"lastCondition" column:"last_condition"`
	CreatedAt     time.Time             `json:"createdAt" column:"created_at"`
	UpdatedAt     time.Time             `json:"updatedAt" column:"updated_at"`
}
