package new_model

import (
	"time"
)

type IotReaderGateReceivingGroup struct {
	ID          int       `json:"_id"         column:"iot_reader_gate_receiving_groups.id"`
	IotReaderID int       `json:"iotReaderId" column:"iot_reader_gate_receiving_groups.iot_reader_id"`
	GroupID     int       `json:"groupId"     column:"iot_reader_gate_receiving_groups.group_id"`
	CreatedAt   time.Time `json:"createdAt"   column:"iot_reader_gate_receiving_groups.created_at"`
	UpdatedAt   time.Time `json:"updatedAt"   column:"iot_reader_gate_receiving_groups.updated_at"`
}
