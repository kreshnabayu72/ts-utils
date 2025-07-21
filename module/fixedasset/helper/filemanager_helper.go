package helper

import (
	"context"
	"fmt"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	utildb "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileManagerHelper interface {
	CreateFileManager(
		serviceFactory utilprovider.ServiceProvider,
		companyCode string,
		collection string,
		data model.FileCommon,
		fileStorageAttribute model.FileStorageAttribute,
	) *entity.HttpError
	DeleteFileManager(
		serviceFactory utilprovider.ServiceProvider,
		companyCode string,
		collection string,
		filter primitive.M,
		fileStorageAttribute model.FileStorageAttribute,
	) *entity.HttpError
}

func MakeFileManagerHelperImplHelperImpl() *FileManagerHelperImpl {
	return &FileManagerHelperImpl{}
}

type FileManagerHelperImpl struct{}

func (f *FileManagerHelperImpl) CreateFileManager(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	collection string,
	data model.FileCommon,
	fileStorageAttribute model.FileStorageAttribute,
) *entity.HttpError {
	// ─────── Check & update storage ───────
	adminDb := utildb.AdminDBName
	clientService := serviceFactory.MakeService(adminDb, db.ClientModelName)

	var client *model.Company
	err := clientService.FindOne(&client, context.Background(), bson.M{"clientId": companyCode})
	if err != nil {
		return &entity.HttpError{}
	}

	if client.FileStorage.UsedKB+data.FileSizeKb > client.FileStorage.TotalKB {
		return entity.BadRequest("Company file storage is full")
	}

	_, err = clientService.UpdateOne(
		context.Background(),
		bson.M{"clientId": companyCode},
		bson.M{
			"$inc": bson.M{
				"fileStorage.usedKB":                                data.FileSizeKb,
				"fileStorage.availableKB":                           -1 * data.FileSizeKb,
				fmt.Sprintf("fileStorage.%v", fileStorageAttribute): data.FileSizeKb,
			},
		},
	)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	// ─────── Insert file ───────
	db := utildb.FixedAssetDBName.CompanyDBName(companyCode)

	fileService := serviceFactory.MakeService(db, collection)

	_, err = fileService.InsertOne(context.Background(), data)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}

func (f *FileManagerHelperImpl) DeleteFileManager(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	collection string,
	filter primitive.M,
	fileStorageAttribute model.FileStorageAttribute,
) *entity.HttpError {
	database := utildb.FixedAssetDBName.CompanyDBName(companyCode)

	fileService := serviceFactory.MakeService(database, collection)

	var data model.FileCommon
	err := fileService.FindOne(&data, context.Background(), filter)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	_, err = fileService.DeleteOne(context.Background(), filter)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	// ─────── Check & update storage ───────
	adminDb := utildb.AdminDBName
	clientService := serviceFactory.MakeService(adminDb, db.ClientModelName)

	var client *model.Company
	err = clientService.FindOne(&client, context.Background(), bson.M{"clientId": companyCode})
	if err != nil {
		return &entity.HttpError{}
	}

	_, err = clientService.UpdateOne(
		context.Background(),
		bson.M{"clientId": companyCode},
		bson.M{
			"$inc": bson.M{
				"fileStorage.usedKB":                                -1 * data.FileSizeKb,
				"fileStorage.availableKB":                           data.FileSizeKb,
				fmt.Sprintf("fileStorage.%v", fileStorageAttribute): -1 * data.FileSizeKb,
			},
		},
	)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}
