package service

import (
	"github.com/stretchr/testify/mock"
)

// MockEmailService implements EmailServicer
//
// Used to mock the EmailService implementation.
type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) Init() {
	// Pass
}

func (m *MockEmailService) SendEmail(
	props SendEmailProps,
) error {
	args := m.Called(props)
	return args.Error(0)
}

func (m *MockEmailService) Close() {
	// Pass
}
