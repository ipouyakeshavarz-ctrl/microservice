package grpc

import (
	"cartapp/internal/delivery/grpc/catrhandler"
	"cartapp/internal/service"
	"myapp/api/gen/cart"
	"myapp/pkg/interceptor"
	"myapp/pkg/metrics"
	"net"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type Server struct {
	service     cartservice.Service
	address     string
	engine      *grpc.Server
	metricsPort int
}

func NewServer(s cartservice.Service, address string, metricsPort int) *Server {

	engine := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcprometheus.UnaryServerInterceptor, // ← اضافه شد
			interceptor.UnaryErrorInterceptor(),
		),
	)

	grpcprometheus.EnableHandlingTimeHistogram()
	grpcprometheus.Register(engine)

	return &Server{
		service:     s,
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

	cart.RegisterCartServiceServer(s.engine, carthandler.New(s.service))

	return s.engine.Serve(lis)
}

func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
