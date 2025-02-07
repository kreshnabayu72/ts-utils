package password

import "github.com/stretchr/testify/mock"

type PasswordUtilMock struct {
	mock.Mock
}

func (p *PasswordUtilMock) HashPassword(password string) string {
	args := p.Called(password)
	return args.Get(0).(string)
}

func (p *PasswordUtilMock) ComparePassword(hashedPassword, plainPassword string) bool {
	args := p.Called(hashedPassword, plainPassword)
	return args.Get(0).(bool)
}
