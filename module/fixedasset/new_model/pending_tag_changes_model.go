// TODO: added table name to column tag.
package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type PendingTagChange struct {
	ID        int       `json:"_id,omitempty"          column:"id" sql:"BIGINT PRIMARY KEY"`
	Type      string    `json:"type"                   column:"type" sql:"VARCHAR"`
	Asset     int       `json:"asset"                  column:"asset" sql:"BIGINT"`
	CreatedAt time.Time `json:"createdAt"              column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt"              column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *PendingTagChange) TableName() string {
	return "pending_tag_changes"
}

func (m *PendingTagChange) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_pending_tag_changes_asset",
		"asset",
		"assets",
		"id",
		"CASCADE",
	)
}
