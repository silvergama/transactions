package transaction

import (
	"context"

	"github.com/silvergama/transations/pkg/logger"
	"go.uber.org/zap"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

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
