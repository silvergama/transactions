package response

import (
	"encoding/json"
	"net/http"
)

// Response represents a common structure for success and error responses
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Error represents an error response
type Error struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Write writes the response to the provided http.ResponseWriter
func Write(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

// WriteServerError writes a server error response
func WriteServerError(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}, http.StatusInternalServerError)
}

// WriteNotFound writes a not found response
func WriteNotFound(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}, http.StatusNotFound)
}

// WriteBadRequest writes a bad request response
func WriteBadRequest(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}, http.StatusBadRequest)
}

// WriteUnprocessableEntity writes an unprocessable entity response
func WriteUnprocessableEntity(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}, http.StatusUnprocessableEntity)
}
