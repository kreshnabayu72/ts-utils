package replica

import "time"

type Attachment struct {
	ID        int       `json:"_id"        column:"attachments.id"          sqlType:"BIGINT"`
	AssetID   int       `json:"assetId"    column:"attachments.asset_id"    sqlType:"BIGINT"`
	Type      string    `json:"type"       column:"attachments.type"        sqlType:"VARCHAR"`
	Name      string    `json:"name"       column:"attachments.name"        sqlType:"VARCHAR"`
	File      string    `json:"file"       column:"attachments.file"        sqlType:"VARCHAR"`
	CreatedBy int       `json:"createdBy"  column:"attachments.created_by"  sqlType:"BIGINT"`
	CreatedAt time.Time `json:"createdAt"  column:"attachments.created_at"  sqlType:"TIMESTAMPTZ"`
	UpdatedAt time.Time `json:"updatedAt"  column:"attachments.updated_at"  sqlType:"TIMESTAMPTZ"`
	FileName  string    `json:"fileName"   column:"attachments.file_name"   sqlType:"VARCHAR"`
}

func (a *Attachment) GetTableName() string {
	return "attachments"
}

func (a *Attachment) GetTableNameWithSchema() string {
	return "public.attachments"
}

func (a *Attachment) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"asset_id":   "BIGINT",
		"type":       "VARCHAR",
		"name":       "VARCHAR",
		"file":       "VARCHAR",
		"created_by": "BIGINT",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
		"file_name":  "VARCHAR",
	}
}

func (a *Attachment) GetColumns() []string {
	return []string{
		"id",
		"asset_id",
		"type",
		"name",
		"file",
		"created_by",
		"created_at",
		"updated_at",
		"file_name",
	}
}
