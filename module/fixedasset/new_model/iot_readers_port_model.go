package new_model

import (
	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type IotReaderPort struct {
	ID          int  `json:"_id"         column:"iot_reader_ports.id"            sql:"BIGINT PRIMARY KEY"`
	IotReaderID int  `json:"iotReaderId" column:"iot_reader_ports.iot_reader_id" sql:"BIGINT NOT NULL"`
	Power       int  `json:"power"       column:"iot_reader_ports.power"         sql:"INTEGER"`
	GroupID     *int `json:"groupId"     column:"iot_reader_ports.group_id"      sql:"BIGINT"`
	IsActive    bool `json:"isActive"    column:"iot_reader_ports.is_active"     sql:"BOOLEAN"`
}

func (p *IotReaderPort) TableName() string {
	return "iot_reader_ports"
}

func (p *IotReaderPort) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(p)

	builder.AddForeignKey(
		"fk_iot_reader_ports_reader_id",
		"iot_reader_id",
		"iot_readers",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_iot_reader_ports_group_id",
		"group_id",
		"groups",
		"id",
		"",
	)
}
