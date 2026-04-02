package market

import (
	"SpotInstrumentService/internal/dto"
	"SpotInstrumentService/internal/usecase"
	"context"

	log "github.com/erdedan1/shared/logger"
)

type Service struct {
	repo usecase.Repository
	l    log.Logger
}

func NewService(repo usecase.Repository, logger log.Logger) *Service {
	return &Service{
		repo: repo,
		l:    logger,
	}
}

const layer = "MarketService"

func (s *Service) ViewMarketsByRole(ctx context.Context, req *dto.ViewMarketsRequest) ([]dto.ViewMarketsResponse, error) {
	const method = "ViewMarketsByRole"
	markets, err := s.repo.InMemoryRepo.ViewMarketsByRole(ctx, req.UserRole)
	if err != nil {
		s.l.Error(layer, method, "get markets repo", err)
		return nil, err
	}
	marketsResp := make([]dto.ViewMarketsResponse, 0, len(markets)) //todo может сделать сразу метод который делает это в dto?
	for _, m := range markets {                                     //todo добавить горутинки для быстроты?
		marketsResp = append(marketsResp, *(new(dto.ViewMarketsResponse).ModelToDto(m)))
	}
	return marketsResp, nil
}
