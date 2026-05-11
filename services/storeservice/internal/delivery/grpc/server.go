package grpc

import (
	"myapp/api/gen/store"
	"myapp/pkg/interceptor"
	"myapp/pkg/metrics"
	"myapp/pkg/richerror"
	"net"
	"storeapp/internal/delivery/grpc/storehandler"
	"storeapp/internal/service"
	storevalidator "storeapp/internal/validator"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type Server struct {
	Validator   storevalidator.Validator
	service     storeservice.Service
	address     string
	engine      *grpc.Server
	metricsPort int
}

func NewServer(V storevalidator.Validator, s storeservice.Service, address string, metricsPort int) *Server {
	engine := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcprometheus.UnaryServerInterceptor,
			interceptor.UnaryErrorInterceptor(),
		),
	)

	grpcprometheus.EnableHandlingTimeHistogram()
	grpcprometheus.Register(engine)

	return &Server{
		Validator:   V,
		service:     s,
		address:     address,
		engine:      engine,
		metricsPort: metricsPort,
	}
}

func (s *Server) Run() error {
	const op = "grpc.server.run"

	go metrics.StartServer(s.metricsPort)

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	store.RegisterStoreServiceServer(s.engine, storehandler.New(s.service, s.Validator))

	return s.engine.Serve(lis)
}
func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
