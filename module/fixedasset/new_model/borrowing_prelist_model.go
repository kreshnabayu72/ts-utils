package new_model

import "time"

type BorrowingPrelist struct {
	ID         int       `json:"id" column:"id"`
	GroupId    int       `json:"groupId" column:"group_id"`
	BorrowerId int       `json:"borrowerId" column:"borrower_id"`
	CreatedAt  time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" column:"updated_at"`
}
