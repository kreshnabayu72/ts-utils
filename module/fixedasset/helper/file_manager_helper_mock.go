package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"github.com/stretchr/testify/mock"
)

type MockFileManagerHelper struct {
	mock.Mock
}

func (m *MockFileManagerHelper) CreateFileManager(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	collection string,
	data model.FileCommon,
	fileStorageAttribute model.FileStorageAttribute,
) *entity.HttpError {
	args := m.Mock.Called(serviceFactory, companyCode, collection, data, fileStorageAttribute)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockFileManagerHelper) DeleteFileManager(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	collection string,
	data model.FileCommon, // need to pass all to get totalkb etc
	fileStorageAttribute model.FileStorageAttribute,
) *entity.HttpError {
	args := m.Mock.Called(serviceFactory, companyCode, collection, data, fileStorageAttribute)
	return args.Get(0).(*entity.HttpError)
}
