package replica

type FixedAssetDisposalTransaction struct {
	ID                 int    `json:"_id"                 column:"fixed_asset_disposal_transactions.id"`
	TransactionID      string `json:"transactionId"       column:"fixed_asset_disposal_transactions.transaction_id"`
	DocumentationPhoto string `json:"documentationPhoto"  column:"fixed_asset_disposal_transactions.documentation_photo"`
}

func (f *FixedAssetDisposalTransaction) GetTableName() string {
	return "fixed_asset_disposal_transactions"
}

func (f *FixedAssetDisposalTransaction) GetTableNameWithSchema() string {
	return "public.fixed_asset_disposal_transactions"
}

func (f *FixedAssetDisposalTransaction) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                  "BIGINT",
		"transaction_id":      "VARCHAR",
		"documentation_photo": "VARCHAR",
	}
}

func (f *FixedAssetDisposalTransaction) GetColumns() []string {
	return []string{
		"id",
		"transaction_id",
		"documentation_photo",
	}
}
