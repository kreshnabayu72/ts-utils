package new_model

import (
	"time"
)

type RoutineTask struct {
	ID            int               `json:"_id"           column:"routine_tasks.id"`
	ScheduleID    int               `json:"scheduleId"    column:"routine_tasks.schedule_id"`
	Name          string            `json:"name" column:"routine_tasks.name"`
	AssetID       int               `json:"assetId"       column:"routine_tasks.asset_id"`
	Status        RoutineTaskStatus `json:"status"        column:"routine_tasks.status"`
	Duration      int               `json:"duration"      column:"routine_tasks.duration"`
	ModifiedByID  int               `json:"modifiedById"  column:"routine_tasks.modified_by_id"`
	ApprovalRound int               `json:"approvalRound" column:"routine_tasks.approval_round"`
	Address       string            `json:"address"       column:"routine_tasks.address"`
	Coordinate    string            `json:"coordinate"    column:"routine_tasks.coordinate"`
	StartAt       time.Time         `json:"startAt"       column:"routine_tasks.start_at"`
	CreatedAt     time.Time         `json:"createdAt"     column:"routine_tasks.created_at"`
	UpdatedAt     time.Time         `json:"updatedAt"     column:"routine_tasks.updated_at"`
}
