package grpc

import (
	"fmt"
	"log"
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
	port      int
}

func NewServer(V validator.Validator, s userservice.Service, port int) *Server {
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

	user.RegisterUserServiceServer(grpcServer, userhandler.New(s.service, s.Validator))

	log.Println("gRPC server started on port", s.port)

	return grpcServer.Serve(lis)
}
