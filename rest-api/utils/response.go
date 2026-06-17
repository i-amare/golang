package utils

import (
	"encoding/json"
	"net/http"
)

// Envelope is a shared JSON response format used across the API.
type Envelope struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// WriteJSON writes a JSON payload with the given status code.
func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// WriteSuccess sends a consistent success response.
func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	WriteJSON(w, status, Envelope{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// WriteError sends a consistent error response.
func WriteError(w http.ResponseWriter, status int, message string, err interface{}) {
	WriteJSON(w, status, Envelope{
		Success: false,
		Message: message,
		Error:   err,
	})
}
