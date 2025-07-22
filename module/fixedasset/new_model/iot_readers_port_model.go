package new_model

type IotReaderPort struct {
	ID          int  `json:"_id" column:"id"`
	IotReaderID int  `json:"iotReaderId" column:"iot_reader_id"`
	Power       int  `json:"power" column:"power"`
	GroupID     *int `json:"groupId" column:"group_id"`
	IsActive    bool `json:"isActive" column:"is_active"`
}
