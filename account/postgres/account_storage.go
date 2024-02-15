package postgres

import (
	"context"
	"database/sql"

	"github.com/silvergama/transations/account"
)

type Repository struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) GetByID(context context.Context, id int) (*account.Account, error) {
	var account account.Account

	query := "SELECT id, document_number FROM accounts WHERE id = $1"

	if err := r.db.QueryRowContext(context, query, id).Scan(&account.AccoundID, &account.DocumentNumber); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &account, nil
}

func (r *Repository) Create(context context.Context, account *account.Account) (int, error) {
	var accountID int
	err := r.db.QueryRowContext(context, `
		INSERT INTO accounts (document_number)
		VALUES ($1)
		RETURNING id
	`, account.DocumentNumber).Scan(&accountID)

	if err != nil {
		return 0, err
	}

	return accountID, nil
}
