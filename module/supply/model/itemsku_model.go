package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type TransactionalPolicy struct {
	AdminOnlyCheckout bool     `json:"adminOnlyCheckout" bson:"adminOnlyCheckout"`
	Returnable        bool     `json:"returnable"        bson:"returnable"`
	Exchangeable      bool     `json:"exchangeable"      bson:"exchangeable"`
	ReturnReason      []string `json:"returnReason"      bson:"returnReason"`
}

type ItemSKU struct {
	ID                            primitive.ObjectID   `json:"_id,omitempty"                 bson:"_id,omitempty"`
	Key                           int                  `json:"key"                           bson:"key"`
	ItemName                      primitive.ObjectID   `json:"itemName"                      bson:"itemName"`
	Category                      primitive.ObjectID   `json:"category"                      bson:"category"`
	Brand                         []primitive.ObjectID `json:"brand"                         bson:"brand"`
	Model                         []primitive.ObjectID `json:"model"                         bson:"model"`
	SKU                           string               `json:"sku"                           bson:"sku"`
	SKUAlias                      *string              `json:"skuAlias"                      bson:"skuAlias"`
	MinimumStock                  int                  `json:"minimumStock"                  bson:"minimumStock"`
	SafetyStock                   *int                 `json:"safetyStock"                   bson:"safetyStock"`
	MinimumAndSafetyStockSettings string               `json:"minimumAndSafetyStockSettings" bson:"minimumAndSafetyStockSettings"`
	TagType                       model.TagType        `json:"tagType"                       bson:"tagType"`
	HasProductionDate             bool                 `json:"hasProductionDate"             bson:"hasProductionDate"`
	HasExpiredDate                bool                 `json:"hasExpiredDate"                bson:"hasExpiredDate"`
	HasBarcode                    bool                 `json:"hasBarcode"                    bson:"hasBarcode"`
	Barcode                       *string              `json:"barcode"                       bson:"barcode"`
	Image                         *string              `json:"image"                         bson:"image"`
	CustomFields                  []CustomField        `json:"customFields"                  bson:"customFields"`
	TransactionalPolicy           TransactionalPolicy  `json:"transactionalPolicy"           bson:"transactionalPolicy"`
	ModifiedBy                    primitive.ObjectID   `json:"modifiedBy"                    bson:"modifiedBy"`
	CreatedAt                     time.Time            `json:"createdAt"                     bson:"createdAt"`
	UpdatedAt                     time.Time            `json:"updatedAt"                     bson:"updatedAt"`
}
