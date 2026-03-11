package replica

type Rfid struct {
	ID int `json:"_id"       column:"rfids.id"`
}

func (r *Rfid) GetTableName() string {
	return "rfids"
}

func (r *Rfid) GetColumnTypes() map[string]string {
	return map[string]string{
		"id": "BIGINT",
	}
}

func (r *Rfid) GetColumns() []string {
	return []string{
		"id",
	}
}
