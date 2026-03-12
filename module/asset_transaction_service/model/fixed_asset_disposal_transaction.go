package model

type FixedAssetDisposalTransaction struct {
	ID                 int    `json:"_id"                 column:"fixed_asset_disposal_transactions.id"`
	TransactionID      string `json:"transactionId"       column:"fixed_asset_disposal_transactions.transaction_id"`
	DocumentationPhoto string `json:"documentationPhoto"  column:"fixed_asset_disposal_transactions.documentation_photo"`
}

func (f *FixedAssetDisposalTransaction) GetColumns() []string {
	return []string{
		"id",
		"transaction_id",
		"documentation_photo",
	}
}
