package new_model

type IotReaderPort struct {
	ID          int  `json:"_id"         column:"iot_reader_ports.id"`
	IotReaderID int  `json:"iotReaderId" column:"iot_reader_ports.iot_reader_id"`
	Power       int  `json:"power"       column:"iot_reader_ports.power"`
	GroupID     *int `json:"groupId"     column:"iot_reader_ports.group_id"`
	IsActive    bool `json:"isActive"    column:"iot_reader_ports.is_active"`
}
