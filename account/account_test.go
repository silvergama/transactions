package account_test

import (
	"context"
	"errors"
	"testing"

	"github.com/silvergama/transations/account"
	"github.com/silvergama/transations/account/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCaseCreate(t *testing.T) {
	tests := []struct {
		name          string
		mockedRepo    func(m *mocks.Repository)
		account       *account.Account
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockedRepo: func(m *mocks.Repository) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(1, nil)
			},
			account:       &account.Account{DocumentNumber: "123456789"},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Error",
			mockedRepo: func(m *mocks.Repository) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(0, errors.New("fail"))
			},
			account:       &account.Account{DocumentNumber: "123456789"},
			expectedID:    0,
			expectedError: errors.New("fail"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Repository{}
			tt.mockedRepo(repo)

			useCase := account.NewService(repo)
			id, err := useCase.Create(context.Background(), tt.account)

			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedError, err)

			repo.AssertExpectations(t)
		})
	}
}

func TestServiceGetByID(t *testing.T) {
	tests := []struct {
		name          string
		mockedRepo    func(m *mocks.Repository)
		accountID     int
		want          *account.Account
		wantErr       bool
		expectedError error
	}{
		{
			name: "Should create account successfully",
			mockedRepo: func(m *mocks.Repository) {
				m.On("GetByID", mock.Anything, mock.Anything).
					Return(&account.Account{AccoundID: 1, DocumentNumber: "123456789"}, nil)
			},
			accountID:     1,
			want:          &account.Account{AccoundID: 1, DocumentNumber: "123456789"},
			expectedError: nil,
		},
		{
			name: "Should return an erro to create account",
			mockedRepo: func(m *mocks.Repository) {
				m.On("GetByID", mock.Anything, mock.Anything).
					Return(nil, errors.New("fail"))
			},
			accountID:     1,
			want:          nil,
			expectedError: errors.New("fail"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Repository{}
			tt.mockedRepo(repo)

			useCase := account.NewService(repo)
			got, err := useCase.GetByID(context.Background(), tt.accountID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.expectedError, err)

			repo.AssertExpectations(t)
		})
	}
}
