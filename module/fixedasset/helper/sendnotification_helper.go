package helper

import (
	"context"
	"encoding/json"
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

		fcmTokens = append(fcmTokens, user.FCMToken)
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
	// cred := os.Getenv("MOBILE_CREDENTIAL_JSON")
	// opt := option.WithCredentialsJSON([]byte(cred))

	sdk := bson.M{
		"type":                        "service_account",
		"project_id":                  "tag-samurai",
		"private_key_id":              "d19ee481e128dff14adee2ee60199cdcec5da6f9",
		"private_key":                 "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCprSiwaneeldw2\nnmlJhxeEpkpn+R6xGJQu7PZkPOGYTbA+kddSpCLFHY55YFYqwpbFgRlxwzJNyQ5S\nNTdkjcSXpOe6zRPgEDTV+XSm1YLdG/OCwCQjrlP5d/bH9T4IolgZOwTQteuTNTGs\nkWpfyz3ldhUd77QonbCGNFSDmjQlE1vTbUd9alg2ILRIe0Vp9KHZ4p3YrNvb112B\n0xKRYbwOhntagVOkx9BwU8N8AYdst7reKdY69dZuxs+cv/5X+gnnQBiAyZk9kbKS\nTg6CIRIDC142mlujlL8SgcCxTJuEDlQF4yvf411ybUt11ieY1C0y5isQfHiPUUIA\nXtj5osxVAgMBAAECggEAKziWosmcCuN1otwbScWFsh6B2ZZVwRrXtFRcNxBjDt7L\nRET2KnDxAnES98wYZvAGY+0MxDPKvpG2ei5gG9cY+t/gv26531uBeff/RPgiS5xx\nHR9AlW40qdlrHWhVEvH3qwWjpl6Q0YC9VI3+IMSu0aKQon2K0MXKT4VI+fTvz3mB\n+CESD16h7Gwc97LTYquOw2b7ZqFpVzC3r8tp+1AlAtCGeKdjvETNwcfX8Frdn8Iw\n/IVSPj1JNsYCqXBKjWFOqqLWMjDAOJXVD8YZUHtaqdVCK5haY8pVK0HFWsE1uOW1\nFwR9nLC6GIfL6u0i+xZRPH0igcmNqR1sECKDvF90iQKBgQC4tUvkc/06zPR6X02m\nI0drm2Czf1bDHr+qE1R6PkVhBXO8xY5aPLFmplyUiJ5dD9TgCUKZV+IFq/3s9MjB\n7/cclPEtj2tTwkEqGGSa3TkHFPMV8DTJ2au+WRsA2Msx7xXQjaZuExd2s6U/3aZy\nBMnw/6O3XaLjFQQ0JrrMkjM7pwKBgQDrKpjXqrz9WEXv0NRjK9N/EerqifdMCjIA\nhwNtzifaDfzLwHDeDfDnX1VIR+aV4ZjGq7WKjKhLvtcYF1fwCzUCi7rFNA9xp//g\nkM3soJJdWh1fmf6M2UO9YT4CunDIOhrs1vOlVo9BWtx25YzDOQ3B2K+uobRj2Aqq\ns6mxk9PHowKBgQCGvbpfLT3NZzXxDxTV3WZ8WRxx1LK2I/gtWWyGwGZcIF4NlxAB\nQ8Ul3LYspHzh5mqkBB1RP4riGFVkRi0tTKA3eW92jOj+cPaYfrLlRfwT7nRlzJS1\nqtReRyIru/p+sqp3ON3v7n1vWh1yOX7yLgwiglorqWLpxKsMYDRkh0JYgQKBgQDg\naYnpZNlsuewjahkDejMds6+FmZsQSTpmHnE7O9vzzh1LIp2SiaM2aFct3s9AgEQW\nco2PZaLEhZ0lNEGKpGc1wrm8rm1+fF9FoWmhQanqw6/8JywupLZr5nwShnBR4oEH\nOijPwsZ6C4cPPwcQ1UU3ktwzvcQy81tagpCREQysGQKBgQCokjPV+yPbqh53aD5P\nGqHfkdnc+mADOI3H5w6t6B0HQ/rAbABAqiPXyCwAQFtRWQj5Skfk+cegNQ3WqmYR\nfcgtlMgrQkZCk4hmVxGFP/oJXhhTqi5gfXpHxLcBY3oi4hzQ9Q8aImjq6H/KzvvQ\njGiufBx0zcUBqihIEsPHgfgDZQ==\n-----END PRIVATE KEY-----\n",
		"client_email":                "firebase-adminsdk-53b72@tag-samurai.iam.gserviceaccount.com",
		"client_id":                   "103240756892132408402",
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-53b72%40tag-samurai.iam.gserviceaccount.com",
		"universe_domain":             "googleapis.com",
	}
	credBytes, err := json.Marshal(sdk)
	if err != nil {
		fmt.Printf("error marshalling cred: %v\n", err)
	}
	opt := option.WithCredentialsJSON(credBytes)

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
		log.Printf("Gagal mengirim notifikasi multicast: %v\n", err)
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
