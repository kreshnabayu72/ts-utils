package entity

import (
	"errors"
	"log/slog"
	"strings"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/helper"
	"go.temporal.io/sdk/temporal"
)

// Define constants for Temporal ApplicationError types
const (
	ErrBadRequest          = "BadRequest"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrNotFound            = "NotFound"
	ErrConflict            = "Conflict"
	ErrInternalServerError = "InternalServerError"
)

type WorkflowErrorWrapperParams struct {
	ErrMessage   string
	ErrType      string
	Data         any
	NonRetryable bool
}

func WorkflowErrorWrapper(params WorkflowErrorWrapperParams) error {
	var details []interface{}
	if params.Data != nil {
		details = append(details, params.Data)
	}

	// If NonRetryable is explicitly set to true, use NewNonRetryableApplicationError instead.
	if params.NonRetryable {
		return temporal.NewNonRetryableApplicationError(
			params.ErrMessage,
			params.ErrType,
			nil, // cause (not used?)
			details...,
		)
	}

	return temporal.NewApplicationError(
		params.ErrMessage,
		params.ErrType,
		details...,
	)
}

// WrapHttpError is a convenience function for lazy writing (support nonRetryable flag)
func WrapHttpError(err *HttpError, nonRetryable ...bool) error {
	if err == nil {
		return nil
	}

	// Map HttpError code to appropriate string type
	// This allows ParseWorkflowError to extract HttpError from ApplicationError
	errType := ErrInternalServerError
	switch err.Code {
	case 400:
		errType = ErrBadRequest
	case 401:
		errType = ErrUnauthorized
	case 403:
		errType = ErrForbidden
	case 404:
		errType = ErrNotFound
	case 409:
		errType = ErrConflict
	}

	if helper.IsZeroValue(nonRetryable) {
		nonRetryable = []bool{false}
	}

	return WorkflowErrorWrapper(WorkflowErrorWrapperParams{
		ErrMessage:   err.Message,
		ErrType:      errType,
		Data:         err.Data,
		NonRetryable: nonRetryable[0],
	})
}

// WorkflowBadRequest is a helper to return a wrapped BadRequest error
func WorkflowBadRequest(message string, nonRetryable ...bool) error {
	return WrapHttpError(BadRequest(message), nonRetryable...)
}

// WorkflowUnauthorized is a helper to return a wrapped Unauthorized error
func WorkflowUnauthorized(message string, nonRetryable ...bool) error {
	return WrapHttpError(Unauthorized(message), nonRetryable...)
}

// WorkflowForbidden is a helper to return a wrapped Forbidden error
func WorkflowForbidden(message string, nonRetryable ...bool) error {
	return WrapHttpError(Forbidden(message), nonRetryable...)
}

// WorkflowNotFound is a helper to return a wrapped NotFound error
func WorkflowNotFound(message string, nonRetryable ...bool) error {
	return WrapHttpError(NotFound(message), nonRetryable...)
}

// WorkflowConflict is a helper to return a wrapped Conflict error
func WorkflowConflict(message string, nonRetryable ...bool) error {
	return WrapHttpError(Conflict(message), nonRetryable...)
}

// WorkflowInternalServerError is a helper to return a wrapped InternalServerError error
func WorkflowInternalServerError(message string, nonRetryable ...bool) error {
	return WrapHttpError(InternalServerError(message), nonRetryable...)
}

// ParseWorkflowError extracts user-friendly error messages from Temporal workflow errors.
// Usage in controller as error catch when getting workflow result
func ParseWorkflowError(workflowName string, err error) *HttpError {
	if err == nil {
		return nil
	}

	// Try to extract ApplicationError details from Temporal-wrapped errors
	var extractedData any
	var appErr *temporal.ApplicationError
	if errors.As(err, &appErr) {
		// Check if the ApplicationError has details
		if appErr.HasDetails() {
			// Try to get HttpError from ApplicationError details
			if err := appErr.Details(&extractedData); err != nil {
				slog.Warn("Failed to extract details from ApplicationError", "error", err)
			}
		}
	}

	var httpErr HttpError
	// If we have an ApplicationError, try to map it to a HttpError
	if appErr != nil {
		switch appErr.Type() {
		case ErrForbidden:
			httpErr = *Forbidden(appErr.Message())
		case ErrNotFound:
			httpErr = *NotFound(appErr.Message())
		case ErrBadRequest:
			httpErr = *BadRequest(appErr.Message())
		case ErrUnauthorized:
			httpErr = *Unauthorized(appErr.Message())
		case ErrConflict:
			httpErr = *Conflict(appErr.Message())
		default:
			httpErr = *InternalServerError(appErr.Message())
		}

		if extractedData != nil {
			httpErr.Data = extractedData
		}

		// If we successfully mapped the type, return it
		if httpErr.Code != 0 {
			activityName := extractActivityName(err.Error())
			// Logging for detailed error information
			slog.Error("Error", "workflowName", workflowName, "activityName", activityName, "error", httpErr.Message, "httpCode", httpErr.Code)
			return &httpErr
		}
	}

	// Final fallback if everything else fails (since workflow's error is not HttpError, status code is most likely 500 as the final fallback)
	return ToHttpError(err)
}

// extractActivityName extracts the activity name from workflow error
func extractActivityName(errorStr string) string {
	// Pattern: "activity error (type: ActivityName, ..."
	if idx := strings.Index(errorStr, "activity error (type: "); idx != -1 {
		start := idx + len("activity error (type: ")
		if end := strings.Index(errorStr[start:], ","); end != -1 {
			return errorStr[start : start+end]
		}
	}
	return ""
}
