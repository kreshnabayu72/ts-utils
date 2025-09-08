// TODO: added table name to column tag.
package new_model

import (
	"time"
)

type PendingTagTypeGroup struct {
	ID   int    `json:"_id"  column:"id"`
	Name string `json:"name" column:"name"`
}

type PendingTagTypeAssetName struct {
	ID               int    `json:"_id"              column:"id"`
	Name             string `json:"name"             column:"name"`
	NameWithSequence string `json:"nameWithSequence" column:"name_with_sequence"`
}

type PendingTagChange struct {
	ID                     int                     `json:"_id,omitempty"          column:"id"`
	Type                   string                  `json:"type"                   column:"type"`
	TagType                string                  `json:"tagType"                column:"tag_type"`
	RFID                   string                  `json:"rfid"                   column:"rfid"`
	QR                     string                  `json:"qr"                     column:"qr"`
	AssetStatus            string                  `json:"assetStatus"            column:"asset_status"`
	AssetMaintenanceStatus string                  `json:"assetMaintenanceStatus" column:"asset_maintenance_status"`
	AssetAuditStatus       string                  `json:"assetAuditStatus"       column:"asset_audit_status"`
	Group                  PendingTagTypeGroup     `json:"group"                  column:"group"`
	AssetName              PendingTagTypeAssetName `json:"assetName"              column:"asset_name"`
	Asset                  int                     `json:"asset"                  column:"asset"`
	CreatedAt              time.Time               `json:"createdAt"              column:"created_at"`
	UpdatedAt              time.Time               `json:"updatedAt"              column:"updated_at"`
}
