package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/entity"
	util_fixedasset_new_model "gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/sql_query"
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
	approvalLevel ...int,
) *entity.HttpError {
	var args mock.Arguments

	if len(approvalLevel) > 0 {
		args = m.Mock.Called(
			serviceProvider,
			companyCode,
			groupIdFilter,
			transactionType,
			roleType,
			notificationDataWithoutUser,
			approvalLevel,
		)
	} else {
		args = m.Mock.Called(
			serviceProvider,
			companyCode,
			groupIdFilter,
			transactionType,
			roleType,
			notificationDataWithoutUser,
		)
	}
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

func (m *MockNotificationSQLHelper) EmitApprovalSseSQL(userIds []string) *entity.HttpError {
	args := m.Mock.Called(userIds)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) EmitLoggedOutSseSQL(userIds []string) *entity.HttpError {
	args := m.Mock.Called(userIds)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) EmitHandoverSseSQL(userIds []string) *entity.HttpError {
	args := m.Mock.Called(userIds)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) UpdateUserSseSQL(userIds []string) *entity.HttpError {
	args := m.Mock.Called(userIds)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) EmitLoggedOutByAdminSseSQL(userIds []string) *entity.HttpError {
	args := m.Mock.Called(userIds)
	return args.Get(0).(*entity.HttpError)
}

func (m *MockNotificationSQLHelper) RefreshNotificationSseWithFilter(
	ctx context.Context,
	serviceProvider provider.ServiceProvider,
	companyCode string,
	filter any,
) error {
	args := m.Mock.Called(ctx, serviceProvider, companyCode, filter)
	return args.Error(0)
}

func (m *MockNotificationSQLHelper) EmitUpdateIotReaderSseSQL(
	readerID []string,
) *entity.HttpError {
	args := m.Mock.Called(readerID)
	return args.Get(0).(*entity.HttpError)
}
