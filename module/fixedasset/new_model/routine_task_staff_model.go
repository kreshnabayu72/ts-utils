package new_model

import (
	"time"
)

type RoutineTaskStaff struct {
	ID            int       `json:"_id"           column:"routine_task_staffs.id"`
	UserID        int       `json:"userId"        column:"routine_task_staffs.user_id"`
	RoutineTaskID int       `json:"routineTaskId" column:"routine_task_staffs.routine_task_id"`
	CreatedAt     time.Time `json:"createdAt"     column:"routine_task_staffs.created_at"`
	UpdatedAt     time.Time `json:"updatedAt"     column:"routine_task_staffs.updated_at"`
}
