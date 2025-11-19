package provider

import "golang.org/x/crypto/bcrypt"

// PasswordProviderer provides functionality to hash and compare password.
type PasswordProviderer interface {
	Init()
	HashPassword(password string) string
	ComparePassword(hashedPassword, plainPassword string) bool
	Close()
}

// PasswordProvider implements PasswordProviderer
type PasswordProvider struct{}

func (p *PasswordProvider) Init() {
	// Pass
}

// HashPassword hashes the provided password using bcrypt.
func (p *PasswordProvider) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(hashed)
}

// ComparePassword compares the hashed password with the plain password using bcrypt.
func (p *PasswordProvider) ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (p *PasswordProvider) Close() {
	// Pass
}
