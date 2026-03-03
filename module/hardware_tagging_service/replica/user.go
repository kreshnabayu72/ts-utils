package replica

import "time"

type User struct {
	ID             int        `json:"_id"            column:"users.id"`
	FirstName      string     `json:"firstName"      column:"users.first_name"`
	LastName       string     `json:"lastName"       column:"users.last_name"`
	FullName       string     `json:"fullName"       column:"users.full_name"`
	Email          string     `json:"email"          column:"users.email"`
	ProfilePicture *string    `json:"profilePicture" column:"users.profile_picture"`
	IsDeleted      bool       `json:"isDeleted"      column:"users.is_deleted"`
	DeletedAt      *time.Time `json:"deletedAt"      column:"users.deleted_at"`
}

func (r *User) GetTableName() string {
	return "users"
}

func (r *User) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":              "BIGINT",
		"first_name":      "VARCHAR",
		"last_name":       "VARCHAR",
		"full_name":       "VARCHAR",
		"email":           "VARCHAR",
		"profile_picture": "VARCHAR",
		"is_deleted":      "BOOLEAN",
		"deleted_at":      "TIMESTAMPTZ",
	}
}

func (r *User) GetColumns() []string {
	return []string{
		"id",
		"first_name",
		"last_name",
		"full_name",
		"email",
		"profile_picture",
		"is_deleted",
		"deleted_at",
	}
}
