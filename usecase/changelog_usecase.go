package service

import (
	"context"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/module/global/model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IChangelogUseCase interface {
	CreateChangelog(ctx context.Context, data model.Changelog, userId primitive.ObjectID) error
}

type ChangelogUseCase struct {
	ChangelogService service.Service
	UserService      service.Service
}

// Deprecated: user service is manually created outside
func MakeChangelogUseCase(companyCode string, user service.Service) IChangelogUseCase {
	dbName := db.GlobalDBName.CompanyDBName(companyCode)
	changelogService := service.MakeService(dbName, db.ChangelogModelName)
	userService := user

	return &ChangelogUseCase{
		ChangelogService: changelogService,
		UserService:      userService,
	}
}

func NewChangelogUseCase(companyCode string) IChangelogUseCase {
	dbName := db.GlobalDBName.CompanyDBName(companyCode)
	return &ChangelogUseCase{
		ChangelogService: service.MakeService(dbName, db.ChangelogModelName),
		UserService:      service.MakeService(dbName, db.UserModelName),
	}
}

func (u *ChangelogUseCase) CreateChangelog(
	ctx context.Context,
	data model.Changelog,
	userId primitive.ObjectID,
) error {
	var loggedInUser model.User
	loggedInErr := u.UserService.FindOne(&loggedInUser, ctx, bson.M{"_id": userId})
	if loggedInErr != nil {
		return loggedInErr
	}

	changelog := model.Changelog{
		ID:           primitive.NewObjectID(),
		Action:       data.Action,
		Field:        data.Field,
		OldValue:     data.OldValue,
		NewValue:     data.NewValue,
		Value:        data.Value,
		ModifiedBy:   loggedInUser.FirstName + " " + loggedInUser.LastName,
		ModifiedById: loggedInUser.ID,
		Object:       data.Object,
		ObjectId:     data.ObjectId,
		ObjectName:   data.ObjectName,
		AssetName:    data.AssetName,
		GroupId:      data.GroupId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := u.ChangelogService.InsertOne(ctx, changelog)
	if err != nil {
		return err
	}

	return nil
}

type MockChangelogUseCase struct {
	mock.Mock
}

func (m *MockChangelogUseCase) CreateChangelog(
	ctx context.Context,
	data model.Changelog,
	userId primitive.ObjectID,
) error {
	args := m.Called(ctx, data, userId)
	return args.Error(0)
}
