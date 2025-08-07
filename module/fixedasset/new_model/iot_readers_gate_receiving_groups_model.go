package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type IotReaderGateReceivingGroup struct {
	ID          int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	IotReaderID int `json:"iotReaderId" column:"iot_reader_id" sql:"BIGINT NOT NULL"`
	GroupID     int `json:"groupId" column:"group_id" sql:"BIGINT NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (g *IotReaderGateReceivingGroup) TableName() string {
	return "iot_reader_gate_receiving_groups"
}

func (g *IotReaderGateReceivingGroup) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(g)

	builder.AddForeignKey(
		"fk_iot_reader_gate_reader_id",
		"iot_reader_id",
		"iot_readers",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_iot_reader_gate_group_id",
		"group_id",
		"groups",
		"id",
		"",
	)
}
