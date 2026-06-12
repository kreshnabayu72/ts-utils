package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/client"
)

type MockWorkflowRun struct {
	mock.Mock
}

func (m *MockWorkflowRun) SetResult(v any) {
	
}

func (m *MockWorkflowRun) Get(ctx context.Context, valuePtr interface{}) error {
	args := m.Called(ctx, valuePtr)
	return args.Error(0)
}

func (m *MockWorkflowRun) GetWithOptions(ctx context.Context, valuePtr interface{}, options client.WorkflowRunGetOptions) error {
	args := m.Called(ctx, valuePtr, options)
	return args.Error(0)
}

func (m *MockWorkflowRun) GetID() string {
	return "mock-workflow-id"
}

func (m *MockWorkflowRun) GetRunID() string {
	return "mock-run-id"
}