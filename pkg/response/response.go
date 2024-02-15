package response

import (
	"encoding/json"
	"net/http"
)

// Response representa uma estrutura comum para respostas de sucesso e erro
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func Write(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func WriteServerError(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}, http.StatusInternalServerError)
}

func WriteNotFound(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}, http.StatusNotFound)
}

func WriteBadRequest(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}, http.StatusBadRequest)
}

func WriteUnprocessableEntity(w http.ResponseWriter, message string) {
	Write(w, Error{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}, http.StatusUnprocessableEntity)
}
