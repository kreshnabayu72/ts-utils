package temporal_helper

import (
	"errors"
	"fmt"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/entity"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/helper"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/service"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

type ExecuteWorkflowConfig struct {
	// name of the workflow to execute
	WorkflowName string
	// message to return on successful completion
	SuccessMessage string
	// HTTP status code for successful response
	SuccessCode int
	// Temporal service instance with client and task queue
	TemporalService *service.TemporalService
	// input parameters to pass to the workflow
	Param any
	// if true, includes error data in response when workflow fails (default: false)
	ErrorHasData bool
}

// ExecuteWorkflowWithDefaults executes a Temporal workflow with standardized configuration and error handling.
// This helper reduces boilerplate in controllers by centralizing workflow execution options, retry policies,
// and response formatting.
//
// Parameters:
//   - ctx: *fiber.Ctx – the current Fiber context
//   - config: ExecuteWorkflowConfig – configuration for workflow execution including:
//   - WorkflowName
//   - SuccessMessage
//   - SuccessCode
//   - TemporalService
//   - Param
//   - ErrorHasData
//   - result: T – pointer to receive the workflow result
//   - customOptions: ...client.StartWorkflowOptions (optional) – override default workflow options if provided
//
// Behavior:
//   - Configures retry policy with MaximumAttempts: 2 and non-retryable WORKFLOW_PANIC errors
//   - Automatically handles errors and converts them to HTTP responses via SendResponse
//   - Parses workflow errors using entity.ParseWorkflowError for user-friendly error messages
//   - When ErrorHasData is true, includes parsed error data (e.g., validation details) in the error response
//   - If customOptions is provided, it completely replaces the default workflow options
//
// Example:
//
//	func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
//	    var req model.CreateUserRequest
//	    if err := ctx.BodyParser(&req); err != nil {
//	        return SendResponse(ctx, fiber.StatusBadRequest, nil, "Invalid request", err)
//	    }
//
//	    config := temporal_helper.ExecuteWorkflowConfig{
//	        WorkflowName:    "CreateUserWorkflow",
//	        SuccessMessage:  "User created successfully",
//	        SuccessCode:     fiber.StatusCreated,
//	        TemporalService: c.TemporalService,
//	        Param:           req,
//	        ErrorHasData:    true,
//	    }
//
//	    var result model.CreateUserResponse
//	    return temporal_helper.ExecuteWorkflowWithDefaults(ctx, config, &result)
//	}
func ExecuteWorkflowWithDefaults[T any](
	ctx *fiber.Ctx,
	config ExecuteWorkflowConfig,
	result T,
	customOptions ...client.StartWorkflowOptions,
) error {
	wo := client.StartWorkflowOptions{
		TaskQueue: config.TemporalService.TaskQueue,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 2,
			NonRetryableErrorTypes: []string{
				"WORKFLOW_PANIC",
			},
		},
	}

	if len(customOptions) > 0 {
		wo = customOptions[0]
	}

	wr, executionErr := config.TemporalService.Client.ExecuteWorkflow(
		ctx.UserContext(),
		wo,
		config.WorkflowName,
		config.Param,
	)
	if executionErr != nil {
		e := entity.ToHttpError(executionErr)
		return helper.SendResponse(ctx, e.Code, nil, e.Message, nil)
	}

	if workflowErr := wr.Get(ctx.UserContext(), result); workflowErr != nil {
		e := entity.ParseWorkflowError(config.WorkflowName, workflowErr)
		if config.ErrorHasData {
			return helper.SendResponse(ctx, e.Code, e.Data, e.Message, nil)
		}
		return helper.SendResponse(ctx, e.Code, nil, e.Message, nil)
	}

	return helper.SendResponse(ctx, config.SuccessCode, result, config.SuccessMessage, nil)
}

// WorkflowRunner wraps workflow logic with panic recovery.
// This helper ensures that panics within workflows are converted to non-retryable errors,
// allowing for fast failure detection and preventing unnecessary retry loops.
//
// Parameters:
//   - fn: func() (any, error) – the workflow logic function that returns a result and error
//
// Returns:
//   - res: any – the result returned by the workflow logic function
//   - runnerErr: error – any error from the function or a NonRetryableApplicationError if panic occurred
//
// Behavior:
//   - Catches panics using defer/recover and converts them to NonRetryableApplicationError
//     with error type "WORKFLOW_PANIC"
//   - Ensures workflows fail fast on panic rather than retrying indefinitely
//
// Example:
//
//	func CreateUserWorkflow(ctx workflow.Context, input model.CreateUserInput) (res any, err error) {
//	    ctx = workflow.WithActivityOptions(ctx, ao)
//
//		defer func() {
//			if err != nil {
//				disconnectCtx, _ := workflow.NewDisconnectedContext(ctx)
//				// Compensate here
//			}
//		}()
//
//	    return temporal_helper.WorkflowRunner(func() (any, error) {
//	        // Execute activities with the provided options
//	        var userID int
//	        if err := workflow.ExecuteActivity(ctx, CreateUserDataActivity, input).Get(ctx, &userID); err != nil {
//	            return nil, err
//	        }
//
//	        return &model.CreateUserOutput{UserID: userID}, nil
//	    })
//	}
func SafeRunner[T any](
	fn func() (res T, err error),
) (res T, runnerErr error) {
	defer func() {
		if r := recover(); r != nil {
			// Convert panic to a NonRetryableApplicationError
			runnerErr = temporal.NewNonRetryableApplicationError(
				fmt.Sprintf("workflow panicked: %v", r),
				"WORKFLOW_PANIC",
				nil,
			)
		}
	}()

	res, runnerErr = fn()
	runnerErr = StopWorkflowRetryOnNonRetryable(runnerErr)
	return res, runnerErr
}

// StopWorkflowRetryOnNonRetryable inspects an error returned from an activity (or any nested error)
// and, if it contains a NonRetryable ApplicationError, converts it into a top-level
// NonRetryableApplicationError so that the WORKFLOW retry policy will not retry the whole workflow.
//
// Why this is needed:
//   - Activity non-retryable errors stop ACTIVITY retries, but the workflow may still be retried
//     if the workflow has a retry policy, because the top-level failure is an ActivityFailure.
//   - By returning a NonRetryableApplicationError at the WORKFLOW level, we signal Temporal to stop
//     retrying the entire workflow execution as well.
//
// Usage inside a workflow:
//
//	var out SomeResult
//	if err := workflow.ExecuteActivity(ctx, DoSomething, arg).Get(ctx, &out); err != nil {
//	    return nil, temporal_helper.StopWorkflowRetryOnNonRetryable(err)
//	}
//
// If the incoming error is not (or does not contain) a non-retryable ApplicationError,
// it is returned unchanged.
func StopWorkflowRetryOnNonRetryable(err error) error {
	if err == nil {
		return nil
	}

	// Try to locate an ApplicationError anywhere in the chain
	var appErr *temporal.ApplicationError
	if errors.As(err, &appErr) && appErr != nil && appErr.NonRetryable() {
		// Try to preserve details if present
		var details []interface{}
		if appErr.HasDetails() {
			var d any
			if derr := appErr.Details(&d); derr == nil {
				details = append(details, d)
			}
		}

		// Return a new top-level non-retryable application error so that workflow retry stops
		return temporal.NewNonRetryableApplicationError(
			appErr.Message(),
			appErr.Type(),
			err, // keep original error as cause for observability
			details...,
		)
	}

	// Not a non-retryable application error -> return as-is
	return err
}
