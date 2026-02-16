package usecase

import (
	"SpotInstrumentService/internal/model"
	"context"

	"github.com/google/uuid"
)

type InMemoryRepo interface {
	ViewMarketsByRoles(ctx context.Context, userRoles []string) ([]model.Market, error)
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
