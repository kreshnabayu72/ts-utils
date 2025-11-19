package helper

import (
	"context"
	"errors"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
)

func GetLoggedInUserHelper(service service.PostgreSqlService, ctx context.Context, loggedInUserId string) (*new_model.UserAttribute, error) {
	user := new(new_model.UserAttribute)
	query, args, _ := sql_builder.GetLoggedInUserBuilder(loggedInUserId)

	if err := service.SelectOne(user, ctx, query, args...); err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("logged in user is not found")
	}

	return user, nil
}
