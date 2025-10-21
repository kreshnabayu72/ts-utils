package helper

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	util_fixedasset_new_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
)

type NotificationHelperSQL interface {
	SendNotificationToTransactionRolesSQL(
		serviceProvider provider.ServiceProvider,
		companyCode string,
		groupIdFilter sql_query.SQLCondition,
		transactionType []util_fixedasset_new_model.TransactionRoleAttribute,
		roleType []util_fixedasset_new_model.TransactionRoleType,
		notificationDataWithoutUser util_fixedasset_new_model.Notification,
		approvalLevel ...int,
	) *entity.HttpError

	SendNotificationToSystemRolesSQL(
		serviceProvider provider.ServiceProvider,
		companyCode string,
		groupIdFilter sql_query.SQLCondition,
		roleAttribute []util_fixedasset_new_model.SystemRoleAttribute,
		notificationDataWithoutUser util_fixedasset_new_model.Notification,
	) *entity.HttpError

	SendNotificationToUsersSQL(
		serviceProvider provider.ServiceProvider,
		companyCode string,
		userIds []string,
		notificationDataWithoutUser util_fixedasset_new_model.Notification,
	) *entity.HttpError

	SendNotificationSQL(
		serviceProvider provider.ServiceProvider,
		companyCode string,
		notifications []util_fixedasset_new_model.Notification,
	) *entity.HttpError

	RefreshNotificationSseWithFilter(
		ctx context.Context,
		serviceProvider provider.ServiceProvider,
		companyCode string,
		filter any,
	) error

	RefreshNotificationSseSQL(userIds []string) *entity.HttpError
	EmitApprovalSseSQL(userIds []string) *entity.HttpError
	EmitLoggedOutSseSQL(userIds []string) *entity.HttpError
	EmitHandoverSseSQL(userIds []string) *entity.HttpError
	UpdateUserSseSQL(userIds []string) *entity.HttpError
	EmitLoggedOutByAdminSseSQL(userIds []string) *entity.HttpError
}

type SendNotificationSQLHelperImpl struct {
	AccessToken           string
	MobileNotificationURI string
}

func MakeNotificationSQLHelperImpl() NotificationHelperSQL {
	result := &SendNotificationSQLHelperImpl{
		MobileNotificationURI: "https://fcm.googleapis.com/v1/projects/tag-samurai/messages:send",
	}
	result.GetAccessTokenSQL()
	return result
}

func (r *SendNotificationSQLHelperImpl) GetAccessTokenSQL() error {
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
	//Add timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	tok, err := cfg.TokenSource(ctx).Token()
	if err != nil {
		return fmt.Errorf("failed to retrieve access token: %w", err)
	}

	r.AccessToken = tok.AccessToken
	return nil
}

func (r *SendNotificationSQLHelperImpl) SendNotificationToTransactionRolesSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	groupIdFilter sql_query.SQLCondition,
	transactionType []util_fixedasset_new_model.TransactionRoleAttribute,
	roleType []util_fixedasset_new_model.TransactionRoleType,
	notificationDataWithoutUser util_fixedasset_new_model.Notification,
	approvalLevel ...int,
) *entity.HttpError {
	//Add timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	userTransactionRoleService := serviceProvider.MakePostgreSqlService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
	)

	builder := sql_query.NewSQLSelectBuilder[any](db.UserTransactionRoleTableName, "utr").
		Select("utr.user_id").
		Join("groups g", "g.ancestor_id = utr.group_id").
		Where(map[string]sql_query.SQLCondition{
			"is_active":      {Operator: sql_query.SQLOperatorEqual, Value: true},
			"role_type":      {Operator: sql_query.SQLOperatorIn, Value: roleType},
			"role_attribute": {Operator: sql_query.SQLOperatorIn, Value: transactionType},
			"g.id":           groupIdFilter,
		})

	if len(approvalLevel) > 0 && approvalLevel[0] != 0 {
		builder.Where(map[string]sql_query.SQLCondition{
			"approval_level": {Operator: sql_query.SQLOperatorEqual, Value: approvalLevel[0]},
		})
	}

	query, args, _ := builder.Build()

	var roles []map[string]int
	err := userTransactionRoleService.SelectMany(
		&roles,
		ctx,
		query,
		args...,
	)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return entity.InternalServerError(err.Error())
	}

	log.Printf("Found %v roles for notif", len(roles))

	var uniqueUserIds []string
	var uniqueUserIdsCheck map[string]bool = make(map[string]bool)

	for _, role := range roles {
		// if user id already used, continue/skip
		if _, ok := uniqueUserIdsCheck[strconv.Itoa(role["user_id"])]; ok {
			continue
		}

		uniqueUserIds = append(uniqueUserIds, strconv.Itoa(role["user_id"]))
		uniqueUserIdsCheck[strconv.Itoa(role["user_id"])] = true
	}

	return r.SendNotificationToUsersSQL(
		serviceProvider,
		companyCode,
		uniqueUserIds,
		notificationDataWithoutUser,
	)
}

func (r *SendNotificationSQLHelperImpl) SendNotificationToSystemRolesSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	groupIdFilter sql_query.SQLCondition,
	roleAttribute []util_fixedasset_new_model.SystemRoleAttribute,
	notificationDataWithoutUser util_fixedasset_new_model.Notification,
) *entity.HttpError {
	//Add timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	userSystemRoleService := serviceProvider.MakePostgreSqlService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
	)

	query, args, _ := sql_query.NewSQLSelectBuilder[util_fixedasset_new_model.UserSystemRole](
		db.UserSystemRoleTableName,
	).
		Where(map[string]sql_query.SQLCondition{
			"is_active":      {Operator: sql_query.SQLOperatorEqual, Value: true},
			"role_attribute": {Operator: sql_query.SQLOperatorIn, Value: roleAttribute},
			"group_id":       groupIdFilter,
		}).
		Build()

	var roles []util_fixedasset_new_model.UserSystemRole
	err := userSystemRoleService.SelectMany(
		&roles,
		ctx,
		query,
		args...,
	)

	log.Printf("Found %v roles for notif", len(roles))

	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	var uniqueUserIds []string
	var uniqueUserIdsCheck map[string]bool = make(map[string]bool)

	for _, role := range roles {
		// if user id already used, continue/skip
		if _, ok := uniqueUserIdsCheck[strconv.Itoa(role.UserID)]; ok {
			continue
		}

		uniqueUserIds = append(uniqueUserIds, strconv.Itoa(role.UserID))
		uniqueUserIdsCheck[strconv.Itoa(role.UserID)] = true
	}

	return r.SendNotificationToUsersSQL(
		serviceProvider,
		companyCode,
		uniqueUserIds,
		notificationDataWithoutUser,
	)
}

func (r *SendNotificationSQLHelperImpl) SendNotificationToUsersSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	userIds []string,
	notificationDataWithoutUser util_fixedasset_new_model.Notification,
) *entity.HttpError {
	//Add timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	log.Printf("SENDING NOTIF %v to %v user", notificationDataWithoutUser.Title, len(userIds))
	if len(userIds) == 0 {
		return nil
	}
	userService := serviceProvider.MakePostgreSqlService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
	)
	query, args, _ := sql_query.NewSQLSelectBuilder[util_fixedasset_new_model.User](
		db.UserTableName,
	).
		Where(map[string]sql_query.SQLCondition{
			"id": {Operator: sql_query.SQLOperatorIn, Value: userIds},
		}).
		Build()

	var users []util_fixedasset_new_model.User
	err := userService.SelectMany(
		&users,
		ctx,
		query,
		args...,
	)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		return entity.InternalServerError(err.Error())
	}

	var pendingNotifications []util_fixedasset_new_model.Notification
	var fcmTokens []string

	for _, user := range users {
		pendingNotifications = append(pendingNotifications, util_fixedasset_new_model.Notification{
			ID:         primitive.NewObjectID(),
			Title:      notificationDataWithoutUser.Title,
			Detail:     notificationDataWithoutUser.Detail,
			App:        notificationDataWithoutUser.App,
			UserKey:    strconv.Itoa(user.ID),
			FCMToken:   user.FCMToken,
			IsRead:     false,
			IsReadOnly: notificationDataWithoutUser.IsReadOnly,
			Module:     notificationDataWithoutUser.Module,
			Status:     notificationDataWithoutUser.Status,
			Severity:   notificationDataWithoutUser.Severity,
			Category:   notificationDataWithoutUser.Category,
			Type:       notificationDataWithoutUser.Type,
			Data:       notificationDataWithoutUser.Data,
			Plan:       notificationDataWithoutUser.Plan,
			IsManager:  notificationDataWithoutUser.IsManager,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})

		if user.FCMToken != "" {
			fcmTokens = append(fcmTokens, user.FCMToken)
		}
	}

	sendNotifErr := r.SendNotificationSQL(serviceProvider, companyCode, pendingNotifications)

	if sendNotifErr != nil {
		log.Printf("Error: %v\n", sendNotifErr.Error())

		return entity.ToHttpError(sendNotifErr)
	}

	mobileNotifErr := r.SendMobileNotificationSQL(notificationDataWithoutUser, fcmTokens)

	if mobileNotifErr != nil {
		log.Println("MOBILE NOTIFICATION ERROR: " + mobileNotifErr.Error())
	}

	return nil
}

func (r *SendNotificationSQLHelperImpl) SendNotificationSQL(
	serviceProvider provider.ServiceProvider,
	companyCode string,
	notifications []util_fixedasset_new_model.Notification,
) *entity.HttpError {
	if len(notifications) == 0 {
		return nil
	}

	notificationApi := os.Getenv("NOTIFICATION_API")

	notificationService := serviceProvider.MakeService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
		db.NotificationModelName,
	)

	notificationsInterface := []any{}
	for _, pendingNotification := range notifications {
		pendingNotification.CreatedAt = time.Now()
		pendingNotification.UpdatedAt = time.Now()
		notificationsInterface = append(notificationsInterface, pendingNotification)

		go func(p util_fixedasset_new_model.Notification) {
			resp, err := http.Get(
				fmt.Sprintf(
					"%s/v2/sse/send?id=%s&message=user:unread-notification",
					notificationApi,
					p.UserKey,
				),
			)
			if err != nil {
				log.Println("NOTIFICATION ERROR: " + err.Error())
				return
			}

			// Drain & close body so the connection can be reused
			_, _ = io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}(pendingNotification)

	}

	_, err := notificationService.InsertMany(context.Background(), notificationsInterface)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		return entity.InternalServerError(err.Error())
	}

	return nil
}

func (r *SendNotificationSQLHelperImpl) SendMobileNotificationSQL(
	data util_fixedasset_new_model.Notification,
	tokens []string,
) *entity.HttpError {
	// ─────── credentials ───────
	cred := os.Getenv("MOBILE_CREDENTIAL_JSON")
	opt := option.WithCredentialsJSON([]byte(cred))

	// ─────── init firebase app ───────
	//Add timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	config := &firebase.Config{
		ProjectID: "tag-samurai", // <-- explicitly set your Firebase project ID
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Printf("error getting Messaging client: %v\n", err)
	}

	if client == nil {
		log.Println("Koneksi client gagal")
		return nil
	}

	dataPayload := make(map[string]string)

	for k, v := range data.Data {
		if strVal, ok := v.(string); ok {
			// The value 'v' was a string, so we can safely add it.
			dataPayload[k] = strVal
			continue
		}

		log.Printf("Notification data attribute isnt string: %v (type=%v)\n", v, reflect.TypeOf(v))
	}

	// Tambahkan field lain sebagai string
	dataPayload["_id"] = data.ID.Hex()
	dataPayload["title"] = data.Title
	dataPayload["detail"] = data.Detail
	dataPayload["user"] = data.UserKey
	dataPayload["isReadOnly"] = fmt.Sprintf("%v", data.IsReadOnly)
	dataPayload["isManager"] = fmt.Sprintf("%v", data.IsManager)
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

func emitRefreshSse(userIds []string, message string, reason ...string) *entity.HttpError {
	notificationApi := os.Getenv("NOTIFICATION_API")
	if notificationApi == "" {
		log.Println("REFRESH SSE: NOTIFICATION_API not set")
		return nil
	}

	for _, userId := range userIds {
		path := fmt.Sprintf(
			"%s/v2/sse/send?id=%s&message=%s",
			notificationApi,
			url.QueryEscape(userId),
			message,
		)

		if len(reason) > 0 {
			path += fmt.Sprintf("&reason=%s", reason[0])
		}

		resp, err := http.Get(path)
		// Don't throw error, if notif is down, let transaction continue
		if err != nil {
			log.Println("RefreshNotificationSseSQL: " + err.Error())
			continue
		}

		// Drain & close body so the connection can be reused
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}

	return nil
}

func (r *SendNotificationSQLHelperImpl) RefreshNotificationSseSQL(
	userIds []string,
) *entity.HttpError {
	return emitRefreshSse(userIds, "user:unread-notification")
}

func (r *SendNotificationSQLHelperImpl) EmitApprovalSseSQL(
	userIds []string,
) *entity.HttpError {
	return emitRefreshSse(userIds, "user:approval")
}

func (r *SendNotificationSQLHelperImpl) EmitLoggedOutSseSQL(
	userIds []string,
) *entity.HttpError {
	return emitRefreshSse(userIds, "user:logout")
}

func (r *SendNotificationSQLHelperImpl) UpdateUserSseSQL(
	userIds []string,
) *entity.HttpError {
	return emitRefreshSse(userIds, "user:update")
}

func (r *SendNotificationSQLHelperImpl) EmitHandoverSseSQL(
	userIds []string,
) *entity.HttpError {
	return emitRefreshSse(userIds, "user:handover")
}

func (r *SendNotificationSQLHelperImpl) EmitLoggedOutByAdminSseSQL(
	userIds []string,
) *entity.HttpError {
	return emitRefreshSse(userIds, "user:logout", "LOGGED_OUT_BY_ADMINISTRATOR")
}

func (r *SendNotificationSQLHelperImpl) RefreshNotificationSseWithFilter(
	ctx context.Context,
	serviceProvider provider.ServiceProvider,
	companyCode string,
	filter any,
) error {
	notificationService := serviceProvider.MakeService(
		db.FixedAssetDBName.CompanyDBName(companyCode),
		db.NotificationModelName,
	)

	// Find notifications based on the provided filter
	var notifications []util_fixedasset_new_model.Notification
	err := notificationService.Find(
		&notifications,
		ctx,
		filter,
	)
	if err != nil {
		return err
	}

	// Update the IsRead field to true for each notification
	_, err = notificationService.UpdateMany(
		ctx,
		filter,
		bson.M{
			"$set": bson.M{"is_read": true, "updated_at": time.Now()},
		},
	)
	if err != nil {
		return err
	}

	var userIds []string
	for _, notif := range notifications {
		userIds = append(userIds, notif.UserKey)
	}

	emitErr := r.RefreshNotificationSseSQL(userIds)
	if emitErr != nil {
		return emitErr
	}

	return nil
}
