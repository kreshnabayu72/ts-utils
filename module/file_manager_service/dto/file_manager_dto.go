package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/file_manager_service/model"

type CreateFileManagerDataActivityParam struct {
	Data     model.FileManagerData
	ClientId string
}

type SaveFileActivityParam struct {
	ClientId   string
	FolderName string
	File       FileUploadData
}

type FileUploadData struct {
	Filename    string
	ContentType string
	Bytes       []byte
}
