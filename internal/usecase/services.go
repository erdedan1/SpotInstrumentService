package usecase

import (
	"SpotInstrumentService/internal/dto"
	"context"
)

type MarketService interface {
	ViewMarketsByRoles(ctx context.Context, req *dto.ViewMarketsRequest) ([]dto.ViewMarketsResponse, error)
}

type Services struct {
	MarketService MarketService
}

func NewServices(market MarketService) *Services {
	return &Services{
		MarketService: market,
	}
}
