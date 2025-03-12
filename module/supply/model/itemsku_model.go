package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

// TransactionalPolicy represents the transactional policy of individual item sku
type TransactionalPolicy struct {
	AdminOnlyCheckout bool     `json:"adminOnlyCheckout" bson:"adminOnlyCheckout"`
	Returnable        bool     `json:"returnable"        bson:"returnable"`
	Exchangeable      bool     `json:"exchangeable"      bson:"exchangeable"`
	ReturnReason      []string `json:"returnReason"      bson:"returnReason"`
}

// SafetyStockSettings represents possible safety stock settings values
type SafetyStockSettings string

const (
	SafetyStockSettingsGlobally       SafetyStockSettings = "Set globally"
	SafetyStockSettingsEachAdminGroup SafetyStockSettings = "Set for each administrative group"
	SafetyStockSettingsEachGroup      SafetyStockSettings = "Set for each group"
)

// ItemSKU represents the schema for a single item sku data
//
// ItemSKU are meant to group stocks under the same settings set in the sku for the stock
type ItemSKU struct {
	ID                            primitive.ObjectID  `json:"_id,omitempty"                 bson:"_id,omitempty"`
	Key                           int                 `json:"key"                           bson:"key"`
	ItemName                      primitive.ObjectID  `json:"itemName"                      bson:"itemName"`
	Category                      primitive.ObjectID  `json:"category"                      bson:"category"`
	Brand                         primitive.ObjectID  `json:"brand"                         bson:"brand"`
	Model                         primitive.ObjectID  `json:"model"                         bson:"model"`
	SKU                           string              `json:"sku"                           bson:"sku"`
	ItemNameWithSKU               string              `json:"itemNameWithSKU"               bson:"itemNameWithSKU"`
	SKUAlias                      *string             `json:"skuAlias"                      bson:"skuAlias"`
	MinimumStock                  int                 `json:"minimumStock"                  bson:"minimumStock"`
	SafetyStock                   *int                `json:"safetyStock"                   bson:"safetyStock"`
	MinimumAndSafetyStockSettings SafetyStockSettings `json:"minimumAndSafetyStockSettings" bson:"minimumAndSafetyStockSettings"`
	TagType                       model.TagType       `json:"tagType"                       bson:"tagType"`
	HasProductionDate             bool                `json:"hasProductionDate"             bson:"hasProductionDate"`
	HasExpiredDate                bool                `json:"hasExpiredDate"                bson:"hasExpiredDate"`
	HasBarcode                    bool                `json:"hasBarcode"                    bson:"hasBarcode"`
	Barcode                       *string             `json:"barcode"                       bson:"barcode"`
	DefaultImage                  *string             `json:"defaultImage"                  bson:"defaultImage"` // Possible Values: "firstImage", "secondImage"
	FirstImage                    *string             `json:"firstImage"                    bson:"firstImage"`
	SecondImage                   *string             `json:"secondImage"                   bson:"secondImage"`
	CustomFields                  []CustomField       `json:"customFields"                  bson:"customFields"`
	TransactionalPolicy           TransactionalPolicy `json:"transactionalPolicy"           bson:"transactionalPolicy"`
	TotalStock                    int                 `json:"totalStock"                    bson:"totalStock"`
	UnplacedStock                 int                 `json:"unplacedStock"                 bson:"unplacedStock"`
	AvailableStock                int                 `json:"availableStock"                bson:"availableStock"`
	HasBeenInitialized            bool                `json:"hasBeenInitialized"            bson:"hasBeenInitialized"`
	ModifiedBy                    primitive.ObjectID  `json:"modifiedBy"                    bson:"modifiedBy"`
	IsDeleted                     bool                `json:"isDeleted"                     bson:"isDeleted"`
	CreatedAt                     time.Time           `json:"createdAt"                     bson:"createdAt"`
	UpdatedAt                     time.Time           `json:"updatedAt"                     bson:"updatedAt"`
}
