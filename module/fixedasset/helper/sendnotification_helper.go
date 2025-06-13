package helper

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"gitea.qwertysystem.net/BETS/ts-utils/db"
	util_db "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
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
	userService := serviceProvider.MakeService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
		util_db.UserModelName,
	)

	var users []model.User
	err := userService.Find(
		&users,
		context.Background(),
		bson.M{
			"_id": bson.M{"$in": userIds},
		},
	)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	var pendingNotifications []model.Notification
	var fcmTokens []string

	for _, user := range users {
		pendingNotifications = append(pendingNotifications, model.Notification{
			ID:         primitive.NewObjectID(),
			Title:      notificationDataWithoutUser.Title,
			Detail:     notificationDataWithoutUser.Detail,
			App:        notificationDataWithoutUser.App,
			User:       user.ID,
			FCMToken:   user.FCMToken,
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

		if user.FCMToken != "" {
			fcmTokens = append(fcmTokens, user.FCMToken)
		}
	}

	mobileNotifErr := r.SendMobileNotification(notificationDataWithoutUser, fcmTokens)

	if mobileNotifErr != nil {
		fmt.Println("MOBILE NOTIFICATION ERROR: " + mobileNotifErr.Error())
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

	}

	_, err := notificationService.InsertMany(context.Background(), notificationsInterface)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}

func (r *SendNotificationHelperImpl) SendMobileNotification(data model.Notification, tokens []string) *entity.HttpError {
	// ─────── credentials ───────
	cred := os.Getenv("MOBILE_CREDENTIAL_JSON")
	opt := option.WithCredentialsJSON([]byte(cred))

	// ─────── init firebase app ───────
	ctx := context.Background()

	config := &firebase.Config{
		ProjectID: "tag-samurai", // <-- explicitly set your Firebase project ID
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err)
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Printf("error getting Messaging client: %v\n", err)
	}

	dataPayload := make(map[string]string)

	for k, v := range data.Data {
		dataPayload[k] = v.(string)
	}

	// Tambahkan field lain sebagai string
	dataPayload["_id"] = data.ID.Hex()
	dataPayload["title"] = data.Title
	dataPayload["detail"] = data.Detail
	dataPayload["user"] = data.User.Hex()
	dataPayload["isReadOnly"] = fmt.Sprintf("%v", data.IsReadOnly)
	dataPayload["plan"] = data.Plan
	dataPayload["module"] = string(data.Module)
	dataPayload["status"] = string(data.Status)
	dataPayload["severity"] = string(data.Severity)
	dataPayload["category"] = string(data.Category)
	dataPayload["type"] = data.Type

	// --- 2. Bangun Pesan Multicast Menggunakan Struct SDK ---
	// Ini cara yang lebih aman dan direkomendasikan daripada map[string]interface{}
	message := &messaging.MulticastMessage{
		Tokens: tokens, // Targetnya adalah daftar token, bukan satu topik
		Data:   dataPayload,
		Notification: &messaging.Notification{
			Title: data.Title,
			Body:  data.Detail,
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				ClickAction: "MOBILE_NOTIFICATION",
			},
		},
	}

	// --- 3. Kirim Pesan dengan SendMulticast ---
	br, err := client.SendEachForMulticast(context.Background(), message)
	if err != nil {
		log.Printf("Gagal mengirim notifikasi multicast: %v %v\n", err, message)
		return nil
	}

	// --- 4. (SANGAT PENTING) Tangani Respons & Bersihkan Token Gagal ---
	// Cek respons untuk setiap token yang dikirim.
	log.Printf("Berhasil mengirim notifikasi ke: %d perangkat\n", br.SuccessCount)
	log.Printf("Gagal mengirim notifikasi ke: %d perangkat\n", br.FailureCount)

	if br.FailureCount > 0 {
		var failedTokens []string
		for idx, resp := range br.Responses {
			if !resp.Success {
				// Token gagal, catat atau langsung hapus dari database bos.
				// Kode error 'messaging.IsRegistrationTokenNotRegistered' menandakan token sudah tidak valid.
				log.Printf("Gagal kirim ke token %s: %v\n", tokens[idx], resp.Error)
				failedTokens = append(failedTokens, tokens[idx])
			}
		}
		log.Printf("Daftar token yang gagal: %v", failedTokens)
	}
	return nil
}
