package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type IotReader struct {
	ID                         int          `json:"_id"                        column:"iot_readers.id"                            sql:"BIGINT PRIMARY KEY"`
	GroupID                    *int         `json:"groupId"                    column:"iot_readers.group_id"                      sql:"BIGINT"`
	ManagerID                  int          `json:"managerId"                  column:"iot_readers.manager_id"                    sql:"BIGINT"`
	QrCode                     *string      `json:"qrCode"                     column:"iot_readers.qr_code"                       sql:"VARCHAR NULL"`
	Image                      string       `json:"image"                      column:"iot_readers.image"                         sql:"VARCHAR"`
	SerialNumber               string       `json:"serialNumber"               column:"iot_readers.serial_number"                 sql:"VARCHAR"`
	NumberOfPorts              int          `json:"numberOfPorts"              column:"iot_readers.number_of_ports"               sql:"INTEGER"`
	InputTopic                 string       `json:"inputTopic"                 column:"iot_readers.input_topic"                   sql:"VARCHAR"`
	OutputTopic                string       `json:"outputTopic"                column:"iot_readers.output_topic"                  sql:"VARCHAR"`
	UniqueIdentifier           string       `json:"uniqueIdentifier"           column:"iot_readers.unique_identifier"             sql:"VARCHAR"`
	DeviceNameID               *int         `json:"deviceNameId"               column:"iot_readers.device_name_id"                sql:"BIGINT"`
	DeviceNameName             string       `json:"deviceNameName"             column:"iot_readers.device_name_name"              sql:"VARCHAR"`
	DeviceNameSequence         int          `json:"deviceNameSequence"         column:"iot_readers.device_name_sequence"          sql:"INTEGER"`
	BrandID                    *int         `json:"brandId"                    column:"iot_readers.brand_id"                      sql:"BIGINT"`
	BrandName                  string       `json:"brandName"                  column:"iot_readers.brand_name"                    sql:"VARCHAR"`
	ModelID                    *int         `json:"modelId"                    column:"iot_readers.model_id"                      sql:"BIGINT"`
	ModelName                  string       `json:"modelName"                  column:"iot_readers.model_name"                    sql:"VARCHAR"`
	Status                     ReaderStatus `json:"status"                     column:"iot_readers.status"                        sql:"reader_status_enum DEFAULT 'Available'"`
	LastReportReason           string       `json:"lastReportReason"           column:"iot_readers.last_report_reason"            sql:"reader_status_enum"`
	LastReportAt               *time.Time   `json:"lastReportAt"               column:"iot_readers.last_report_at"                sql:"TIMESTAMPTZ DEFAULT NULL"`
	NetworkStatus              string       `json:"networkStatus"              column:"iot_readers.network_status"                sql:"network_status_enum DEFAULT 'Offline'"`
	ResponseID                 string       `json:"responseId"                 column:"iot_readers.response_id"                   sql:"VARCHAR"`
	AliasName                  string       `json:"aliasName"                  column:"iot_readers.alias_name"                    sql:"VARCHAR"`
	AssetControlActive         bool         `json:"assetControlActive"         column:"iot_readers.asset_control_active"          sql:"BOOLEAN"`
	AssetControlRoleRestricted bool         `json:"assetControlRoleRestricted" column:"iot_readers.asset_control_role_restricted" sql:"BOOLEAN"`
	CreatedAt                  time.Time    `json:"createdAt"                  column:"iot_readers.created_at"                    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                  time.Time    `json:"updatedAt"                  column:"iot_readers.updated_at"                    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
