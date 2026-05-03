package grpc

import (
	"fmt"
	"log"
	"myapp/api/gen/store"
	"net"
	"storeapp/internal/delivery/grpc/storehandler"
	"storeapp/internal/service"
	storevalidator "storeapp/internal/validator"

	"google.golang.org/grpc"
)

type Server struct {
	Validator storevalidator.Validator
	service   storeservice.Service
	port      int
}

func NewServer(V storevalidator.Validator, s storeservice.Service, port int) *Server {
	return &Server{
		Validator: V,
		service:   s,
		port:      port,
	}
}

func (s *Server) Run() error {

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	store.RegisterStoreServiceServer(grpcServer, storehandler.New(s.service, s.Validator))

	log.Println("gRPC server started on port", s.port)

	return grpcServer.Serve(lis)
}
