package service

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockTemporalService implements TemporalServicer
//
// Used to mock the TemporalService implementation.
type MockTemporalService struct {
	mock.Mock
}

func (t *MockTemporalService) Init() {
	// Pass
}

func (t *MockTemporalService) RegisterWorkflow(w any) {
	t.Called(w)
}

func (t *MockTemporalService) RegisterActivity(a any) {
	t.Called(a)
}

func (t *MockTemporalService) Run(ctx context.Context, interruptCh <-chan any) error {
	arg := t.Called(ctx, interruptCh)
	return arg.Error(0)
}

func (t *MockTemporalService) getHostPort() string {
	return "localhost:7233"
}

func (t *MockTemporalService) Close() {
	// Pass
}
