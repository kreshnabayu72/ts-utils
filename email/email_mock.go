package email

import "github.com/stretchr/testify/mock"

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) SendEmail(
	props SendEmailProps,
) error {
	args := m.Called(props)
	return args.Error(0)
}
