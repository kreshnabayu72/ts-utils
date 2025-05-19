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
)

type FileManagerHelper interface {
	CreateFileManager(
		serviceFactory utilprovider.ServiceProvider,
		companyCode string,
		collection string,
		data model.FileCommon,
		fileStorageAttribute model.FileStorageAttribute,
	) *entity.HttpError
}

type SendFileManagerHelperImpl struct {
	AccessToken          string
	MobileFileManagerURI string
}

func CreateFileManager(
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

	_, err = clientService.UpdateOne(context.Background(), bson.M{"clientId": companyCode},
		bson.M{"$inc": bson.M{"fileStorage.usedKB": data.FileSizeKb, "fileStorage.availableKB": -1 * data.FileSizeKb, fmt.Sprintf("fileStorage.%v", fileStorageAttribute): data.FileSizeKb}})
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
