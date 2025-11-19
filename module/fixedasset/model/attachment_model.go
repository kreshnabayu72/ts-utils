package model

import (
	"time"
)

type Attachment struct {
	ID        int            `json:"_id"       column:"attachments.id"`
	AssetID   int            `json:"assetId"   column:"attachments.asset_id"`
	Type      AttachmentType `json:"type"      column:"attachments.type"`
	Name      string         `json:"name"      column:"attachments.name"`
	File      string         `json:"file"      column:"attachments.file"`
	FileName  string         `json:"fileName"      column:"attachments.file_name"`
	CreatedBy int            `json:"createdBy" column:"attachments.created_by"`
	CreatedAt time.Time      `json:"createdAt" column:"attachments.created_at"`
	UpdatedAt time.Time      `json:"updatedAt" column:"attachments.updated_at"`
}
