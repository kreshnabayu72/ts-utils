package password

type PasswordUtil interface {
	HashPassword(password string) string
	ComparePassword(hashedPassword, plainPassword string) bool
}
