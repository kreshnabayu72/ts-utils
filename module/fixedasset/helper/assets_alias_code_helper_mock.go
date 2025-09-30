package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/stretchr/testify/mock"
)

type MockAssetsAliasCodeHelperImpl struct {
	mock.Mock
}

func (m *MockAssetsAliasCodeHelperImpl) UpdateAssetsAliasCode(
	ctx context.Context,
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	filter map[string]sql_query.SQLCondition,
) *entity.HttpError {
	args := m.Called(ctx, serviceFactory, companyCode, filter)
	return args.Get(0).(*entity.HttpError)
}
