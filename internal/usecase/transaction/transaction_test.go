package transaction

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/silvergama/transations/internal/domain"
	"github.com/silvergama/transations/internal/domain/mocks"
)

func TestServiceCreate(t *testing.T) {

	type args struct {
		ctx         context.Context
		transaction *domain.Transaction
	}
	tests := []struct {
		name    string
		prepare func(t *mocks.TransactionRepositoryInterface, a *mocks.AccountRepositoryInterface)
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Success Purchase",
			prepare: func(t *mocks.TransactionRepositoryInterface, a *mocks.AccountRepositoryInterface) {
				a.On("GetByID", mock.Anything, mock.Anything).Return(&domain.Account{
					AccoundID:            1,
					DocumentNumber:       "123",
					AvailableCreditLimit: 100,
				}, nil)
				t.On("Create", mock.Anything, mock.Anything).Return(1, nil)
			},
			args: args{
				ctx: context.Background(),
				transaction: &domain.Transaction{
					ID:              1,
					AccountID:       2,
					OperationTypeID: 1,
					Amount:          20.0,
				},
			},
			want: 1,
		},
		{
			name: "InvalidAccount",
			prepare: func(t *mocks.TransactionRepositoryInterface, a *mocks.AccountRepositoryInterface) {
				a.On("GetByID", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("account not found"))
			},
			args: args{
				ctx: context.Background(),
				transaction: &domain.Transaction{
					ID:              3,
					AccountID:       4,
					OperationTypeID: 1,
					Amount:          30.0,
				},
			},
			wantErr: true,
		},
		{
			name: "InsufficientBalance",
			prepare: func(t *mocks.TransactionRepositoryInterface, a *mocks.AccountRepositoryInterface) {
				a.On("GetByID", mock.Anything, mock.Anything).Return(&domain.Account{
					AccoundID:            1,
					DocumentNumber:       "123",
					AvailableCreditLimit: 10, // Defina um limite de crédito inferior ao valor da transação
				}, nil)
			},
			args: args{
				ctx: context.Background(),
				transaction: &domain.Transaction{
					ID:              2,
					AccountID:       3,
					OperationTypeID: 1,
					Amount:          20.0,
				},
			},
			wantErr: true,
		},
		{
			name: "TransactionCreationFailure",
			prepare: func(t *mocks.TransactionRepositoryInterface, a *mocks.AccountRepositoryInterface) {
				a.On("GetByID", mock.Anything, mock.Anything).Return(&domain.Account{
					AccoundID:            1,
					DocumentNumber:       "123",
					AvailableCreditLimit: 100,
				}, nil)
				t.On("Create", mock.Anything, mock.Anything).Return(0, fmt.Errorf("transaction creation failed"))
			},
			args: args{
				ctx: context.Background(),
				transaction: &domain.Transaction{
					ID:              4,
					AccountID:       5,
					OperationTypeID: 1,
					Amount:          40.0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tMockedRepo := &mocks.TransactionRepositoryInterface{}
			aMockedRepo := &mocks.AccountRepositoryInterface{}
			s := &service{
				transactionRepo: tMockedRepo,
				accountRepo:     aMockedRepo,
			}
			tt.prepare(tMockedRepo, aMockedRepo)
			got, err := s.Create(tt.args.ctx, tt.args.transaction)
			if err != nil && !tt.wantErr {
				assert.Fail(t, fmt.Sprintf(
					"Error not expected but got one:\n"+
						"error: %q", err),
				)
				return
			}
			assert.Equal(t, tt.want, got)

			aMockedRepo.AssertExpectations(t)
			tMockedRepo.AssertExpectations(t)
		})
	}
}
func TestProcessTransaction(t *testing.T) {
	type args struct {
		account     *domain.Account
		transaction *domain.Transaction
	}
	tests := []struct {
		name string
		args args
		want *domain.Account
	}{
		{
			name: "Payment",
			args: args{
				account: &domain.Account{
					AvailableCreditLimit: 100,
				},
				transaction: &domain.Transaction{
					OperationTypeID: domain.Payment,
					Amount:          10,
				},
			},
			want: &domain.Account{
				AvailableCreditLimit: 110,
			},
		},
		{
			name: "Withdrawal",
			args: args{
				account: &domain.Account{
					AvailableCreditLimit: 100,
				},
				transaction: &domain.Transaction{
					OperationTypeID: domain.Withdrawal,
					Amount:          10,
				},
			},
			want: &domain.Account{
				AvailableCreditLimit: 90,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProcessTransaction(tt.args.account, tt.args.transaction)
			assert.Equal(t, tt.args.account, tt.want)
		})
	}
}

func TestIsValidOperationType(t *testing.T) {
	tests := []struct {
		name     string
		input    domain.OperationType
		expected bool
	}{
		{
			name:     "valid operation type",
			input:    domain.Payment,
			expected: true,
		},
		{
			name:     "invalid operation type",
			input:    100,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isValidOperationType(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
