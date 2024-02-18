package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name   string
		body   interface{}
		status int
	}{
		{
			name:   "Success",
			body:   Response{Message: "Success", Data: "Data"},
			status: http.StatusOK,
		},
		{
			name:   "CustomStatus",
			body:   Response{Message: "Custom Status", Data: "Custom Data"},
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			Write(w, tt.body, tt.status)

			assert.Equal(t, tt.status, w.Code)

			var response Response
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, tt.body, response)
		})
	}
}

func TestWriteServerError(t *testing.T) {
	w := httptest.NewRecorder()
	message := "Internal Server Error"
	WriteServerError(w, message)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response Error
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	assert.Equal(t, message, response.Message)
}

func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	message := "Not Found"
	WriteNotFound(w, message)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response Error
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, message, response.Message)
}

func TestWriteBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	message := "Bad Request"
	WriteBadRequest(w, message)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response Error
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, message, response.Message)
}

func TestWriteUnprocessableEntity(t *testing.T) {
	w := httptest.NewRecorder()
	message := "Unprocessable Entity"
	WriteUnprocessableEntity(w, message)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var response Error
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)
	assert.Equal(t, message, response.Message)
}
