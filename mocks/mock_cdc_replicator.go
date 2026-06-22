package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockCDCReplicator struct {
	mock.Mock	
}

func (m *MockCDCReplicator) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCDCReplicator) Stop() {
	m.Called()
}