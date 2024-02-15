package account

import "context"

type Account struct {
	AccoundID      int    `json:"accound_id,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
}

type Reader interface {
	GetByID(context context.Context, id int) (*Account, error)
}

type Writer interface {
	Create(context context.Context, account *Account) (int, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Create(context context.Context, account *Account) (int, error)
	GetByID(context context.Context, id int) (*Account, error)
}
