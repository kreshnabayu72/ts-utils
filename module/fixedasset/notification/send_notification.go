package notification

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

func SendNotificationToTransactionRoles(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	groupIdFilter primitive.M,
	transactionType []model.TransactionGroupAttribute,
	roleType []model.TransactionRoleType,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	userTransactionRoleService := serviceProvider.MakeService(db.FixedAssetDBName.CompanyDBName(companyCode), util_db.UserTransactionRoleModelName)

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
		//if user id already used, continue/skip
		if _, ok := uniqueUserIdsCheck[role.User.ID]; ok {
			continue
		}

		uniqueUserIds = append(uniqueUserIds, role.User.ID)
		uniqueUserIdsCheck[role.User.ID] = true
	}

	return SendNotificationToUsers(serviceProvider, companyCode, uniqueUserIds, notificationDataWithoutUser)
}

func SendNotificationToSystemRoles(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	groupIdFilter primitive.M,
	roleType []model.RoleType,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	userTransactionRoleService := serviceProvider.MakeService(db.FixedAssetDBName.CompanyDBName(companyCode), util_db.UserTransactionRoleModelName)

	var roles []model.UserSystemRole
	err := userTransactionRoleService.Find(
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
		//if user id already used, continue/skip
		if _, ok := uniqueUserIdsCheck[*role.User.ID]; ok {
			continue
		}

		uniqueUserIds = append(uniqueUserIds, *role.User.ID)
		uniqueUserIdsCheck[*role.User.ID] = true
	}

	return SendNotificationToUsers(serviceProvider, companyCode, uniqueUserIds, notificationDataWithoutUser)
}

func SendNotificationToUsers(
	serviceProvider util_provider.ServiceProvider,
	companyCode string,
	userIds []primitive.ObjectID,
	notificationDataWithoutUser model.Notification,
) *entity.HttpError {
	notificationApi := os.Getenv("NOTIFICATION_API")
	var pendingNotifications []model.Notification

	notificationService := serviceProvider.MakeService(db.FixedAssetDBName.CompanyDBName(companyCode), util_db.NotificationModelName)

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

	notificationsInterface := []interface{}{}
	for _, pendingNotification := range pendingNotifications {
		notificationsInterface = append(notificationsInterface, pendingNotification)
		_, err := http.Get(fmt.Sprintf("%s/v2/sse/send?id=%s&message=user:unread-notification", notificationApi, pendingNotification.User.Hex()))
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
