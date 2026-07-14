package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type Asset struct {
	ID               int              `json:"_id"             column:"assets.id"`
	NameWithSequence string           `json:"nameWithSequence" column:"assets.name_with_sequence"`
	Status           enum.AssetStatus `json:"status"           column:"assets.status"`
	IsDeleted        bool             `json:"isDeleted"        column:"assets.is_deleted"`
	DeletedAt        *time.Time       `json:"deletedAt"        column:"assets.deleted_at"`
	CreatedAt        time.Time        `json:"createdAt"        column:"assets.created_at"`
	UpdatedAt        time.Time        `json:"updatedAt"        column:"assets.updated_at"`
}

