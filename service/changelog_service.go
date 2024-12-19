package service

import (
	"context"
	"time"

	"github.com/betsv2/tsv2-go-utils/db"
	"github.com/betsv2/tsv2-go-utils/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IChangelogUseCase interface {
	CreateChangelog(ctx context.Context, data model.Changelog, userId primitive.ObjectID) error
}

type ChangelogUseCase struct {
	ChangelogService Service
	UserService      Service
}

// Deprecated: user service is manually created outside
func MakeChangelogUseCase(companyCode string, user Service) IChangelogUseCase {
	changelogService := MakeService(companyCode, db.ChangelogModelName)
	userService := user

	return &ChangelogUseCase{
		ChangelogService: changelogService,
		UserService:      userService,
	}
}

func NewChangelogUseCase(companyCode string) IChangelogUseCase {
	return &ChangelogUseCase{
		ChangelogService: MakeService(companyCode, db.ChangelogModelName),
		UserService:      MakeService(companyCode, db.UserModelName),
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
