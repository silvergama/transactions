package postgres

import (
	"context"
	"database/sql"

	"github.com/silvergama/transations/account"
)

// Repository represents the PostgreSQL implementation of the account.Repository interface
type Repository struct {
	db *sql.DB
}

// NewAccount creates a new instance of the Repository
func NewAccount(db *sql.DB) *Repository {
	return &Repository{
		db,
	}
}

// GetByID retrieves an account by its ID
func (r *Repository) GetByID(context context.Context, id int) (*account.Account, error) {
	var acc account.Account

	query := "SELECT id, document_number FROM accounts WHERE id = $1"

	if err := r.db.QueryRowContext(context, query, id).Scan(&acc.AccoundID, &acc.DocumentNumber); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &acc, nil
}

// Create creates a new account
func (r *Repository) Create(context context.Context, acc *account.Account) (int, error) {
	var accountID int
	err := r.db.QueryRowContext(context, `
		INSERT INTO accounts (document_number)
		VALUES ($1)
		RETURNING id
	`, acc.DocumentNumber).Scan(&accountID)

	if err != nil {
		return 0, err
	}

	return accountID, nil
}
