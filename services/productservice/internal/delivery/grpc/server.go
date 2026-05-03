package grpc

import (
	"fmt"
	"log"
	"myapp/api/gen/product"
	"net"
	"productapp/internal/delivery/grpc/producthandler"
	productservice "productapp/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	service productservice.Service
	port    int
}

func NewServer(s productservice.Service, port int) *Server {
	return &Server{
		service: s,
		port:    port,
	}
}

func (s *Server) Run() error {

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	product.RegisterProductServiceServer(grpcServer, producthandler.New(s.service))

	log.Println("gRPC server started on port", s.port)

	return grpcServer.Serve(lis)
}
