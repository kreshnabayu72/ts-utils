package new_model

import "time"

type AssignmentPrelist struct {
	ID           int                   `json:"id" column:"id"`
	GroupId      int                   `json:"groupId" column:"group_id"`
	AssignedToId int                   `json:"assignedToId" column:"assigned_to_id"`
	Type         AssignmentPrelistType `json:"type" column:"type"`
	CreatedAt    time.Time             `json:"createdAt" column:"created_at"`
	UpdatedAt    time.Time             `json:"updatedAt" column:"updated_at"`
}
