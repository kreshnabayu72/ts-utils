package helper

import (
	"context"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	utildb "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	admin_model "gitea.qwertysystem.net/BETS/ts-utils/module/admin/new_model"
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

func (f *FileManagerHelperSQLImpl) CreateFileManagerSQL(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	data new_model.FileInStorage,
) *entity.HttpError {
	// ─────── Check & update storage ───────
	adminDb := utildb.AdminDBName
	clientService := serviceFactory.MakePostgreSqlService(adminDb)

	// Check if the company exists
	queryStr, args, _ := sql_query.NewSQLSelectBuilder[admin_model.Client](db.ClientTableName).
		Where(map[string]sql_query.SQLCondition{
			"client_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    companyCode,
			},
		}).
		SetLimit(1).
		Build()

	var client admin_model.Client
	err := clientService.SelectOne(&client, context.Background(), queryStr, args...)
	if err != nil {
		return &entity.HttpError{Message: "Company not found"}
	}

	// Check if the file storage has enough space
	if client.FileStorageUsedKB+data.FileSizeKB > client.FileStorageTotalKB {
		return entity.BadRequest("Company file storage is full")
	}

	// Update the file storage usage
	storageFieldName := ""
	switch data.Type {
	case new_model.FileTypeAttachment:
		storageFieldName = "file_storage_attachment_kb"
	case new_model.FileTypeDamage:
		storageFieldName = "file_storage_damaged_kb"
	case new_model.FileTypeDisposal:
		storageFieldName = "file_storage_disposal_kb"
	case new_model.FileTypeRepairTicketing:
		storageFieldName = "file_storage_repair_ticketing_kb"
	case new_model.FileTypeMaintenanceRoutine:
		storageFieldName = "file_storage_routine_kb"
	default:
		return entity.BadRequest("Invalid file type")
	}
	queryStr, args, _ = sql_query.NewSQLUpdateBuilder(db.ClientTableName).
		Increment(map[string]any{
			"file_storage_used_kb":      data.FileSizeKB,
			"file_storage_available_kb": -1 * data.FileSizeKB,
			storageFieldName:            data.FileSizeKB,
		}).
		Where(map[string]sql_query.SQLCondition{
			"client_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    companyCode,
			},
		}).
		Build()

	_, err = clientService.UpdateOne(
		context.Background(),
		queryStr,
		args...,
	)
	if err != nil {
		return &entity.HttpError{Message: "Failed to update file storage"}
	}

	// ─────── Insert file ───────
	fixedAssetDBName := utildb.FixedAssetDBName.CompanyDBName(companyCode)
	fileService := serviceFactory.MakePostgreSqlService(
		fixedAssetDBName,
	)

	queryStr, args, _ = sql_query.NewSQLInsertBuilder(db.FilesInStorageTableName).
		Insert(data).
		Build()

	_, err = fileService.InsertOne(context.Background(), queryStr, args...)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}

func (f *FileManagerHelperSQLImpl) DeleteFileManagerSQL(
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
	filter map[string]sql_query.SQLCondition,
	fileType new_model.FileType,
) *entity.HttpError {
	fixedAssetDB := utildb.FixedAssetDBName.CompanyDBName(companyCode)
	fileService := serviceFactory.MakePostgreSqlService(fixedAssetDB)

	// ─────── Find the file ───────
	if len(filter) == 0 {
		return entity.BadRequest("Filter cannot be empty")
	}

	queryStr, args, _ := sql_query.NewSQLSelectBuilder[int](db.FilesInStorageTableName).
		Where(filter).
		SetLimit(1).
		Build()

	var data new_model.FileInStorage
	err := fileService.SelectOne(&data, context.Background(), queryStr, args...)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	// ─────── Delete the file ───────
	queryStr, args, _ = sql_query.NewSQLDeleteBuilder(db.FilesInStorageTableName).
		Delete("id").
		Where(filter).
		Build()
	_, err = fileService.DeleteOne(context.Background(), queryStr, args...)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	// ─────── Check & update storage ───────
	adminDb := utildb.AdminDBName
	clientService := serviceFactory.MakePostgreSqlService(adminDb)

	queryStr, args, _ = sql_query.NewSQLSelectBuilder[int](db.ClientTableName).
		Where(map[string]sql_query.SQLCondition{
			"client_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    companyCode,
			},
		}).
		SetLimit(1).
		Build()

	var client admin_model.Client
	err = clientService.SelectOne(&client, context.Background(), queryStr, args...)
	if err != nil {
		return &entity.HttpError{}
	}

	storageFieldName := ""
	switch fileType {
	case new_model.FileTypeAttachment:
		storageFieldName = "file_storage_attachment_kb"
	case new_model.FileTypeDamage:
		storageFieldName = "file_storage_damaged_kb"
	case new_model.FileTypeDisposal:
		storageFieldName = "file_storage_disposal_kb"
	case new_model.FileTypeRepairTicketing:
		storageFieldName = "file_storage_repair_ticketing_kb"
	case new_model.FileTypeMaintenanceRoutine:
		storageFieldName = "file_storage_routine_kb"
	default:
		return entity.BadRequest("Invalid file type")
	}
	queryStr, args, _ = sql_query.NewSQLUpdateBuilder(db.ClientTableName).
		Increment(map[string]any{
			"file_storage_used_kb":      -1 * data.FileSizeKB,
			"file_storage_available_kb": data.FileSizeKB,
			storageFieldName:            -1 * data.FileSizeKB,
		}).
		Where(map[string]sql_query.SQLCondition{
			"client_id": {
				Operator: sql_query.SQLOperatorEqual,
				Value:    companyCode,
			},
		}).
		Build()

	_, err = clientService.UpdateOne(
		context.Background(),
		queryStr,
		args...,
	)
	if err != nil {
		return entity.InternalServerError(err.Error())
	}

	return nil
}
