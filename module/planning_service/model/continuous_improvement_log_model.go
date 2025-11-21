package new_model

import (
	"time"
)

type ContinuousImprovementLog struct {
	ID int `json:"id" column:"continuous_improvement_logs.id"`

	ImprovementId   string `json:"improvementId" column:"continuous_improvement_logs.improvement_id"`
	Improvement     string `json:"improvement" column:"continuous_improvement_logs.improvement"`
	Category        string `json:"category" column:"continuous_improvement_logs.category"`
	ExpectedBenefit string `json:"expectedBenefit" column:"continuous_improvement_logs.expected_benefit"`
	PicId           int    `json:"picId" column:"continuous_improvement_logs.pic_id"`
	PicFullName     string `json:"picFullName" column:"continuous_improvement_logs.pic_full_name"`
	TargetDate      string `json:"targetDate" column:"continuous_improvement_logs.target_date"`
	Status          string `json:"status" column:"continuous_improvement_logs.status"`
	IsClosed        string `json:"isClosed" column:"continuous_improvement_logs.is_closed"`

	IsDeleted bool       `json:"isDeleted"        column:"continuous_improvement_logs.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"continuous_improvement_logs.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"continuous_improvement_logs.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"continuous_improvement_logs.updated_at"`
}
