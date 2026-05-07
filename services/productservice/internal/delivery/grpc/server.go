package grpc

import (
	"myapp/api/gen/product"
	"myapp/pkg/interceptor"
	"net"
	"productapp/internal/delivery/grpc/producthandler"
	productservice "productapp/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	service productservice.Service
	address string
	engine  *grpc.Server
}

func NewServer(s productservice.Service, address string) *Server {
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

	product.RegisterProductServiceServer(s.engine, producthandler.New(s.service))

	return s.engine.Serve(lis)
}

func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
