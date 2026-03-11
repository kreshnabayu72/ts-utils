package replica

import "time"

type SmartFixedStock struct {
	ID                 int        `json:"_id" column:"smart_fixed_stocks.id"`
	SkuID              int        `json:"skuId" column:"smart_fixed_stocks.sku_id"`
	SerialNumber       string     `json:"serialNumber" column:"smart_fixed_stocks.serial_number"`
	Type               string     `json:"type" column:"smart_fixed_stocks.type"`
	DeviceNameID       int        `json:"deviceNameId" column:"smart_fixed_stocks.device_name_id"`
	DeviceNameSequence int        `json:"deviceNameSequence" column:"smart_fixed_stocks.device_name_sequence"`
	IotCode            string     `json:"iotCode" column:"smart_fixed_stocks.iot_code"`
	LicenseKey         string     `json:"licenseKey" column:"smart_fixed_stocks.license_key"`
	ClientID           *string    `json:"clientId" column:"smart_fixed_stocks.client_id"`
	IsDeleted          bool       `json:"isDeleted" column:"smart_fixed_stocks.is_deleted"`
	DeletedAt          *time.Time `json:"deletedAt" column:"smart_fixed_stocks.deleted_at"`
	CreatedAt          time.Time  `json:"createdAt" column:"smart_fixed_stocks.created_at"`
	UpdatedAt          time.Time  `json:"updatedAt" column:"smart_fixed_stocks.updated_at"`
}

func (r *SmartFixedStock) GetTableName() string {
	return "smart_fixed_stocks"
}

func (r *SmartFixedStock) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                   "BIGINT",
		"sku_id":               "BIGINT",
		"serial_number":        "VARCHAR",
		"type":                 "VARCHAR",
		"device_name_id":       "BIGINT",
		"device_name_sequence": "INT",
		"iot_code":             "VARCHAR",
		"license_key":          "VARCHAR",
		"client_id":            "VARCHAR",
		"is_deleted":           "BOOL",
		"deleted_at":           "TIMESTAMPTZ",
		"created_at":           "TIMESTAMPTZ",
		"updated_at":           "TIMESTAMPTZ",
	}
}

func (r *SmartFixedStock) GetColumns() []string {
	return []string{
		"id",
		"sku_id",
		"serial_number",
		"type",
		"device_name_id",
		"device_name_sequence",
		"iot_code",
		"license_key",
		"client_id",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}
