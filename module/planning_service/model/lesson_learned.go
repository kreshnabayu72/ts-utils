package new_model

import (
	"time"
)

type LessonLearned struct {
	ID int `json:"id" column:"lesson_learned.id"`

	LessonId string `json:"lessonId" column:"lesson_learned.lesson_id"`
	Lesson   string `json:"lesson" column:"lesson_learned.lesson"`
	Category string `json:"category" column:"lesson_learned.category"`

	RootCause       string `json:"rootCause" column:"lesson_learned.root_cause"`
	ExpectedBenefit string `json:"expectedBenefit" column:"lesson_learned.expected_benefit"`
	ActionTaken     string `json:"actionTaken" column:"lesson_learned.action_taken"`

	Status string `json:"status" column:"lesson_learned.status"`

	IsDeleted bool       `json:"isDeleted"        column:"lesson_learned.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"lesson_learned.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"lesson_learned.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"lesson_learned.updated_at"`
}
