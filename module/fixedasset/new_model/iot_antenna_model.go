package new_model

import "time"

type IotAntenna struct {
	ID int `json:"_id" column:"id"`

	IotReaderPortID *int   `json:"iotReaderPortId" column:"iot_reader_port_id"`
	Image           string `json:"image" column:"image"`

	DeviceNameID       *int   `json:"deviceNameId" column:"device_name_id"`
	DeviceNameName     string `json:"deviceNameName" column:"device_name_name"`
	DeviceNameSequence int    `json:"deviceNameSequence" column:"device_name_sequence"`

	BrandID   *int   `json:"brandId" column:"brand_id"`
	BrandName string `json:"brandName" column:"brand_name"`

	ModelID   *int   `json:"modelId" column:"model_id"`
	ModelName string `json:"modelName" column:"model_name"`

	ManagerID *int `json:"managerId" column:"manager_id"`

	Status     ReaderStatus `json:"status" column:"status"`
	ReportedAt *time.Time   `json:"reportedAt" column:"reported_at"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
