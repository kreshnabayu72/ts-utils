package new_model

import (
	"time"
)

type UserTransactionRole struct {
	ID                int  `json:"_id" column:"id"`
	IsActive          bool `json:"isActive" column:"is_active"`
	UserID            int  `json:"userId" column:"user_id"`
	IsManagedByParent bool `json:"isManagedByParent" column:"is_managed_by_parent"`

	RoleAttribute TransactionRoleAttribute `json:"roleAttribute" column:"role_attribute"`
	RoleName      TransactionRoleName      `json:"roleName" column:"role_name"`
	RoleType      TransactionRoleType      `json:"roleType" column:"role_type"`

	GroupID       int    `json:"groupId" column:"group_id"`
	ApprovalLevel *int   `json:"approvalLevel" column:"approval_level"`
	ApprovalType  string `json:"approvalType" column:"approval_type"` // "Or" or "And"

	CreatedAt *time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at"`
}
