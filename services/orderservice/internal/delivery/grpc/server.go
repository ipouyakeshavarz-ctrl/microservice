package grpc

import (
	"myapp/api/gen/order"
	"myapp/pkg/interceptor"
	"myapp/pkg/metrics"
	"net"
	orderhandler "orderapp/internal/delivery/grpc/producthandler"
	orderservice "orderapp/internal/service"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type Server struct {
	service     orderservice.Service
	address     string
	engine      *grpc.Server
	metricsPort int
}

func NewServer(s *orderservice.Service, address string, metricsPort int) *Server {
	engine := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcprometheus.UnaryServerInterceptor,
			interceptor.UnaryErrorInterceptor(),
		),
	)

	grpcprometheus.EnableHandlingTimeHistogram()
	grpcprometheus.Register(engine)

	return &Server{
		service:     *s,
		address:     address,
		engine:      engine,
		metricsPort: metricsPort,
	}
}

func (s *Server) Run() error {
	go metrics.StartServer(s.metricsPort)

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	order.RegisterOrderServiceServer(s.engine, orderhandler.New(s.service))

	return s.engine.Serve(lis)
}

func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
