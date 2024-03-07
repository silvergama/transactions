package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/silvergama/transations/internal/domain"
	"github.com/silvergama/transations/pkg/logger"
	"go.uber.org/zap"
)

// UseCase defines the business logic for transaction operations
type UseCase interface {
	Create(ctx context.Context, transaction *domain.Transaction) (int, error)
}

// Service represents the business logic for transaction operations
type service struct {
	transactionRepo domain.TransactionRepositoryInterface
	accountRepo     domain.AccountRepositoryInterface
}

// NewService creates a new Service instance with the provided repository
func NewService(t domain.TransactionRepositoryInterface, a domain.AccountRepositoryInterface,
) UseCase {
	return &service{
		transactionRepo: t,
		accountRepo:     a,
	}
}

// ProcessTransaction updates the account balance based on the transaction type
func ProcessTransaction(account *domain.Account, transaction *domain.Transaction) {

	if transaction.OperationTypeID == domain.Payment {
		account.AvailableCreditLimit = account.AvailableCreditLimit + transaction.Amount
		return
	}

	account.AvailableCreditLimit = account.AvailableCreditLimit - transaction.Amount
	transaction.Amount = -transaction.Amount
}

func (s *service) Create(ctx context.Context, transaction *domain.Transaction) (int, error) {
	if !isValidOperationType(transaction.OperationTypeID) {
		return 0, errors.New("invalid transaction type")
	}
	account, err := s.accountRepo.GetByID(ctx, transaction.AccountID)
	if err != nil {
		return 0, err
	}

	ProcessTransaction(account, transaction)

	if !s.hasLimit(account.AvailableCreditLimit) {
		return 0, fmt.Errorf("no balance available")
	}

	transactionID, err := s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		logger.Error("failed to create transaction",
			zap.Error(err),
			zap.Any("request_id", ctx.Value("request_id")),
			zap.Any("transaction", transaction),
		)
		return 0, err
	}

	// update account limit

	return transactionID, nil
}

// validateBalance determines if the transaction amount is within the account balance
func (s *service) hasLimit(limit float64) bool {
	return limit >= 0
}

// isValidOperationType determines if the given operation type is valid
func isValidOperationType(operationType domain.OperationType) bool {
	switch operationType {
	case domain.Purchase, domain.Installment, domain.Withdrawal, domain.Payment:
		return true
	default:
		return false
	}
}
