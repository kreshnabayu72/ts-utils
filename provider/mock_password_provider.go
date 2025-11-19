package provider

import "github.com/stretchr/testify/mock"

// MockPasswordProvider implements PasswordProviderer
//
// Used to mock the PasswordProvider implementation.
type MockPasswordProvider struct {
	mock.Mock
}

func (p *MockPasswordProvider) Init() {
	// Pass
}

func (p *MockPasswordProvider) HashPassword(password string) string {
	args := p.Called(password)
	return args.Get(0).(string)
}

func (p *MockPasswordProvider) ComparePassword(hashedPassword string, plainPassword string) bool {
	args := p.Called(hashedPassword, plainPassword)
	return args.Get(0).(bool)
}

func (p *MockPasswordProvider) Close() {
	// Pass
}
