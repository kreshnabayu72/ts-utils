package new_model

import "time"

type IotReaderGateReceivingGroup struct {
	ID          int `json:"_id" column:"id"`
	IotReaderID int `json:"iotReaderId" column:"iot_reader_id"`
	GroupID     int `json:"groupId" column:"group_id"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
