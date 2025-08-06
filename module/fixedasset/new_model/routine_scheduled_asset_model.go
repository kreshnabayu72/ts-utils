package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineScheduledAsset struct {
	ID         int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	ScheduleID int       `json:"scheduleId" column:"schedule_id" sql:"BIGINT NOT NULL"`
	AssetID    int       `json:"assetId" column:"asset_id" sql:"BIGINT NOT NULL"`
	IsActive   bool      `json:"isActive" column:"is_active" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	CreatedAt  time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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
