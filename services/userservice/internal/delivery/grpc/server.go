package grpc

import (
	"log"
	"myapp/pkg/richerror"
	"net"
	"userapp/internal/delivery/grpc/userhandler"
	userservice "userapp/internal/service"
	"userapp/internal/validator"

	"myapp/api/gen/user"

	"google.golang.org/grpc"
)

type Server struct {
	Validator validator.Validator
	service   userservice.Service
	address   string
}

func NewServer(V validator.Validator, s userservice.Service, address string) *Server {
	return &Server{
		Validator: V,
		service:   s,
		address:   address,
	}
}

func (s *Server) Run() error {
	const op = "grpc.server.Run"

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, userhandler.New(s.service, s.Validator))

	log.Println("🚀gRPC server started on ", s.address)

	return grpcServer.Serve(lis)
}
