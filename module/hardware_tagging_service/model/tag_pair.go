package model

import (
	"time"

	util_enum "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type TagPair struct {
	ID int `json:"id" column:"tag_pairs.id"`

	// Nullable Foreign Keys
	RfidID *int `json:"rfidId" column:"tag_pairs.rfid_id"`
	NfcID  *int `json:"nfcId" column:"tag_pairs.nfc_id"`
	QrID   *int `json:"qrId" column:"tag_pairs.qr_id"`

	TagType util_enum.TagType `json:"tagType" column:"tag_pairs.tag_type"`

	CreatedAt time.Time `json:"createdAt" column:"tag_pairs.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"tag_pairs.updated_at"`
}
func (r *TagPair) GetColumns() []string {
	return []string{
		"id",
		"rfid_id",
		"nfc_id",
		"qr_id",
		"tag_type",
		"created_at",
		"updated_at",
	}
}
