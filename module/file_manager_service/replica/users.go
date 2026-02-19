package replica

import "time"

type User struct {
	ID             int       `json:"_id"            column:"users.id"              sqlType:"BIGINT"`
	FirstName      string    `json:"firstName"      column:"users.first_name"      sqlType:"VARCHAR"`
	LastName       string    `json:"lastName"       column:"users.last_name"       sqlType:"VARCHAR"`
	FullName       string    `json:"fullName"       column:"users.full_name"       sqlType:"VARCHAR"`
	Email          string    `json:"email"          column:"users.email"           sqlType:"VARCHAR"`
	ProfilePicture string    `json:"profilePicture" column:"users.profile_picture" sqlType:"VARCHAR"`
	IsDeleted      bool      `json:"isDeleted"      column:"users.is_deleted"      sqlType:"BOOLEAN"`
	DeletedAt      time.Time `json:"deletedAt"      column:"users.deleted_at"      sqlType:"TIMESTAMPTZ"`
}

func (u *User) GetTableName() string {
	return "users"
}

func (u *User) GetTableNameWithSchema() string {
	return "public.users"
}

func (u *User) GetColumnTypes() map[string]string {
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

func (u *User) GetColumns() []string {
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
