package new_model

import "time"

type Attachment struct {
	ID        int            `json:"id" column:"id"`
	AssetID   int            `json:"assetId" column:"asset_id"`
	Type      AttachmentType `json:"type" column:"type"`
	Name      string         `json:"name" column:"name"`
	File      string         `json:"file" column:"file"`
	CreatedBy int            `json:"createdBy" column:"created_by"`
	CreatedAt time.Time      `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" column:"updated_at"`
}
