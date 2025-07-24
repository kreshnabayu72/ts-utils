package helper

import (
	"context"
	"errors"

	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query/common_builders"
)

func GetLoggedInUserHelper(service service.PostgreSqlService, ctx context.Context, loggedInUserId string) (*new_model.UserAttribute, error) {
	user := new(new_model.UserAttribute)
	query, args, _ := common_builders.GetLoggedInUserBuilder(loggedInUserId)

	if err := service.SelectOne(user, ctx, query, args...); err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("logged in user is not found")
	}

	return user, nil
}
