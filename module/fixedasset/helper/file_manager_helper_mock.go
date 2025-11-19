package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/v2/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/v2/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
	"github.com/stretchr/testify/mock"
)

type MockFileManagerHelperSQLImpl struct {
	mock.Mock
}

func (m *MockFileManagerHelperSQLImpl) CreateFileManagerSQL(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	data new_model.FileInStorage,
) *entity.HttpError {
	args := m.Called(serviceFactory, companyCode, data)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockFileManagerHelperSQLImpl) DeleteFileManagerSQL(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	filter map[string]sql_query.SQLCondition,
	fileType new_model.FileType,
) *entity.HttpError {
	args := m.Called(serviceFactory, companyCode, filter, fileType)
	return args.Get(0).(*entity.HttpError)
}
