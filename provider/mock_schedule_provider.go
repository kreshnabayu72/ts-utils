package provider

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockScheduleProvider struct {
	mock.Mock
}

func (m *MockScheduleProvider) Init() {}

func (m *MockScheduleProvider) Close() {}

func (m *MockScheduleProvider) CreateSchedule(ctx context.Context, options ScheduleOptions) (string, error) {
	args := m.Called(ctx, options)
	return args.String(0), args.Error(1)
}

func (m *MockScheduleProvider) UpdateSchedule(ctx context.Context, scheduleID string, options ScheduleOptions) error {
	args := m.Called(ctx, scheduleID, options)
	return args.Error(0)
}

func (m *MockScheduleProvider) DeleteSchedule(ctx context.Context, scheduleID string) error {
	args := m.Called(ctx, scheduleID)
	return args.Error(0)
}
