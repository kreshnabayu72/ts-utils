package helper

import (
	"context"
	"errors"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
)

// Cant use new_model.UserTransactionRole since the builder has join & need table alias (utr.id, etc)
type UserTransactionRolesData struct {
	ID       int  `json:"_id" column:"utr.id"`
	IsActive bool `json:"isActive" column:"utr.is_active"`
	UserID   int  `json:"userId" column:"utr.user_id"`

	RoleAttribute new_model.TransactionRoleAttribute `json:"roleAttribute" column:"utr.role_attribute"`
	RoleName      new_model.TransactionRoleName      `json:"roleName" column:"utr.role_name"`
	RoleType      new_model.TransactionRoleType      `json:"roleType" column:"utr.role_type"`

	GroupID       int    `json:"groupId" column:"utr.group_id"`
	ApprovalLevel *int   `json:"approvalLevel" column:"utr.approval_level"`
	ApprovalType  string `json:"approvalType" column:"utr.approval_type"`

	CreatedAt *time.Time `json:"createdAt" column:"utr.created_at"`
	UpdatedAt *time.Time `json:"updatedAt" column:"utr.updated_at"`
}

func GetLoggedInUserTransactionRolesHelper(service service.PostgreSqlService, ctx context.Context, loggedInUserId int, roleFilter ...sql_builder.TransactionRoleFilter) (*[]UserTransactionRolesData, error) {
	userTransactionRoles := new([]UserTransactionRolesData)
	if err := sql_builder.FindUserTransactionRoles(userTransactionRoles, service, ctx, loggedInUserId, roleFilter...); err != nil {
		return nil, errors.New("failed to get user transaction roles")
	}

	return userTransactionRoles, nil
}
