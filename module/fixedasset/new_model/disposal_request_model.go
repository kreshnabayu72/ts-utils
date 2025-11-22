package new_model

import (
	"time"
)

// DisposalRequest represents the disposal_requests table in the database.
type DisposalRequest struct {
	ID                        int                   `json:"id"                        column:"disposal_requests.id"`
	Asset                     int                   `json:"asset"                     column:"disposal_requests.asset"`
	Transaction               int                   `json:"transaction"               column:"disposal_requests.transaction"`
	ReportID                  *int                  `json:"reportId"                  column:"disposal_requests.report_id"`
	PreviousStatus            AssetStatus           `json:"previousStatus"            column:"disposal_requests.previous_status"`
	Status                    DisposalRequestStatus `json:"status"                    column:"disposal_requests.status"`
	SalvageableParts          []string              `json:"salvageableParts"          column:"disposal_requests.salvageable_parts"`
	SalePrice                 *string               `json:"salePrice"                 column:"disposal_requests.sale_price"`
	Notes                     string                `json:"notes"                     column:"disposal_requests.notes"`
	CompletionApprovalRound   *int                  `json:"completionApprovalRound"   column:"disposal_requests.completion_approval_round"`
	IsVerified                bool                  `json:"isVerified"                column:"disposal_requests.is_verified"`
	DocumentationPhoto        *string               `json:"documentationPhoto"        column:"disposal_requests.documentation_photo"`
	DocumentationPhotoAddress *string               `json:"documentationPhotoAddress" column:"disposal_requests.documentation_photo_address"`
	ReaderID                  *int                  `json:"readerId"                  column:"disposal_requests.reader_id"`
	CreatedAt                 time.Time             `json:"createdAt"                 column:"disposal_requests.created_at"`
	UpdatedAt                 time.Time             `json:"updatedAt"                 column:"disposal_requests.updated_at"`
}
