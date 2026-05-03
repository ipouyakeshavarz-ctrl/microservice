package grpc

import (
	"log"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"
	"net"
	"storeapp/internal/delivery/grpc/storehandler"
	"storeapp/internal/service"
	storevalidator "storeapp/internal/validator"

	"google.golang.org/grpc"
)

type Server struct {
	Validator storevalidator.Validator
	service   storeservice.Service
	address   string
}

func NewServer(V storevalidator.Validator, s storeservice.Service, address string) *Server {
	return &Server{
		Validator: V,
		service:   s,
		address:   address,
	}
}

func (s *Server) Run() error {
	const op = "grpc.server.run"

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	grpcServer := grpc.NewServer()

	store.RegisterStoreServiceServer(grpcServer, storehandler.New(s.service, s.Validator))

	log.Println("🚀gRPC server started on", s.address)

	return grpcServer.Serve(lis)
}
