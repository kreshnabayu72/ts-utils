package new_model

import "time"

type AuditIotScanAssetLog struct {
	ID           int       `json:"_id"          bson:"_id"`
	IotReaderID  string    `json:"iotReaderId"  bson:"iotReaderId"`
	AuditTaskID  string    `json:"auditTaskId"  bson:"auditTaskId"`
	SerialNumber string    `json:"serialNumber" bson:"serialNumber"`
	AssetId      string    `json:"assetId"      bson:"assetId"`
	CreatedAt    time.Time `json:"createdAt"    bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"    bson:"updatedAt"`
}
