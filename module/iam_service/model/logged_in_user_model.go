package model

type LoggedInUser struct {
	ID        int    `json:"_id"       column:"users.id"`
	Email     string `json:"email"     column:"users.email"`
	FirstName string `json:"firstName" column:"users.first_name"`
	LastName  string `json:"lastName"  column:"users.last_name"`
	FullName  string `json:"fullName"  column:"users.full_name"`
}
