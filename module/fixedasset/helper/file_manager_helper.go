package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	utildb "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type FileManagerHelperSQL interface {
	CreateFileManagerSQL(
		serviceFactory utilprovider.ServiceProvider,
		companyCode string,
		data new_model.FileInStorage,
	) *entity.HttpError

	DeleteFileManagerSQL(
		serviceFactory utilprovider.ServiceProvider,
		companyCode string,
		filter map[string]sql_query.SQLCondition,
		fileType new_model.FileType,
	) *entity.HttpError
}

type FileManagerHelperSQLImpl struct{}

func MakeFileManagerHelperSQLImpl() *FileManagerHelperSQLImpl {
	return &FileManagerHelperSQLImpl{}
}

func (f *FileManagerHelperImpl) CreateFileManagerSQL(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	data new_model.FileInStorage,
) *entity.HttpError {
	// ─────── Check & update storage ───────
	adminDb := utildb.AdminDBName
	clientService := serviceFactory.MakePostgreSqlService(adminDb, db.ClientTableName)

	// Check if the company exists
	queryStr, args, _ := sql_query.NewSQLSelectBuilder[int](db.ClientTableName).
		Where(map[string]sql_query.SQLCondition{
			"cleint_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    companyCode,
			},
		}).
		SetLimit(1).
		Build()

	var client new_model.Company
	err := clientService.SelectOne(&client, context.Background(), queryStr, args...)
	if err != nil {
		return &entity.HttpError{Message: "Company not found"}
	}

	// Check if the file storage has enough space
	if client.FileStorageUsedKB+data.FileSizeKB > client.FileStorageTotalKB {
		return entity.BadRequest("Company file storage is full")
	}

	// Update the file storage usage

	// _, err = clientService.UpdateOne(
	// 	map[string]interface{}{"clientId": companyCode},
	// 	map[string]interface{}{
	// 		"$inc": map[string]interface{}{
	// 			"fileStorage.usedKB": data.FileSizeKb,
	// 		},
	// 	},
	// )
	// if err != nil {
	// 	return &entity.HttpError{Message: "Failed to update file storage"}
	// }

	return nil
}
