package model

import "time"

type TabletStock struct {
	ID                 int        `json:"_id" column:"tablet_stocks.id"`
	DeviceID           string     `json:"deviceId" column:"tablet_stocks.device_id"`
	DeviceNameID       int        `json:"deviceNameId" column:"tablet_stocks.device_name_id"`
	DeviceNameSequence int        `json:"deviceNameSequence" column:"tablet_stocks.device_name_sequence"`
	ClientID           *string    `json:"clientId" column:"tablet_stocks.client_id"`
	IsDeleted          bool       `json:"isDeleted" column:"tablet_stocks.is_deleted"`
	DeletedAt          *time.Time `json:"deletedAt" column:"tablet_stocks.deleted_at"`
	CreatedAt          time.Time  `json:"createdAt" column:"tablet_stocks.created_at"`
	UpdatedAt          time.Time  `json:"updatedAt" column:"tablet_stocks.updated_at"`
}

