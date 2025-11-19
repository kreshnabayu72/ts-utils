package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/v2/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/model"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/v2/provider"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	filter primitive.M,
	fileStorageAttribute model.FileStorageAttribute,
) *entity.HttpError {
	args := m.Mock.Called(serviceFactory, companyCode, collection, filter, fileStorageAttribute)
	return args.Get(0).(*entity.HttpError)
}
