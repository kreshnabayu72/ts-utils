package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"github.com/stretchr/testify/mock"
)

type MockAssetsAliasCodeHelperImpl struct {
	mock.Mock
}

func (m *MockAssetsAliasCodeHelperImpl) UpdateAssetsAliasCode(
	ctx context.Context,
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
) *entity.HttpError {
	args := m.Called(ctx, serviceFactory, companyCode)
	return args.Get(0).(*entity.HttpError)
}
