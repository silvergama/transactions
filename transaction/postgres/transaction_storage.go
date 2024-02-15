package postgres

import (
	"context"
	"database/sql"

	"github.com/silvergama/transations/transaction"
)

// Repository represents the database repository for transactions
type Repository struct {
	db *sql.DB
}

// NewTransaction creates a new instance of the Transaction repository
func NewTransaction(db *sql.DB) *Repository {
	return &Repository{db}
}

// Create creates a new transaction in the database
func (r *Repository) Create(ctx context.Context, transaction *transaction.Transaction) (int, error) {
	var transactionID int
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO transactions (account_id, operation_type_id, amount)
		VALUES ($1, $2, $3)
		RETURNING id`, transaction.AccountID, transaction.OperationTypeID, transaction.Amount).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}
