package dto

import (
	"time"

	"SpotInstrumentService/internal/errs"
	"SpotInstrumentService/internal/model"

	pb "github.com/erdedan1/protocol/proto/spot_instrument_service/gen"
	m "github.com/erdedan1/shared/mapper"
	"github.com/google/uuid"
)

type ViewMarketsRequest struct {
	UserRoles []string
}

func NewViewMarketsRequest(req *pb.ViewMarketsRequest) (*ViewMarketsRequest, error) {
	if req == nil {
		return nil, errs.ErrInvalidArgument
	}

	return &ViewMarketsRequest{
		UserRoles: req.UserRoles,
	}, nil
}

type ViewMarketsResponse struct {
	ID        uuid.UUID
	Name      string
	Enabled   bool
	CreatedAt *time.Time
	UpdateAt  *time.Time
	DeletedAt *time.Time
}

func (vmr *ViewMarketsResponse) ModelToDto(market model.Market) *ViewMarketsResponse {
	vmr.ID = market.ID
	vmr.Name = market.Name
	vmr.Enabled = market.Enabled
	vmr.CreatedAt = market.CreatedAt
	vmr.UpdateAt = market.UpdateAt
	vmr.DeletedAt = market.DeletedAt
	return vmr
}

func (vmr *ViewMarketsResponse) DtoToProto() *pb.Market {
	return &pb.Market{
		Id:        m.ToUUIDProto(vmr.ID),
		Name:      vmr.Name,
		Enabled:   vmr.Enabled,
		CreatedAt: m.ToTimestampProto(vmr.CreatedAt),
		UpdatedAt: m.ToTimestampProto(vmr.UpdateAt),
		DeletedAt: m.ToTimestampProto(vmr.DeletedAt),
	}
}
