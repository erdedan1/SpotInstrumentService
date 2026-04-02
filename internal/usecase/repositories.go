package usecase

import (
	"SpotInstrumentService/internal/model"
	"context"

	"github.com/google/uuid"
)

type InMemoryRepo interface {
	ViewMarketsByRole(ctx context.Context, userRole string) ([]model.Market, error)
	CreateMarket(ctx context.Context, id uuid.UUID, market model.Market) error
}

type Repository struct {
	InMemoryRepo InMemoryRepo
}

func NewRepositories(market InMemoryRepo) *Repository {
	return &Repository{
		InMemoryRepo: market,
	}
}
