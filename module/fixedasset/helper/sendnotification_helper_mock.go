package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockNotificationHelper struct {
	mock.Mock
}

func (m *MockNotificationHelper) SendNotificationToTransactionRoles(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	groupIdFilter primitive.M,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(
		serviceProvider,
		companyCode,
		groupIdFilter,
		transactionType,
		roleType,
		notificationDataWithoutUser,
	)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationHelper) SendNotificationToSystemRoles(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	groupIdFilter primitive.M,
	roleType []model.RoleType,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(
		serviceProvider,
		companyCode,
		groupIdFilter,
		roleType,
		notificationDataWithoutUser,
	)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationHelper) SendNotificationToUsers(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	userIds []primitive.ObjectID,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(serviceProvider, companyCode, userIds, notificationDataWithoutUser)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationHelper) SendNotification(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	notifications []model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(serviceProvider, companyCode, notifications)
	return args.Get(0).(*entity.HttpError)
}
