package new_model

import "time"

type IotReader struct {
	ID int `json:"_id" column:"id"`

	GroupID   *int `json:"groupId" column:"group_id"`
	ManagerID int  `json:"managerId" column:"manager_id"`
	QrCode    *int `json:"qrCode" column:"qr_code"`

	Image            string `json:"image" column:"image"`
	SerialNumber     string `json:"serialNumber" column:"serial_number"`
	NumberOfPorts    int    `json:"numberOfPorts" column:"number_of_ports"`
	InputTopic       string `json:"inputTopic" column:"input_topic"`
	OutputTopic      string `json:"outputTopic" column:"output_topic"`
	UniqueIdentifier string `json:"uniqueIdentifier" column:"unique_identifier"`

	DeviceNameID       *int   `json:"deviceNameId" column:"device_name_id"`
	DeviceNameName     string `json:"deviceNameName" column:"device_name_name"`
	DeviceNameSequence int    `json:"deviceNameSequence" column:"device_name_sequence"`

	BrandID   *int   `json:"brandId" column:"brand_id"`
	BrandName string `json:"brandName" column:"brand_name"`

	ModelID   *int   `json:"modelId" column:"model_id"`
	ModelName string `json:"modelName" column:"model_name"`

	Status           ReaderStatus `json:"status" column:"status"`
	LastReportReason string       `json:"lastReportReason" column:"last_report_reason"`
	LastReportAt     *time.Time   `json:"lastReportAt" column:"last_report_at"`

	NetworkStatus string `json:"networkStatus" column:"network_status"` // "Online" or "Offline"
	ResponseID    string `json:"responseId" column:"response_id"`
	AliasName     string `json:"aliasName" column:"alias_name"`

	AssetControlActive         bool `json:"assetControlActive" column:"asset_control_active"`
	AssetControlRoleRestricted bool `json:"assetControlRoleRestricted" column:"asset_control_role_restricted"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
