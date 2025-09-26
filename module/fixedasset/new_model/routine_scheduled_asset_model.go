package new_model

import (
	"time"
)

type RoutineScheduledAsset struct {
	ID         int       `json:"_id"        column:"routine_scheduled_assets.id"`
	ScheduleID int       `json:"scheduleId" column:"routine_scheduled_assets.schedule_id"`
	AssetID    int       `json:"assetId"    column:"routine_scheduled_assets.asset_id"`
	IsActive   bool      `json:"isActive"   column:"routine_scheduled_assets.is_active"`
	CreatedAt  time.Time `json:"createdAt"  column:"routine_scheduled_assets.created_at"`
	UpdatedAt  time.Time `json:"updatedAt"  column:"routine_scheduled_assets.updated_at"`
}
