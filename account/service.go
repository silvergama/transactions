package account

import "context"

// Service represents the business logic for account operations
type Service struct {
	repo Repository
}

// NewService creates a new Service instance with the provided repository
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// GetByID retrieves an account by its ID
func (s *Service) GetByID(context context.Context, id int) (*Account, error) {
	return s.repo.GetByID(context, id)
}

// Create creates a new account and returns its ID
func (s *Service) Create(ctx context.Context, account *Account) (int, error) {
	accountID, err := s.repo.Create(ctx, account)
	if err != nil {
		return 0, err
	}

	return accountID, nil
}
