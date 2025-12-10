package temporal_helper

import (
	"fmt"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/entity"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/helper"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/service"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
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
	result *T,
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

// WorkflowRunner wraps workflow logic with panic recovery and standard activity options.
// This helper ensures that panics within workflows are converted to non-retryable errors,
// allowing for fast failure detection and preventing unnecessary retry loops.
//
// Parameters:
//   - fn: func(ao workflow.ActivityOptions) (any, error) – the workflow logic function that receives
//     pre-configured activity options and returns a result and error
//   - customOptions: ...workflow.ActivityOptions (optional) – override default activity options if provided
//
// Returns:
//   - res: any – the result returned by the workflow logic function
//   - runnerErr: error – any error from the function or a NonRetryableApplicationError if panic occurred
//
// Behavior:
//   - Provides standard ActivityOptions with:
//   - StartToCloseTimeout: 1 minute
//   - RetryPolicy: MaximumAttempts of 2
//   - Catches panics using defer/recover and converts them to NonRetryableApplicationError
//     with error type "WORKFLOW_PANIC"
//   - Ensures workflows fail fast on panic rather than retrying indefinitely
//   - If customOptions is provided, it completely replaces the default activity options
//
// Example:
//
//	func CreateUserWorkflow(ctx workflow.Context, input model.CreateUserInput) (*model.CreateUserOutput, error) {
//	    res, err := temporal_helper.WorkflowRunner(func(ao workflow.ActivityOptions) (any, error) {
//	        ctx = workflow.WithActivityOptions(ctx, ao)
//
//	        // Execute activities with the provided options
//	        var userID int
//	        if err := workflow.ExecuteActivity(ctx, CreateUserDataActivity, input).Get(ctx, &userID); err != nil {
//	            return nil, err
//	        }
//
//	        return &model.CreateUserOutput{UserID: userID}, nil
//	    })
//
//	    if err != nil {
//	        return nil, err
//	    }
//	    return res.(*model.CreateUserOutput), nil
//	}
func WorkflowRunner(fn func(ao workflow.ActivityOptions) (any, err error), customOptions ...workflow.ActivityOptions) (res any, runnerErr error) {
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

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 2,
		},
	}

	if len(customOptions) > 0 {
		ao = customOptions[0]
	}

	res, runnerErr = fn(ao)
	return res, runnerErr
}
