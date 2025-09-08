package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineScheduledAsset struct {
	ID         int       `json:"_id"        column:"routine_scheduled_assets.id"          sql:"BIGINT PRIMARY KEY"`
	ScheduleID int       `json:"scheduleId" column:"routine_scheduled_assets.schedule_id" sql:"BIGINT NOT NULL"`
	AssetID    int       `json:"assetId"    column:"routine_scheduled_assets.asset_id"    sql:"BIGINT NOT NULL"`
	IsActive   bool      `json:"isActive"   column:"routine_scheduled_assets.is_active"   sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	CreatedAt  time.Time `json:"createdAt"  column:"routine_scheduled_assets.created_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updatedAt"  column:"routine_scheduled_assets.updated_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RoutineScheduledAsset) TableName() string {
	return "routine_scheduled_assets"
}

func (m *RoutineScheduledAsset) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_routine_schedule_asset_schedule_id",
		"schedule_id",
		"routine_schedules",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_schedule_asset_asset_id",
		"asset_id",
		"assets",
		"id",
		"",
	)
}
