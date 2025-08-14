package helper

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	util_fixedasset_new_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/stretchr/testify/mock"
)

type MockNotificationSQLHelper struct {
	mock.Mock
}

func (m *MockNotificationSQLHelper) SendNotificationToTransactionRolesSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	groupIdFilter sql_query.SQLCondition,
	transactionType []util_fixedasset_new_model.TransactionRoleAttribute,
	roleType []util_fixedasset_new_model.TransactionRoleType,
	notificationDataWithoutUser util_fixedasset_new_model.Notification,
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

func (m *MockNotificationSQLHelper) SendNotificationToSystemRolesSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	groupIdFilter sql_query.SQLCondition,
	roleAttribute []util_fixedasset_new_model.SystemRoleAttribute,
	notificationDataWithoutUser util_fixedasset_new_model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(
		serviceProvider,
		companyCode,
		groupIdFilter,
		roleAttribute,
		notificationDataWithoutUser,
	)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) SendNotificationToUsersSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	userIds []string,
	notificationDataWithoutUser util_fixedasset_new_model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(serviceProvider, companyCode, userIds, notificationDataWithoutUser)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) SendNotificationSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	notifications []util_fixedasset_new_model.Notification,
) *entity.HttpError {
	args := m.Mock.Called(serviceProvider, companyCode, notifications)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) RefreshNotificationSseSQL(userIds []string) *entity.HttpError {
	args := m.Mock.Called(userIds)
	return args.Get(0).(*entity.HttpError)
}
