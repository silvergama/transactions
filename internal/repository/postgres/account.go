package postgres

import (
	"context"
	"database/sql"

	"github.com/silvergama/transations/internal/domain"
)

// AccountRepository represents the PostgreSQL implementation of the domain.AccountRepository interface
type accountRepository struct {
	db *sql.DB
}

// NewAccount creates a new instance of the Repository
func NewAccount(db *sql.DB) *accountRepository {
	return &accountRepository{
		db,
	}
}

// GetByID retrieves an account by its ID
func (r *accountRepository) GetByID(context context.Context, id int) (*domain.Account, error) {
	var acc domain.Account

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
func (r *accountRepository) Create(context context.Context, acc *domain.Account) (int, error) {
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

func (r *accountRepository) UpdateLimit(ctx context.Context, idAccount int, limit float64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE INTO account (id, limit)
		VALUES ($1, $2)`, idAccount, limit)
	if err != nil {
		return err
	}

	return nil
}
