package grpc

import (
	"cartapp/internal/delivery/grpc/catrhandler"
	"cartapp/internal/service"
	"myapp/api/gen/cart"
	"myapp/pkg/interceptor"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	service cartservice.Service
	address string
	engine  *grpc.Server
}

func NewServer(s cartservice.Service, address string) *Server {
	return &Server{
		service: s,
		address: address,
		engine:  grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryErrorInterceptor())),
	}
}

func (s *Server) Run() error {

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
