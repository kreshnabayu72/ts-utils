package replica

import "time"

type HandheldStock struct {
	ID                 int        `json:"_id" column:"handheld_stocks.id" sqlType:"BIGINT"`
	SkuID              int        `json:"skuId" column:"handheld_stocks.sku_id" sqlType:"BIGINT"`
	SerialNumber       string     `json:"serialNumber" column:"handheld_stocks.serial_number" sqlType:"VARCHAR"`
	Type               string     `json:"type" column:"handheld_stocks.type" sqlType:"VARCHAR"`
	DeviceNameID       int        `json:"deviceNameId" column:"handheld_stocks.device_name_id" sqlType:"BIGINT"`
	DeviceNameSequence int        `json:"deviceNameSequence" column:"handheld_stocks.device_name_sequence" sqlType:"INT"`
	HardwareID         *string    `json:"hardwareId" column:"handheld_stocks.hardware_id" sqlType:"VARCHAR"`
	LicenseKey         string     `json:"licenseKey" column:"handheld_stocks.license_key" sqlType:"VARCHAR"`
	ClientID           *string    `json:"clientId" column:"handheld_stocks.client_id" sqlType:"VARCHAR"`
	DeletedAt          *time.Time `json:"deletedAt" column:"handheld_stocks.deleted_at" sqlType:"TIMESTAMPTZ"`
	IsDeleted          bool       `json:"isDeleted" column:"handheld_stocks.is_deleted" sqlType:"BOOL"`
	CreatedAt          time.Time  `json:"createdAt" column:"handheld_stocks.created_at" sqlType:"TIMESTAMPTZ"`
	UpdatedAt          time.Time  `json:"updatedAt" column:"handheld_stocks.updated_at" sqlType:"TIMESTAMPTZ"`
}

func (r *HandheldStock) GetTableName() string {
	return "handheld_stocks"
}

func (r *HandheldStock) GetTableNameWithSchema() string {
	return "public.handheld_stocks"
}

func (r *HandheldStock) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                   "BIGINT",
		"sku_id":               "BIGINT",
		"serial_number":        "VARCHAR",
		"type":                 "VARCHAR",
		"device_name_id":       "BIGINT",
		"device_name_sequence": "INT",
		"hardware_id":          "VARCHAR",
		"license_key":          "VARCHAR",
		"client_id":            "VARCHAR",
		"deleted_at":           "TIMESTAMPTZ",
		"is_deleted":           "BOOL",
		"created_at":           "TIMESTAMPTZ",
		"updated_at":           "TIMESTAMPTZ",
	}
}

func (r *HandheldStock) GetColumns() []string {
	return []string{
		"id",
		"sku_id",
		"serial_number",
		"type",
		"device_name_id",
		"device_name_sequence",
		"hardware_id",
		"license_key",
		"client_id",
		"deleted_at",
		"is_deleted",
		"created_at",
		"updated_at",
	}
}
