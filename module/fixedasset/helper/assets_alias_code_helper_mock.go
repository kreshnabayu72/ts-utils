package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"github.com/stretchr/testify/mock"
)

type MockAssetsAliasCodeHelperImpl struct {
	mock.Mock
}

func (m *MockAssetsAliasCodeHelperImpl) UpdateAssetsAliasCode(
	ctx context.Context,
	fixedAssetService service.PostgreSqlService,
	companyCode string,
) *entity.HttpError {
	args := m.Called(ctx, fixedAssetService, companyCode)
	return args.Get(0).(*entity.HttpError)
}
