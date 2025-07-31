package helper

import (
	"context"
	"errors"

	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

func GetLoggedInUserTransactionRolesHelper(service service.PostgreSqlService, ctx context.Context, loggedInUserId int, roleFilter ...sql_builder.TransactionRoleFilter) (*[]new_model.UserTransactionRole, error) {
	userTransactionRoles := new([]new_model.UserTransactionRole)
	if err := sql_builder.FindUserTransactionRoles(userTransactionRoles, service, ctx, loggedInUserId, roleFilter...); err != nil {
		return nil, errors.New("failed to get user transaction roles")
	}

	return userTransactionRoles, nil
}
