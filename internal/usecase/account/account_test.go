package account_test

import (
	"context"
	"errors"
	"testing"

	"github.com/silvergama/transations/internal/domain"
	"github.com/silvergama/transations/internal/domain/mocks"
	"github.com/silvergama/transations/internal/usecase/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestNewAccountService(t *testing.T) {
// 	type args struct {
// 		r domain.AccountRepositoryInterface
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want account.UseCase
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewAccountService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewAccountService() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestUseCaseCreate(t *testing.T) {
	tests := []struct {
		name          string
		mockedRepo    func(m *mocks.AccountRepositoryInterface)
		account       *domain.Account
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockedRepo: func(m *mocks.AccountRepositoryInterface) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(1, nil)
			},
			account:       &domain.Account{DocumentNumber: "123456789"},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Error",
			mockedRepo: func(m *mocks.AccountRepositoryInterface) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(0, errors.New("fail"))
			},
			account:       &domain.Account{DocumentNumber: "123456789"},
			expectedID:    0,
			expectedError: errors.New("fail"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.AccountRepositoryInterface{}
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
		mockedRepo    func(m *mocks.AccountRepositoryInterface)
		accountID     int
		want          *domain.Account
		wantErr       bool
		expectedError error
	}{
		{
			name: "Should create account successfully",
			mockedRepo: func(m *mocks.AccountRepositoryInterface) {
				m.On("GetByID", mock.Anything, mock.Anything).
					Return(&domain.Account{AccoundID: 1, DocumentNumber: "123456789"}, nil)
			},
			accountID:     1,
			want:          &domain.Account{AccoundID: 1, DocumentNumber: "123456789"},
			expectedError: nil,
		},
		{
			name: "Should return an erro to create account",
			mockedRepo: func(m *mocks.AccountRepositoryInterface) {
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
			repo := &mocks.AccountRepositoryInterface{}
			tt.mockedRepo(repo)

			useCase := account.NewService(repo)
			got, err := useCase.GetByID(context.Background(), tt.accountID)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.expectedError, err)

			repo.AssertExpectations(t)
		})
	}
}
