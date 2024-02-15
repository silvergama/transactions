package mux

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/silvergama/transations/pkg/response"
	"github.com/silvergama/transations/transaction"
	"github.com/silvergama/transations/transaction/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionHandler(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      string
		expectedStatus   int
		expectedResponse response.Response
		serviceMock      func(m *mocks.UseCase) // Use o mock gerado pelo mockery
	}{
		{
			name:           "Success",
			requestBody:    `{"operation_type_id": 1, "amount": 100}`,
			expectedStatus: http.StatusCreated,
			expectedResponse: response.Response{
				Message: "transaction created successfully",
				Data:    map[string]interface{}{"transaction_id": float64(1)},
			},
			serviceMock: func(m *mocks.UseCase) {
				m.On("Create", mock.Anything, mock.Anything).Return(1, nil)
			},
		},
		{
			name:           "InvalidOperationType",
			requestBody:    `{"operation_type_id": 99, "amount": 100}`,
			expectedStatus: http.StatusBadRequest,
			expectedResponse: response.Response{
				Message: "invalid operation type",
			},
			serviceMock: func(m *mocks.UseCase) {},
		},
		{
			name:           "InvalidRequestBody",
			requestBody:    `{"invalid_json"`,
			expectedStatus: http.StatusBadRequest,
			expectedResponse: response.Response{
				Message: "failed to decode payload",
			},
			serviceMock: func(m *mocks.UseCase) {},
		},
		{
			name:           "ServerError",
			requestBody:    `{"operation_type_id": 1, "amount": 100}`,
			expectedStatus: http.StatusInternalServerError,
			expectedResponse: response.Response{
				Message: "failed to create transaction",
			},
			serviceMock: func(m *mocks.UseCase) {
				m.On("Create", mock.Anything, mock.Anything).Return(0, errors.New("fail"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := &mocks.UseCase{}

			tt.serviceMock(useCase)

			transaction.NewService(useCase)
			handler := NewTransactionHandler(useCase)

			req, err := http.NewRequest("POST", "/transactions", strings.NewReader(tt.requestBody))
			assert.NoError(t, err)

			router := mux.NewRouter()
			router.HandleFunc("/transactions", handler.CreateTransactionHandler).Methods(http.MethodPost)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var responseObj response.Response
			err = json.Unmarshal(rr.Body.Bytes(), &responseObj)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedResponse, responseObj)

			if tt.serviceMock != nil {
				useCase.AssertExpectations(t)
			}
		})
	}
}
