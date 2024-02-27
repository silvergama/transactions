package transaction

import (
	"context"

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
func NewService(t domain.TransactionRepositoryInterface, a domain.AccountRepositoryInterface) UseCase {
	return &service{
		transactionRepo: t,
		accountRepo:     a,
	}
}

// Create handles the creation of a new transaction
func (s *service) Create(ctx context.Context, transaction *domain.Transaction) (int, error) {

	ProcessTransaction(transaction)

	_, err := s.accountRepo.GetByID(ctx, transaction.AccountID)
	if err != nil {
		return 0, err
	}

	// validate balance

	// account, err := s.(ctx, transaction)
	// if !s.validateBalance(transaction.Amount, account.AvaillableCreditLimit) {
	// 	//log
	// 	return 0, err
	// }

	transactionID, err := s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		logger.Error("failed to create transaction",
			zap.Error(err),
			zap.Any("request_id", ctx.Value("request_id")),
			zap.Any("transaction", transaction),
		)
		return 0, err
	}

	// err, a := s.AccountRepo.UpdateLimit(limit)
	// if err != nil {
	// 	//log
	// 	return 0, nil
	// }

	return transactionID, nil
}

// ProcessTransaction processes the given transaction
func ProcessTransaction(t *domain.Transaction) {
	switch t.OperationTypeID {
	case domain.Purchase, domain.Withdrawal, domain.Installment:
		t.Amount = -t.Amount
	case domain.Payment:
		t.Amount = +t.Amount
	default:
		return
	}
}

// func (s *service) validateBalance(value, limit float64) bool {
// 	return value <= limit
// }
