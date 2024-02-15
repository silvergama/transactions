package transaction

import "context"

type Transaction struct {
	ID              int           `json:"id,omitempty"`
	AccountID       int           `json:"account_id,omitempty"`
	OperationTypeID OperationType `json:"operation_type_id,omitempty"`
	Amount          float64       `json:"amount,omitempty"`
}

type OperationType int

const (
	Purchase OperationType = iota + 1
	Installment
	Withdrawal
	Payment
)

type Repository interface {
	Create(ctx context.Context, transaction *Transaction) (int, error)
}

type UseCase interface {
	Create(ctx context.Context, transaction *Transaction) (int, error)
}

func isValidOperationType(opType OperationType) bool {
	switch opType {
	case Purchase, Installment, Withdrawal, Payment:
		return true
	default:
		return false
	}
}
