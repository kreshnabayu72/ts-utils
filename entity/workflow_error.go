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
	// If NonRetryable is explicitly set to true, use NewNonRetryableApplicationError instead.
	if params.NonRetryable {
		return temporal.NewNonRetryableApplicationError(
			params.ErrMessage,
			params.ErrType,
			nil, // cause (not used?)
			params.Data,
		)
	}

	return temporal.NewApplicationError(
		params.ErrMessage,
		params.ErrType,
		params.Data,
	)
}

// WrapHttpError is a convenience function for lazy writing (support nonRetryable flag)
func WrapHttpError(err *HttpError, nonRetryable ...bool) error {
	if err == nil {
		return nil
	}

	// Map HttpError code to appropriate string type
	// This allows ParseWorkflowError to fallback to type checking if details are lost
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
		Data:         err,
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

// ParseWorkflowError extracts user-friendly error messages from Temporal workflow errors
func ParseWorkflowError(workflowName string, err error) *HttpError {
	if err == nil {
		return nil
	}

	var httpErr HttpError

	// Handle direct HttpError returns (when not wrapped in Temporal errors)
	if directHttpErr, ok := err.(*HttpError); ok {
		slog.Error("Error", "workflowName", workflowName, "error", directHttpErr.Message, "httpCode", directHttpErr.Code)
		return directHttpErr
	}

	// Try to extract ApplicationError details from Temporal-wrapped errors
	var appErr *temporal.ApplicationError
	if errors.As(err, &appErr) {
		// Check if the ApplicationError has details
		if appErr.HasDetails() {
			// Try to get HttpError from details (4th parameter)
			if err := appErr.Details(&httpErr); err == nil {
				// Validate that we got actual data, not zero values
				if httpErr.Code != 0 || httpErr.Message != "" {
					errorStr := appErr.Error()
					activityName := extractActivityName(errorStr)

					slog.Error("Error", "workflowName", workflowName, "activityName", activityName, "error", httpErr.Message, "httpCode", httpErr.Code)
					return &httpErr
				}
			} else {
				slog.Warn("Failed to extract HttpError from details", "error", err)
			}
		}
	} else {
		// Not a Temporal ApplicationError and not an HttpError - convert generic error
		slog.Error("Converting generic error to HttpError", "error", err)
		return ToHttpError(err)
	}

	// Fallback: Check ApplicationError Type if we failed to extract details
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

		// If we successfully mapped the type, return it
		if httpErr.Code != 0 {
			slog.Error("Error", "workflowName", workflowName, "type", appErr.Type(), "message", appErr.Message(), "httpCode", httpErr.Code)
			return &httpErr
		}
	}

	// Final Fallback if everything else fails
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
