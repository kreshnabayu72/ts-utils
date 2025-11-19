package helper

import (
	"context"
	"errors"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
)

func GetLoggedInUserSystemRolesHelper(service service.PostgreSqlService, ctx context.Context, loggedInUserId int, roleFilter ...sql_builder.SystemRoleFilter) (*[]new_model.UserSystemRole, error) {
	userSystemRoles := new([]new_model.UserSystemRole)
	if err := sql_builder.FindUserSystemRoles(userSystemRoles, service, ctx, loggedInUserId, roleFilter...); err != nil {
		return nil, errors.New("failed to get user system roles")
	}

	return userSystemRoles, nil
}
