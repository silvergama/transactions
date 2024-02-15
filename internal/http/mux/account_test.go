package mux

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/silvergama/transations/account"
	"github.com/silvergama/transations/account/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAccountHandler(t *testing.T) {
	tests := []struct {
		name            string
		requestURL      string
		expectedStatus  int
		expectedAccount *account.Account       // Preencha com o valor esperado para a conta
		serviceMock     func(m *mocks.UseCase) // Use o mock gerado pelo mockery
	}{
		{
			name:           "Success",
			requestURL:     "/accounts/1",
			expectedStatus: http.StatusOK,
			expectedAccount: &account.Account{
				AccoundID:      1,
				DocumentNumber: "1010101",
			},
			serviceMock: func(m *mocks.UseCase) {
				m.On("GetByID", mock.Anything, 1).Return(&account.Account{
					AccoundID:      1,
					DocumentNumber: "1010101",
				}, nil)
			},
		},
		{
			name:            "InvalidID",
			requestURL:      "/accounts/invalid",
			expectedStatus:  http.StatusNotFound,
			expectedAccount: nil,
			serviceMock:     func(m *mocks.UseCase) {},
		},
		{
			name:            "NotFound",
			requestURL:      "/accounts/2",
			expectedStatus:  http.StatusNotFound,
			expectedAccount: nil,
			serviceMock: func(m *mocks.UseCase) {
				m.On("GetByID", mock.Anything, 2).Return(nil, errors.New("not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := &mocks.UseCase{}
			tt.serviceMock(useCase)

			account.NewService(useCase)
			handler := NewAccountHandler(useCase)

			req, err := http.NewRequest("GET", tt.requestURL, nil)
			assert.NoError(t, err)

			router := mux.NewRouter()
			router.HandleFunc("/accounts/{id:[0-9]+}", handler.GetAccountHandler).Methods("GET")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.serviceMock != nil {
				useCase.AssertExpectations(t)
			}
		})
	}
}
