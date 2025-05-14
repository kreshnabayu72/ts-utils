package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2/jwt"
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

type SendNotificationHelperImpl struct {
	AccessToken           string
	MobileNotificationURI string
}

func MakeNotificationHelperImpl() *SendNotificationHelperImpl {
	result := &SendNotificationHelperImpl{

		MobileNotificationURI: "https://fcm.googleapis.com/v1/projects/tag-samurai/messages:send",
	}
	result.GetAccessToken()
	return result
}

func (r *SendNotificationHelperImpl) GetAccessToken() error {
	// 1. Read the raw private key from env and replace literal "\n" with actual newlines
	rawKey := os.Getenv("MOBILE_PRIVATE_KEY")
	if rawKey == "" {
		return fmt.Errorf("MOBILE_PRIVATE_KEY is not set")
	}
	privateKey := strings.ReplaceAll(rawKey, `\n`, "\n")

	// 2. Configure the JWT with your service-account details
	cfg := &jwt.Config{
		Email:        "firebase-adminsdk-53b72@tag-samurai.iam.gserviceaccount.com",
		PrivateKey:   []byte(privateKey),
		PrivateKeyID: "efd5f98c2704770fd75b4e65d977c412d093ec19",
		Scopes:       []string{"https://www.googleapis.com/auth/firebase.messaging"},
		TokenURL:     "https://oauth2.googleapis.com/token",
	}

	// 3. Request a token
	tok, err := cfg.TokenSource(context.Background()).Token()
	if err != nil {
		return fmt.Errorf("failed to retrieve access token: %w", err)
	}

	r.AccessToken = tok.AccessToken
	return nil
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
		//Dont throw error, if notif is down, let transaction continue
		if err != nil {
			fmt.Println("NOTIFICATION ERROR: " + err.Error())
		}

		mobileNotifErr := r.SendMobileNotification(pendingNotification)

		if mobileNotifErr != nil {
			fmt.Println("MOBILE NOTIFICATION ERROR: " + mobileNotifErr.Error())
		}
	}

	_, err := notificationService.InsertMany(context.Background(), notificationsInterface)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}

func (r *SendNotificationHelperImpl) SendMobileNotification(data model.Notification) *entity.HttpError {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	// build payload
	payload := map[string]interface{}{
		"message": map[string]interface{}{
			"topic": data.User,
			"data": func() map[string]interface{} {
				// copy original data and add title/module
				d := make(map[string]interface{}, len(data.Data)+2)
				for k, v := range data.Data {
					d[k] = v
				}
				d["title"] = data.Title
				d["module"] = data.Module
				return d
			}(),
			"notification": map[string]string{
				"body":  data.Detail,
				"title": data.Title,
			},
			"android": map[string]interface{}{
				"priority": "HIGH",
				"notification": map[string]string{
					"click_action": "MOBILE_NOTIFICATION",
				},
			},
		},
	}

	// marshal JSON
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal notification body: %v", err)
		return entity.InternalServerError(err.Error())
	}

	// create request with context
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, r.MobileNotificationURI, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Printf("error creating HTTP request: %v", err)
		return entity.InternalServerError(err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+r.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// do the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("error sending notification: %v", err)
		return entity.InternalServerError(err.Error())

	}
	defer resp.Body.Close()

	// optional: decode and log response body
	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		log.Printf("error decoding response: %v", err)
		return entity.InternalServerError(err.Error())

	}
	log.Printf("notification response: %+v", respBody)

	return nil
}
