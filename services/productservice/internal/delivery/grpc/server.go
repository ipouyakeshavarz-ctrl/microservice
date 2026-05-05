package grpc

import (
	"myapp/api/gen/product"
	"net"
	"productapp/internal/delivery/grpc/producthandler"
	productservice "productapp/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	service productservice.Service
	address string
}

func NewServer(s productservice.Service, address string) *Server {
	return &Server{
		service: s,
		address: address,
	}
}

func (s *Server) Run() error {

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	product.RegisterProductServiceServer(grpcServer, producthandler.New(s.service))

	return grpcServer.Serve(lis)
}
