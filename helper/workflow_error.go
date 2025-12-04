package helper

import (
	"log/slog"
	"strings"
)

// ParseWorkflowError extracts user-friendly error messages from Temporal workflow errors
func ParseWorkflowError(workflowName string, err error) string {
	if err == nil {
		return ""
	}

	errorStr := err.Error()

	// Extract activity name if present
	activityName := extractActivityName(errorStr)

	// Extract the innermost error message by removing workflow/activity metadata
	cleanedError := extractInnerError(errorStr)

	slog.Error("Error", "workflowName", workflowName, "activityName", activityName, "cleanedError", cleanedError)

	// Format: (WorkflowName) error message
	return cleanedError
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

// extractInnerError removes Temporal workflow/activity metadata
func extractInnerError(errorStr string) string {
	// Remove trailing metadata like "(type: ..., retryable: ...)"
	if idx := strings.LastIndex(errorStr, " (type:"); idx != -1 {
		errorStr = errorStr[:idx]
	}

	// Extract message after the last "): "
	if idx := strings.LastIndex(errorStr, "): "); idx != -1 {
		errorStr = errorStr[idx+3:]
	}

	return strings.TrimSpace(errorStr)
}
