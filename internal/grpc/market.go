package grpc

import (
	"SpotInstrumentService/internal/dto"
	"SpotInstrumentService/internal/errs"
	"SpotInstrumentService/internal/usecase"
	"context"

	pb "github.com/erdedan1/protocol/proto/spot_instrument_service/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	srv usecase.Services
	pb.MarketServiceServer
}

func NewService(srv usecase.Services) *Service {
	return &Service{
		srv: srv,
	}
}

func (s *Service) ViewMarketsByRoles(ctx context.Context, req *pb.ViewMarketsRequest) (*pb.ViewMarketsResponse, error) {
	dtoReq, err := dto.NewViewMarketsRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	markets, err := s.srv.MarketService.ViewMarketsByRoles(ctx, dtoReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(markets) == 0 {
		return nil, status.Error(codes.NotFound, errs.ErrNotFound.Error())
	}

	marketsProto := make([]*pb.Market, 0, len(markets))
	for _, m := range markets {
		marketsProto = append(marketsProto, m.DtoToProto())
	}

	return &pb.ViewMarketsResponse{
		Markets: marketsProto,
	}, nil
}
