package mocks

import (
	"context"

	ems_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/enterprise_management_service/model"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/service"
	"github.com/stretchr/testify/mock"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/operatorservice/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
)

type MockTemporalClient struct {
	mock.Mock
}

func (m *MockTemporalClient) ExecuteWorkflow(ctx context.Context, options client.StartWorkflowOptions, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
	callArgs := m.Called(ctx, options, workflow, args)
	if wr := callArgs.Get(0); wr != nil {
		return wr.(client.WorkflowRun), callArgs.Error(1)
	}
	return nil, callArgs.Error(1)
}

func (m *MockTemporalClient) GetWorkflow(ctx context.Context, workflowID string, runID string) client.WorkflowRun {
	args := m.Called(ctx, workflowID, runID)
	return args.Get(0).(client.WorkflowRun)
}

func (m *MockTemporalClient) SignalWorkflow(ctx context.Context, workflowID string, runID string, signalName string, arg interface{}) error {
	args := m.Called(ctx, workflowID, runID, signalName, arg)
	return args.Error(0)
}

func (m *MockTemporalClient) SignalWithStartWorkflow(ctx context.Context, workflowID string, signalName string, signalArg interface{}, options client.StartWorkflowOptions, workflow interface{}, workflowArgs ...interface{}) (client.WorkflowRun, error) {
	args := m.Called(ctx, workflowID, signalName, signalArg, options, workflow, workflowArgs)
	if wr := args.Get(0); wr != nil {
		return wr.(client.WorkflowRun), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTemporalClient) NewWithStartWorkflowOperation(options client.StartWorkflowOptions, workflow interface{}, args ...interface{}) client.WithStartWorkflowOperation {
	callArgs := m.Called(options, workflow, args)
	return callArgs.Get(0).(client.WithStartWorkflowOperation)
}

func (m *MockTemporalClient) CancelWorkflow(ctx context.Context, workflowID string, runID string) error {
	args := m.Called(ctx, workflowID, runID)
	return args.Error(0)
}

func (m *MockTemporalClient) TerminateWorkflow(ctx context.Context, workflowID string, runID string, reason string, details ...interface{}) error {
	args := m.Called(ctx, workflowID, runID, reason, details)
	return args.Error(0)
}

func (m *MockTemporalClient) GetWorkflowHistory(ctx context.Context, workflowID string, runID string, isLongPoll bool, filterType enumspb.HistoryEventFilterType) client.HistoryEventIterator {
	args := m.Called(ctx, workflowID, runID, isLongPoll, filterType)
	return args.Get(0).(client.HistoryEventIterator)
}

func (m *MockTemporalClient) CompleteActivity(ctx context.Context, taskToken []byte, result interface{}, err error) error {
	args := m.Called(ctx, taskToken, result, err)
	return args.Error(0)
}

func (m *MockTemporalClient) CompleteActivityByID(ctx context.Context, namespace, workflowID, runID, activityID string, result interface{}, err error) error {
	args := m.Called(ctx, namespace, workflowID, runID, activityID, result, err)
	return args.Error(0)
}

func (m *MockTemporalClient) RecordActivityHeartbeat(ctx context.Context, taskToken []byte, details ...interface{}) error {
	args := m.Called(ctx, taskToken, details)
	return args.Error(0)
}

func (m *MockTemporalClient) RecordActivityHeartbeatByID(ctx context.Context, namespace, workflowID, runID, activityID string, details ...interface{}) error {
	args := m.Called(ctx, namespace, workflowID, runID, activityID, details)
	return args.Error(0)
}

func (m *MockTemporalClient) ListClosedWorkflow(ctx context.Context, request *workflowservice.ListClosedWorkflowExecutionsRequest) (*workflowservice.ListClosedWorkflowExecutionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.ListClosedWorkflowExecutionsResponse), args.Error(1)
}

func (m *MockTemporalClient) ListOpenWorkflow(ctx context.Context, request *workflowservice.ListOpenWorkflowExecutionsRequest) (*workflowservice.ListOpenWorkflowExecutionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.ListOpenWorkflowExecutionsResponse), args.Error(1)
}

func (m *MockTemporalClient) ListWorkflow(ctx context.Context, request *workflowservice.ListWorkflowExecutionsRequest) (*workflowservice.ListWorkflowExecutionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.ListWorkflowExecutionsResponse), args.Error(1)
}

func (m *MockTemporalClient) ListArchivedWorkflow(ctx context.Context, request *workflowservice.ListArchivedWorkflowExecutionsRequest) (*workflowservice.ListArchivedWorkflowExecutionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.ListArchivedWorkflowExecutionsResponse), args.Error(1)
}

func (m *MockTemporalClient) ScanWorkflow(ctx context.Context, request *workflowservice.ScanWorkflowExecutionsRequest) (*workflowservice.ScanWorkflowExecutionsResponse, error) { //nolint:staticcheck
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.ScanWorkflowExecutionsResponse), args.Error(1)
}

func (m *MockTemporalClient) CountWorkflow(ctx context.Context, request *workflowservice.CountWorkflowExecutionsRequest) (*workflowservice.CountWorkflowExecutionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.CountWorkflowExecutionsResponse), args.Error(1)
}

func (m *MockTemporalClient) GetSearchAttributes(ctx context.Context) (*workflowservice.GetSearchAttributesResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).(*workflowservice.GetSearchAttributesResponse), args.Error(1)
}

func (m *MockTemporalClient) QueryWorkflow(ctx context.Context, workflowID string, runID string, queryType string, args ...interface{}) (converter.EncodedValue, error) {
	callArgs := m.Called(ctx, workflowID, runID, queryType, args)
	return callArgs.Get(0).(converter.EncodedValue), callArgs.Error(1)
}

func (m *MockTemporalClient) QueryWorkflowWithOptions(ctx context.Context, request *client.QueryWorkflowWithOptionsRequest) (*client.QueryWorkflowWithOptionsResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*client.QueryWorkflowWithOptionsResponse), args.Error(1)
}

func (m *MockTemporalClient) DescribeWorkflowExecution(ctx context.Context, workflowID, runID string) (*workflowservice.DescribeWorkflowExecutionResponse, error) {
	args := m.Called(ctx, workflowID, runID)
	return args.Get(0).(*workflowservice.DescribeWorkflowExecutionResponse), args.Error(1)
}

func (m *MockTemporalClient) DescribeWorkflow(ctx context.Context, workflowID, runID string) (*client.WorkflowExecutionDescription, error) {
	args := m.Called(ctx, workflowID, runID)
	return args.Get(0).(*client.WorkflowExecutionDescription), args.Error(1)
}

func (m *MockTemporalClient) DescribeTaskQueue(ctx context.Context, taskqueue string, taskqueueType enumspb.TaskQueueType) (*workflowservice.DescribeTaskQueueResponse, error) {
	args := m.Called(ctx, taskqueue, taskqueueType)
	return args.Get(0).(*workflowservice.DescribeTaskQueueResponse), args.Error(1)
}

func (m *MockTemporalClient) DescribeTaskQueueEnhanced(ctx context.Context, options client.DescribeTaskQueueEnhancedOptions) (client.TaskQueueDescription, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(client.TaskQueueDescription), args.Error(1)
}

func (m *MockTemporalClient) ResetWorkflowExecution(ctx context.Context, request *workflowservice.ResetWorkflowExecutionRequest) (*workflowservice.ResetWorkflowExecutionResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*workflowservice.ResetWorkflowExecutionResponse), args.Error(1)
}

func (m *MockTemporalClient) UpdateWorkerBuildIdCompatibility(ctx context.Context, options *client.UpdateWorkerBuildIdCompatibilityOptions) error {
	args := m.Called(ctx, options)
	return args.Error(0)
}

func (m *MockTemporalClient) GetWorkerBuildIdCompatibility(ctx context.Context, options *client.GetWorkerBuildIdCompatibilityOptions) (*client.WorkerBuildIDVersionSets, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(*client.WorkerBuildIDVersionSets), args.Error(1)
}

func (m *MockTemporalClient) GetWorkerTaskReachability(ctx context.Context, options *client.GetWorkerTaskReachabilityOptions) (*client.WorkerTaskReachability, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(*client.WorkerTaskReachability), args.Error(1)
}

func (m *MockTemporalClient) UpdateWorkerVersioningRules(ctx context.Context, options client.UpdateWorkerVersioningRulesOptions) (*client.WorkerVersioningRules, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(*client.WorkerVersioningRules), args.Error(1)
}

func (m *MockTemporalClient) GetWorkerVersioningRules(ctx context.Context, options client.GetWorkerVersioningOptions) (*client.WorkerVersioningRules, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(*client.WorkerVersioningRules), args.Error(1)
}

func (m *MockTemporalClient) CheckHealth(ctx context.Context, request *client.CheckHealthRequest) (*client.CheckHealthResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*client.CheckHealthResponse), args.Error(1)
}

func (m *MockTemporalClient) UpdateWorkflow(ctx context.Context, options client.UpdateWorkflowOptions) (client.WorkflowUpdateHandle, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(client.WorkflowUpdateHandle), args.Error(1)
}

func (m *MockTemporalClient) UpdateWorkflowExecutionOptions(ctx context.Context, options client.UpdateWorkflowExecutionOptionsRequest) (client.WorkflowExecutionOptions, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(client.WorkflowExecutionOptions), args.Error(1)
}

func (m *MockTemporalClient) UpdateWithStartWorkflow(ctx context.Context, options client.UpdateWithStartWorkflowOptions) (client.WorkflowUpdateHandle, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(client.WorkflowUpdateHandle), args.Error(1)
}

func (m *MockTemporalClient) GetWorkflowUpdateHandle(ref client.GetWorkflowUpdateHandleOptions) client.WorkflowUpdateHandle {
	args := m.Called(ref)
	return args.Get(0).(client.WorkflowUpdateHandle)
}

func (m *MockTemporalClient) WorkflowService() workflowservice.WorkflowServiceClient {
	args := m.Called()
	return args.Get(0).(workflowservice.WorkflowServiceClient)
}

func (m *MockTemporalClient) OperatorService() operatorservice.OperatorServiceClient {
	args := m.Called()
	return args.Get(0).(operatorservice.OperatorServiceClient)
}

func (m *MockTemporalClient) ScheduleClient() client.ScheduleClient {
	args := m.Called()
	return args.Get(0).(client.ScheduleClient)
}

func (m *MockTemporalClient) DeploymentClient() client.DeploymentClient {
	args := m.Called()
	return args.Get(0).(client.DeploymentClient)
}

func (m *MockTemporalClient) WorkerDeploymentClient() client.WorkerDeploymentClient {
	args := m.Called()
	return args.Get(0).(client.WorkerDeploymentClient)
}

func (m *MockTemporalClient) Close() {
	m.Called()
}

func CreateTemporalService(mockClients []ems_model.Client, executeErr, getErr error) *service.TemporalService {
	mockClient := new(MockTemporalClient)
	mockRun := new(MockWorkflowRun)

	if executeErr != nil {
		mockClient.On("ExecuteWorkflow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, executeErr)
	} else {
		mockClient.On("ExecuteWorkflow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(mockRun, nil)

		if getErr != nil {
			mockRun.On("Get", mock.Anything, mock.Anything).
				Return(getErr)
		} else {
			mockRun.On("Get", mock.Anything, mock.Anything).
				Run(func(args mock.Arguments) {
					if ptr, ok := args.Get(1).(*[]ems_model.Client); ok {
						*ptr = mockClients
					}
				}).
				Return(nil)
		}
	}

	return &service.TemporalService{Client: mockClient}
}
