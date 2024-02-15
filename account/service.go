package account

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetByID(context context.Context, id int) (*Account, error) {
	return s.repo.GetByID(context, id)
}

func (s *Service) Create(ctx context.Context, account *Account) (int, error) {
	accountID, err := s.repo.Create(ctx, account)
	if err != nil {
		return 0, err
	}

	return accountID, nil
}
