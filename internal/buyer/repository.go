package buyer

import (
	"context"

	"github.com/agustincaro/gRPC/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {

	return []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "card1",
			FirstName:    "Carl",
			LastName:     "Cox",
		},
		{
			ID:           2,
			CardNumberID: "card2",
			FirstName:    "Charlotte",
			LastName:     "DeWitte",
		},
	}, nil
}
