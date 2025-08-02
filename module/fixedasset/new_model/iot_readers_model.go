package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type IotReader struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	GroupID   *int `json:"groupId" column:"group_id" sql:"BIGINT"`
	ManagerID int  `json:"managerId" column:"manager_id" sql:"BIGINT"`
	QrCode    *int `json:"qrCode" column:"qr_code"` // NOTE: Go struct uses INT, SQL uses VARCHAR

	Image            string `json:"image" column:"image" sql:"VARCHAR"`
	SerialNumber     string `json:"serialNumber" column:"serial_number" sql:"VARCHAR"`
	NumberOfPorts    int    `json:"numberOfPorts" column:"number_of_ports" sql:"INTEGER"`
	InputTopic       string `json:"inputTopic" column:"input_topic" sql:"VARCHAR"`
	OutputTopic      string `json:"outputTopic" column:"output_topic" sql:"VARCHAR"`
	UniqueIdentifier string `json:"uniqueIdentifier" column:"unique_identifier" sql:"VARCHAR"`

	DeviceNameID       *int   `json:"deviceNameId" column:"device_name_id" sql:"BIGINT"`
	DeviceNameName     string `json:"deviceNameName" column:"device_name_name" sql:"VARCHAR"`
	DeviceNameSequence int    `json:"deviceNameSequence" column:"device_name_sequence" sql:"INTEGER"`

	BrandID   *int   `json:"brandId" column:"brand_id" sql:"BIGINT"`
	BrandName string `json:"brandName" column:"brand_name" sql:"VARCHAR"`

	ModelID   *int   `json:"modelId" column:"model_id" sql:"BIGINT"`
	ModelName string `json:"modelName" column:"model_name" sql:"VARCHAR"`

	Status           ReaderStatus `json:"status" column:"status" sql:"reader_status_enum DEFAULT 'Available'"`
	LastReportReason string       `json:"lastReportReason" column:"last_report_reason" sql:"reader_status_enum"`
	LastReportAt     *time.Time   `json:"lastReportAt" column:"last_report_at" sql:"TIMESTAMP DEFAULT NULL"`

	NetworkStatus string `json:"networkStatus" column:"network_status" sql:"network_status_enum DEFAULT 'Offline'"`
	ResponseID    string `json:"responseId" column:"response_id" sql:"VARCHAR"`
	AliasName     string `json:"aliasName" column:"alias_name" sql:"VARCHAR"`

	AssetControlActive         bool `json:"assetControlActive" column:"asset_control_active" sql:"BOOLEAN"`
	AssetControlRoleRestricted bool `json:"assetControlRoleRestricted" column:"asset_control_role_restricted" sql:"BOOLEAN"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (i *IotReader) TableName() string {
	return "iot_readers"
}

func (i *IotReader) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(i)

	builder.AddForeignKey(
		"fk_iot_readers_group_id",
		"group_id",
		"groups",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_iot_readers_manager_id",
		"manager_id",
		"users",
		"id",
		"",
	)

	builder.AddIndex(
		"iot_readers_qr_code_uindex",
		"qr_code",
		true,
		"qr_code IS NOT NULL",
	)

	builder.AddIndex(
		"iot_readers_serial_number_uindex",
		"serial_number",
		true,
		"",
	)
}
