package transaction_test

import (
	"context"
	"errors"
	"testing"

	"github.com/silvergama/transations/transaction"
	"github.com/silvergama/transations/transaction/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCaseCreate(t *testing.T) {
	tests := []struct {
		name          string
		mockedRepo    func(m *mocks.Repository)
		transaction   *transaction.Transaction
		expectedID    int
		expectedError error
	}{
		{
			name: "Should create account successfully",
			mockedRepo: func(m *mocks.Repository) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(1, nil)
			},
			transaction:   &transaction.Transaction{AccountID: 1, OperationTypeID: transaction.Purchase, Amount: 100.0},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Should return an erro to create transaction",
			mockedRepo: func(m *mocks.Repository) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(0, errors.New("database error"))
			},
			transaction:   &transaction.Transaction{AccountID: 1, OperationTypeID: transaction.Purchase, Amount: 100.0},
			expectedID:    0,
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Repository{}
			tt.mockedRepo(repo)

			useCase := transaction.NewService(repo)
			id, err := useCase.Create(context.Background(), tt.transaction)

			assert.Equal(t, tt.expectedID, id)
			assert.Equal(t, tt.expectedError, err)

			repo.AssertExpectations(t)
		})
	}
}

func TestProcessTransaction(t *testing.T) {
	type args struct {
		t *transaction.Transaction
	}
	tests := []struct {
		name string
		args args
		want *transaction.Transaction
	}{
		{
			name: "should convert amout to negative when the operation type is Purchase",
			args: args{
				t: &transaction.Transaction{
					OperationTypeID: 1,
					Amount:          10,
				},
			},
			want: &transaction.Transaction{
				OperationTypeID: 1,
				Amount:          -10,
			},
		},
		{
			name: "should convert amout to negative when the operation type is Payment",
			args: args{
				t: &transaction.Transaction{
					OperationTypeID: 4,
					Amount:          10,
				},
			},
			want: &transaction.Transaction{
				OperationTypeID: 4,
				Amount:          10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction.ProcessTransaction(tt.args.t)
			assert.Equal(t, tt.args.t, tt.want)
		})
	}
}
