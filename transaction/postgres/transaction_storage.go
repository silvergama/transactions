package postgres

import (
	"context"
	"database/sql"

	"github.com/silvergama/transations/transaction"
)

type Repository struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Repository {
	return &Repository{db}
}

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
