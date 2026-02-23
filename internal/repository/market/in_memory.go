package market

import (
	"context"
	"sync"
	"time"

	"SpotInstrumentService/internal/model"

	log "github.com/erdedan1/shared/logger"
	"github.com/google/uuid"
)

type InMemory struct {
	Markets map[uuid.UUID]model.Market
	mu      *sync.RWMutex
	l       log.Logger
}

func NewRepo(logger log.Logger) *InMemory {
	marketRepo := &InMemory{
		Markets: make(map[uuid.UUID]model.Market),
		mu:      &sync.RWMutex{},
		l:       logger.Layer("Repository"),
	}

	now := time.Now()
	markets := []model.Market{
		{
			ID:           uuid.MustParse("0179803e-06f0-4369-b94f-14e26ec190a1"),
			Name:         "BTC-USDT",
			Enabled:      true,
			AllowedRoles: []string{"USER_ROLE_TRAIDER"},
			DeletedAt:    nil,
		},
		{
			ID:           uuid.MustParse("0179803e-06f0-4369-b94f-14e26ec190a2"),
			Name:         "DOGE-USDT",
			Enabled:      true,
			AllowedRoles: []string{"USER_ROLE_ADMIN"},
			DeletedAt:    &now,
		},
		{
			ID:           uuid.MustParse("0179803e-06f0-4369-b94f-14e26ec190a3"),
			Name:         "ETH-USDT",
			Enabled:      false,
			AllowedRoles: []string{"USER_ROLE_TRAIDER"},
			DeletedAt:    nil,
		},
		{
			ID:           uuid.MustParse("0179803e-06f0-4369-b94f-14e26ec190a4"),
			Name:         "SOL-USDT",
			Enabled:      false,
			AllowedRoles: []string{"USER_ROLE_TRAIDER"},
			DeletedAt:    &now,
		},
	}

	for _, market := range markets {
		marketRepo.Markets[market.ID] = market
	}

	return marketRepo
}

func (r *InMemory) ViewMarketsByRoles(ctx context.Context, userRoles []string) ([]model.Market, error) {
	const method = "ViewMarketsByRoles"
	r.mu.RLock()
	defer r.mu.RUnlock()

	setRoles := make(map[string]struct{}, len(userRoles))
	for _, role := range userRoles {
		setRoles[role] = struct{}{}
	}
	var result []model.Market

	for _, m := range r.Markets {
		if !m.Enabled || m.DeletedAt != nil {
			continue
		}

		for _, ar := range m.AllowedRoles {
			if _, ok := setRoles[ar]; ok {
				result = append(result, m)
				break
			}
		}
	}

	r.l.Debug(method, "count_markets", len(result))

	return result, nil
}

func (r *InMemory) CreateMarket(ctx context.Context, id uuid.UUID, market model.Market) error {
	const method = "CreateMarket"
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Markets[id] = market
	r.l.Debug(method, "Market created", id)
	return nil
}
