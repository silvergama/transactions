package transaction

import "context"

// Transaction represents a financial transaction
type Transaction struct {
	ID              int           `json:"id,omitempty"`
	AccountID       int           `json:"account_id,omitempty"`
	OperationTypeID OperationType `json:"operation_type_id,omitempty"`
	Amount          float64       `json:"amount,omitempty"`
}

// OperationType represents the type of financial operation
type OperationType int

const (
	Purchase OperationType = iota + 1
	Installment
	Withdrawal
	Payment
)

// Repository defines the interface for transaction database operations
type Repository interface {
	Create(ctx context.Context, transaction *Transaction) (int, error)
}

// UseCase defines the business logic for transaction operations
type UseCase interface {
	Create(ctx context.Context, transaction *Transaction) (int, error)
}

// isValidOperationType checks if the provided operation type is valid
func isValidOperationType(opType OperationType) bool {
	switch opType {
	case Purchase, Installment, Withdrawal, Payment:
		return true
	default:
		return false
	}
}
