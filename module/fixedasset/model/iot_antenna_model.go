package model

import (
	"time"
)

type IotAntenna struct {
	ID                 int          `json:"_id"                column:"iot_antenna.id"`
	IotReaderPortID    *int         `json:"iotReaderPortId"    column:"iot_antenna.iot_reader_port_id"`
	Image              string       `json:"image"              column:"iot_antenna.image"`
	DeviceNameID       *int         `json:"deviceNameId"       column:"iot_antenna.device_name_id"`
	DeviceNameName     string       `json:"deviceNameName"     column:"iot_antenna.device_name_name"`
	DeviceNameSequence int          `json:"deviceNameSequence" column:"iot_antenna.device_name_sequence"`
	BrandID            *int         `json:"brandId"            column:"iot_antenna.brand_id"`
	BrandName          string       `json:"brandName"          column:"iot_antenna.brand_name"`
	ModelID            *int         `json:"modelId"            column:"iot_antenna.model_id"`
	ModelName          string       `json:"modelName"          column:"iot_antenna.model_name"`
	ManagerID          *int         `json:"managerId"          column:"iot_antenna.manager_id"`
	Status             ReaderStatus `json:"status"             column:"iot_antenna.status"`
	ReportedAt         *time.Time   `json:"reportedAt"         column:"iot_antenna.reported_at"`
	CreatedAt          time.Time    `json:"createdAt"          column:"iot_antenna.created_at"`
	UpdatedAt          time.Time    `json:"updatedAt"          column:"iot_antenna.updated_at"`
}
