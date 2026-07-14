package model

import "time"

type Tablet struct {
	ID            int        `json:"_id" column:"tablets.id"`
	ReaderId      *int       `json:"readerId" column:"tablets.reader_id"`
	AliasName     string     `json:"aliasName" column:"tablets.alias_name"`
	Status        string     `json:"status" column:"tablets.status"`
	TabletStockId int        `json:"tabletStockId" column:"tablets.tablet_stock_id"`
	ReportedById  *int       `json:"reportedById" column:"tablets.reported_by_id"`
	ReportedAt    *time.Time `json:"reportedAt" column:"tablets.reported_at"`
	IsDeleted     bool       `json:"isDeleted" column:"tablets.is_deleted"`
	DeletedAt     *time.Time `json:"deletedAt" column:"tablets.deleted_at"`
	CreatedAt     time.Time  `json:"createdAt" column:"tablets.created_at"`
	UpdatedAt     time.Time  `json:"updatedAt" column:"tablets.updated_at"`
}

