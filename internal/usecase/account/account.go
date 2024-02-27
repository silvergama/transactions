package account

import (
	"context"

	"github.com/silvergama/transations/internal/domain"
)

// UseCase defines the use cases for interacting with account data
type UseCase interface {
	Create(context context.Context, account *domain.Account) (int, error)
	GetByID(context context.Context, id int) (*domain.Account, error)
}

// Service represents the business logic for account operations
type service struct {
	repo domain.AccountRepositoryInterface
}

// NewService creates a new Service instance with the provided repository
func NewService(r domain.AccountRepositoryInterface) UseCase {
	return &service{
		repo: r,
	}
}

// Create creates a new account and returns its ID
func (s *service) Create(context context.Context, account *domain.Account) (int, error) {
	accountID, err := s.repo.Create(context, account)
	if err != nil {
		return 0, err
	}

	return accountID, nil
}

// GetByID retrieves an account by its ID
func (s *service) GetByID(context context.Context, id int) (*domain.Account, error) {
	return s.repo.GetByID(context, id)
}
