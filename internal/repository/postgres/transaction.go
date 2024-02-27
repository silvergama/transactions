package postgres

import (
	"context"
	"database/sql"

	"github.com/silvergama/transations/internal/domain"
)

// transactionRepository represents the database repository for transactions
type transactionRepository struct {
	db *sql.DB
}

// NewTransaction creates a new instance of the Transaction transactionrepository
func NewTransaction(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

// Create creates a new transaction in the database
func (r *transactionRepository) Create(ctx context.Context, transaction *domain.Transaction) (int, error) {
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
