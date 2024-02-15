package transaction

import (
	"context"

	"github.com/silvergama/transations/pkg/logger"
	"go.uber.org/zap"
)

// Service represents the business logic for transaction operations
type Service struct {
	repo Repository
}

// NewService creates a new Service instance with the provided repository
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

// CreateTransactionHandler handles the creation of a new transaction
func (s *Service) Create(ctx context.Context, transaction *Transaction) (int, error) {

	ProcessTransaction(transaction)

	transactionID, err := s.repo.Create(ctx, transaction)
	if err != nil {
		logger.Error("failed to create transaction",
			zap.Error(err),
			zap.Any("request_id", ctx.Value("request_id")),
			zap.Any("transaction", transaction),
		)
		return 0, err
	}

	return transactionID, nil
}

// ProcessTransaction processes the given transaction
func ProcessTransaction(t *Transaction) {
	switch t.OperationTypeID {
	case Purchase, Withdrawal, Installment:
		t.Amount = -t.Amount
	case Payment:
		t.Amount = +t.Amount
	default:
		return
	}
}
