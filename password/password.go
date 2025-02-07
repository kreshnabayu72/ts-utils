package password

import "golang.org/x/crypto/bcrypt"

type PasswordUtilImpl struct{}

func (p *PasswordUtilImpl) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(hashed)
}

func (p *PasswordUtilImpl) ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
