package buyer

import (
	"context"
	"errors"

	"github.com/agustincaro/gRPC/internal/domain"
)

var (
	ErrNotFound = errors.New("buyer not found")
	ErrDb       = errors.New("no buyers in the database")
	ErrCid      = errors.New("the company id already exists")
	ErrServer   = errors.New("internal Server Error")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{
		repo: rep,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	buyers, err := s.repo.GetAll(ctx)
	return buyers, err
}
