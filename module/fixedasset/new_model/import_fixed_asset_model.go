package new_model

import "time"

type ImportFixedAsset struct {
	ID  int `json:"id"  column:"id"`
	Key int `json:"key" column:"key"`

	Group       string `json:"group"       column:"group"`
	Category    string `json:"category"    column:"category"`
	Name        string `json:"name"        column:"name"`
	Quantity    int    `json:"quantity"    column:"quantity"`
	Brand       string `json:"brand"       column:"brand"`
	Model       string `json:"model"       column:"model"`
	Measurement string `json:"measurement" column:"measurement"`

	// ─────── Standard & Enterprise plan ───────
	GroupCode          string     `json:"groupCode"           column:"group_code"`
	CategoryCode       string     `json:"categoryCode"        column:"category_code"`
	NameCode           string     `json:"nameCode"            column:"name_code"`
	Currency           string     `json:"currency"             column:"currency"`
	PurchaseCost       int        `json:"purchaseCost"        column:"purchase_cost"`
	TangibleAssetGroup string     `json:"tangibleAssetGroup" column:"tangible_asset_group"`
	UsefulLife         int        `json:"usefulLife"          column:"useful_life"`
	PurchaseDate       *time.Time `json:"purchaseDate"        column:"purchase_date"`
	AliasName          string     `json:"aliasName"           column:"alias_name"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}
