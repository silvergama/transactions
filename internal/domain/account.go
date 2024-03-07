package domain

import "context"

// Account represents the account entity with its attributes
type Account struct {
	AccoundID            int     `json:"accound_id,omitempty"`
	DocumentNumber       string  `json:"document_number,omitempty"`
	AvailableCreditLimit float64 `json:"avalable_credit_limit,omitempty"`
}

// Reader provides methods for reading account information
type Reader interface {
	GetByID(context context.Context, id int) (*Account, error)
}

// Writer provides methods for creating account records
type Writer interface {
	Create(context context.Context, account *Account) (int, error)
}

// Repository combines the Reader and Writer interfaces to represent a complete set of account data operations
type AccountRepositoryInterface interface {
	Reader
	Writer
}
