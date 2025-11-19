package model

import (
	"time"
)

// AssignmentPrelist represents the assignment_prelists table in the database.
type AssignmentPrelist struct {
	ID           int                   `json:"id"           column:"assignment_prelists.id"`
	GroupID      int                   `json:"groupId"      column:"assignment_prelists.group_id"`
	AssignedToID int                   `json:"assignedToId" column:"assignment_prelists.assigned_to_id"`
	Type         AssignmentPrelistType `json:"type"         column:"assignment_prelists.type"`
	CreatedAt    time.Time             `json:"createdAt"    column:"assignment_prelists.created_at"`
	UpdatedAt    time.Time             `json:"updatedAt"    column:"assignment_prelists.updated_at"`
}
