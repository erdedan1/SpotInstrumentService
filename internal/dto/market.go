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
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (vmr *ViewMarketsResponse) ModelToDto(market model.Market) *ViewMarketsResponse {
	vmr.ID = market.ID
	vmr.Name = market.Name
	vmr.Enabled = market.Enabled
	vmr.CreatedAt = market.CreatedAt
	vmr.UpdatedAt = market.UpdateAt
	vmr.DeletedAt = market.DeletedAt
	return vmr
}

func (vmr *ViewMarketsResponse) DtoToProto() *pb.Market {

	market := &pb.Market{
		Id:      vmr.ID.String(),
		Name:    vmr.Name,
		Enabled: vmr.Enabled,
	}

	if vmr.CreatedAt != nil {
		market.CreatedAt = timestamppb.New(*vmr.CreatedAt)
	}

	if vmr.UpdatedAt != nil {
		market.UpdatedAt = timestamppb.New(*vmr.UpdatedAt)
	}

	if vmr.DeletedAt != nil {
		market.DeletedAt = timestamppb.New(*vmr.DeletedAt)
	}

	return market
}
