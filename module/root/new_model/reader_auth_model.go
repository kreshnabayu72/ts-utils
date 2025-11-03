package new_model

import (
	"time"
)

type ReaderAuth struct {
	ID          int        `json:"_id"         column:"reader_auths.id"`
	CompanyCode string     `json:"companyCode" column:"reader_auths.company_code"`
	ReaderId    string     `json:"readerId" column:"reader_auths.reader_id"`
	CreatedAt   *time.Time `json:"createdAt"   column:"reader_auths.created_at"`
	UpdatedAt   *time.Time `json:"updatedAt"   column:"reader_auths.updated_at"`
}
