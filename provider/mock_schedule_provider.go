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

func (m *MockScheduleProvider) CreateSchedule(ctx context.Context, scheduleIdentifier string, options ScheduleOptions) (string, error) {
	args := m.Called(ctx, scheduleIdentifier, options)
	return args.String(0), args.Error(1)
}

func (m *MockScheduleProvider) CreateOrUpdateSchedule(ctx context.Context, scheduleIdentifier string, options ScheduleOptions) (string, error) {
	args := m.Called(ctx, scheduleIdentifier, options)
	return args.String(0), args.Error(1)
}

func (m *MockScheduleProvider) UpdateSchedule(ctx context.Context, scheduleIdentifier string, options ScheduleOptions) error {
	args := m.Called(ctx, scheduleIdentifier, options)
	return args.Error(0)
}

func (m *MockScheduleProvider) DeleteSchedule(ctx context.Context, scheduleIdentifier string) error {
	args := m.Called(ctx, scheduleIdentifier)
	return args.Error(0)
}
