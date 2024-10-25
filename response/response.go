package response

import (
	"encoding/json"
	"net/http"
)

// sendResponse is a helper to send JSON responses
func SendResponse(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	response := map[string]interface{}{
		"statusCode": statusCode,
		"data":       data,
		"message":    message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
