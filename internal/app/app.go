package app

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"SpotInstrumentService/config"
	grpcHandler "SpotInstrumentService/internal/grpc"
	marketRepo "SpotInstrumentService/internal/repository/market"
	marketSrv "SpotInstrumentService/internal/service/market"
	"SpotInstrumentService/internal/usecase"

	pb "github.com/erdedan1/protocol/proto/spot_instrument_service/gen"
	"github.com/erdedan1/shared/interceptors/logger"
	"github.com/erdedan1/shared/interceptors/recovery"
	"github.com/erdedan1/shared/interceptors/request_id"
	log "github.com/erdedan1/shared/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	L          log.Logger
}

func New(cfg *config.Config) *App {

	logger, err := log.NewLogger(cfg.SpotInstrumentService.Log_LVL)
	if err != nil {
		panic(err)
	}

	return &App{
		cfg: cfg,
		L:   logger,
	}
}

const layer = "App"

func (a *App) Start() error {
	const method = "Start"
	a.L.Info(layer, method, "Start app")
	repos := usecase.NewRepositories(marketRepo.NewRepo(a.L))
	srvs := usecase.NewServices(marketSrv.NewService(*repos, a.L))

	go func() {
		if err := a.startGRPCServer(*srvs); err != nil {
			a.L.Error(layer, method, "failed to start grpc server", err)
			os.Exit(1)
		}
	}()
	a.L.Info(layer, method, "service work")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	a.L.Info(layer, method, "waiting for shutdown signal")
	<-quit
	a.L.Info(layer, method, "shutdown signal received")
	a.stopGRPCServer()
	a.L.Info(layer, method, "service stopped gracefully")
	return nil
}

func (a *App) startGRPCServer(usecase usecase.Services) error {
	const method = "startGRPCServer"
	logZap, _ := zap.NewProduction()
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			request_id.XRequestIDServerInterceptor(),
			logger.LoggerServerInterceptor(logZap),
			recovery.RecoveryServerInterceptor(logZap),
		),
	)

	a.grpcServer = grpcServer
	grpcHandler := grpcHandler.NewService(usecase)
	pb.RegisterMarketServiceServer(grpcServer, grpcHandler)

	lis, err := net.Listen("tcp", a.cfg.GRPCServer.Address)
	if err != nil {
		a.L.Error(layer, method, "failed to listenv", err)
		return err
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		a.L.Error(layer, method, "grpc serve error", err)
		return err
	}
	return nil
}

func (a *App) stopGRPCServer() {
	const method = "stopGRPCServer"
	a.grpcServer.GracefulStop()
	a.L.Info(layer, method, "grpc server stopped gracefully")
}
