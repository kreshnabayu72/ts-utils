package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type IotAntenna struct {
	ID                 int          `json:"_id"                column:"iot_antenna.id"                   sql:"BIGINT PRIMARY KEY"`
	IotReaderPortID    *int         `json:"iotReaderPortId"    column:"iot_antenna.iot_reader_port_id"   sql:"BIGINT"`
	Image              string       `json:"image"              column:"iot_antenna.image"                sql:"VARCHAR"`
	DeviceNameID       *int         `json:"deviceNameId"       column:"iot_antenna.device_name_id"       sql:"BIGINT"`
	DeviceNameName     string       `json:"deviceNameName"     column:"iot_antenna.device_name_name"     sql:"VARCHAR"`
	DeviceNameSequence int          `json:"deviceNameSequence" column:"iot_antenna.device_name_sequence" sql:"INTEGER"`
	BrandID            *int         `json:"brandId"            column:"iot_antenna.brand_id"             sql:"BIGINT"`
	BrandName          string       `json:"brandName"          column:"iot_antenna.brand_name"           sql:"VARCHAR"`
	ModelID            *int         `json:"modelId"            column:"iot_antenna.model_id"             sql:"BIGINT"`
	ModelName          string       `json:"modelName"          column:"iot_antenna.model_name"           sql:"VARCHAR"`
	ManagerID          *int         `json:"managerId"          column:"iot_antenna.manager_id"           sql:"BIGINT"`
	Status             ReaderStatus `json:"status"             column:"iot_antenna.status"               sql:"reader_status_enum"`
	ReportedAt         *time.Time   `json:"reportedAt"         column:"iot_antenna.reported_at"          sql:"TIMESTAMPTZ DEFAULT NULL"`
	CreatedAt          time.Time    `json:"createdAt"          column:"iot_antenna.created_at"           sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time    `json:"updatedAt"          column:"iot_antenna.updated_at"           sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (a *IotAntenna) TableName() string {
	return "iot_antenna"
}

func (a *IotAntenna) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(a)

	builder.AddForeignKey(
		"fk_antenna_reader_port",
		"iot_reader_port_id",
		"iot_reader_ports",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_antenna_manager",
		"manager_id",
		"users",
		"id",
		"",
	)
}
