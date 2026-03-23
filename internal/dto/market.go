package dto

import (
	"time"

	"SpotInstrumentService/internal/errs"
	"SpotInstrumentService/internal/model"

	pb "github.com/erdedan1/protocol/proto/spot_instrument_service/gen/v2"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Id:        vmr.ID.String(),
		Name:      vmr.Name,
		Enabled:   vmr.Enabled,
		CreatedAt: timestamppb.New(*vmr.CreatedAt),
		UpdatedAt: timestamppb.New(*vmr.UpdateAt),
		DeletedAt: timestamppb.New(*vmr.DeletedAt),
	}
}
