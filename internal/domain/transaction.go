package domain

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
type TransactionRepositoryInterface interface {
	Create(ctx context.Context, transaction *Transaction) (int, error)
}
