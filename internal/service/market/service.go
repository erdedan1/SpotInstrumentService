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
		l:    logger.Layer("Market.Service"),
	}
}

func (s *Service) ViewMarketsByRoles(ctx context.Context, req *dto.ViewMarketsRequest) ([]dto.ViewMarketsResponse, error) {
	const method = "ViewMarketsByRoles"
	markets, err := s.repo.InMemoryRepo.ViewMarketsByRoles(ctx, req.UserRoles)
	if err != nil {
		s.l.Error(method, "get markets repo", err)
		return nil, err
	}
	marketsResp := make([]dto.ViewMarketsResponse, 0, len(markets)) //todo может сделать сразу метод который делает это в dto?
	for _, m := range markets {                                     //todo добавить горутинки для быстроты?
		marketsResp = append(marketsResp, *(new(dto.ViewMarketsResponse).ModelToDto(m)))
	}
	return marketsResp, nil
}
