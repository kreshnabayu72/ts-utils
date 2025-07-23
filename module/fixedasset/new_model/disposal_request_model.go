package new_model

import "time"

type DisposalRequest struct {
	ID                      int                   `json:"id" column:"id"`
	AssetID                 int                   `json:"assetId" column:"asset_id"`
	TransactionID           int                   `json:"transactionId" column:"transaction_id"`
	ReportId                *int                  `json:"reportId" column:"report_id"`
	PreviousStatus          AssetStatus           `json:"previousStatus" column:"previous_status"`
	Status                  DisposalRequestStatus `json:"status" column:"status"`
	SalvageableParts        []string              `json:"salvageableParts" column:"salvageable_parts"`
	Notes                   string                `json:"notes" column:"notes"`
	CompletionApprovalRound string                `json:"completionApprovalRound" column:"completion_approval_round"`
	IsVerified              bool                  `json:"isVerified" column:"is_verified"`
	CreatedAt               time.Time             `json:"createdAt" column:"created_at"`
	UpdatedAt               time.Time             `json:"updatedAt" column:"updated_at"`
}
