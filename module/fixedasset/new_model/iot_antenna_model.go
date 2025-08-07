package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type IotAntenna struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	IotReaderPortID *int   `json:"iotReaderPortId" column:"iot_reader_port_id" sql:"BIGINT"`
	Image           string `json:"image" column:"image" sql:"VARCHAR"`

	DeviceNameID       *int   `json:"deviceNameId" column:"device_name_id" sql:"BIGINT"`
	DeviceNameName     string `json:"deviceNameName" column:"device_name_name" sql:"VARCHAR"`
	DeviceNameSequence int    `json:"deviceNameSequence" column:"device_name_sequence" sql:"INTEGER"`

	BrandID   *int   `json:"brandId" column:"brand_id" sql:"BIGINT"`
	BrandName string `json:"brandName" column:"brand_name" sql:"VARCHAR"`

	ModelID   *int   `json:"modelId" column:"model_id" sql:"BIGINT"`
	ModelName string `json:"modelName" column:"model_name" sql:"VARCHAR"`

	ManagerID *int `json:"managerId" column:"manager_id" sql:"BIGINT"`

	Status     ReaderStatus `json:"status" column:"status" sql:"reader_status_enum"`
	ReportedAt *time.Time   `json:"reportedAt" column:"reported_at" sql:"TIMESTAMPTZ DEFAULT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
