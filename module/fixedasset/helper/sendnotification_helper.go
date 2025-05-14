package helper

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationHelper interface {
	SendNotificationToTransactionRoles(
		serviceProvider util_provider.ServiceProvider,
		companyCode string,
		groupIdFilter primitive.M,
		transactionType []model.TransactionGroupAttribute,
		roleType []model.TransactionRoleType,
		notificationDataWithoutUser model.Notification,
	) *entity.HttpError

	SendNotificationToSystemRoles(
		serviceProvider util_provider.ServiceProvider,
		companyCode string,
		groupIdFilter primitive.M,
		roleType []model.RoleType,
		notificationDataWithoutUser model.Notification,
	) *entity.HttpError

	SendNotificationToUsers(
		serviceProvider util_provider.ServiceProvider,
		companyCode string,
		userIds []primitive.ObjectID,
		notificationDataWithoutUser model.Notification,
	) *entity.HttpError

	SendNotification(
		serviceProvider util_provider.ServiceProvider,
		companyCode string,
		notifications []model.Notification,
	) *entity.HttpError
}

type SendNotificationHelperImpl struct{}

func MakeNotificationHelperImpl() *SendNotificationHelperImpl {
	return &SendNotificationHelperImpl{}
}

func (r *SendNotificationHelperImpl) SendNotificationToTransactionRoles(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	groupIdFilter primitive.M,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	userTransactionRoleService := serviceProvider.MakeService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
		util_db.UserTransactionRoleModelName,
	)

	var roles []model.UserTransactionRole
	err := userTransactionRoleService.Find(
		&roles,
		context.Background(),
		bson.M{
			"isActive":                  true,
			"roleType":                  bson.M{"$in": roleType},
			"transactionGroupAttribute": bson.M{"$in": transactionType},
			"group._id":                 groupIdFilter,
		},
	)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	var uniqueUserIds []primitive.ObjectID
	var uniqueUserIdsCheck map[primitive.ObjectID]bool = make(map[primitive.ObjectID]bool)

	for _, role := range roles {
		// if user id already used, continue/skip
		if _, ok := uniqueUserIdsCheck[role.User.ID]; ok {
			continue
		}

		uniqueUserIds = append(uniqueUserIds, role.User.ID)
		uniqueUserIdsCheck[role.User.ID] = true
	}

	return r.SendNotificationToUsers(
		serviceProvider,
		companyCode,
		uniqueUserIds,
		notificationDataWithoutUser,
	)
}

func (r *SendNotificationHelperImpl) SendNotificationToSystemRoles(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	groupIdFilter primitive.M,
	roleType []model.RoleType,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	userSystemRoleService := serviceProvider.MakeService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
		util_db.UserSystemRoleModelName,
	)

	var roles []model.UserSystemRole
	err := userSystemRoleService.Find(
		&roles,
		context.Background(),
		bson.M{
			"isActive":  true,
			"type":      bson.M{"$in": roleType},
			"group._id": groupIdFilter,
		},
	)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	var uniqueUserIds []primitive.ObjectID
	var uniqueUserIdsCheck map[primitive.ObjectID]bool = make(map[primitive.ObjectID]bool)

	for _, role := range roles {
		// if user id already used, continue/skip
		if _, ok := uniqueUserIdsCheck[*role.User.ID]; ok {
			continue
		}

		uniqueUserIds = append(uniqueUserIds, *role.User.ID)
		uniqueUserIdsCheck[*role.User.ID] = true
	}

	return r.SendNotificationToUsers(
		serviceProvider,
		companyCode,
		uniqueUserIds,
		notificationDataWithoutUser,
	)
}

func (r *SendNotificationHelperImpl) SendNotificationToUsers(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	userIds []primitive.ObjectID,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	if len(userIds) == 0 {
		return nil
	}

	var pendingNotifications []model.Notification

	for _, userId := range userIds {
		pendingNotifications = append(pendingNotifications, model.Notification{
			ID:         primitive.NewObjectID(),
			Title:      notificationDataWithoutUser.Title,
			Detail:     notificationDataWithoutUser.Detail,
			User:       userId,
			IsRead:     false,
			IsReadOnly: notificationDataWithoutUser.IsReadOnly,
			Module:     notificationDataWithoutUser.Module,
			Status:     notificationDataWithoutUser.Status,
			Severity:   notificationDataWithoutUser.Severity,
			Category:   notificationDataWithoutUser.Category,
			Type:       notificationDataWithoutUser.Type,
			Data:       notificationDataWithoutUser.Data,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})
	}

	return r.SendNotification(serviceProvider, companyCode, pendingNotifications)
}

func (r *SendNotificationHelperImpl) SendNotification(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	notifications []model.Notification,
) *entity.HttpError {
	if len(notifications) == 0 {
		return nil
	}

	notificationApi := os.Getenv("NOTIFICATION_API")

	notificationService := serviceProvider.MakeService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
		util_db.NotificationModelName,
	)

	notificationsInterface := []any{}
	for _, pendingNotification := range notifications {
		notificationsInterface = append(notificationsInterface, pendingNotification)
		_, err := http.Get(
			fmt.Sprintf(
				"%s/v2/sse/send?id=%s&message=user:unread-notification",
				notificationApi,
				pendingNotification.User.Hex(),
			),
		)
		if err != nil {
			return entity.InternalServerError(err.Error())
		}
	}

	_, err := notificationService.InsertMany(context.Background(), notificationsInterface)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}
