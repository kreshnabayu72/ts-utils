package helper

import (
	"context"
	"errors"

	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

func GetLoggedInUserSystemRolesHelper(service service.PostgreSqlService, ctx context.Context, companyCode string, loggedInUserId int, groupId ...int) (*[]new_model.UserSystemRole, error) {
	userSystemRoles := new([]new_model.UserSystemRole)
	if err := sql_builder.FindUserSystemRoles(userSystemRoles, ctx, companyCode, loggedInUserId, groupId...); err != nil {
		return nil, errors.New("failed to get user system roles")
	}

	return userSystemRoles, nil
}
