package grpc

import (
	authhandler "authapp/internal/delivery/grpc/authhamdler"
	authservice "authapp/internal/service"
	"myapp/api/gen/auth"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	service authservice.Service
	address string
	engine  *grpc.Server
}

func NewServer(s authservice.Service, address string) *Server {
	return &Server{
		service: s,
		address: address,
		engine:  grpc.NewServer(),
	}
}

func (s *Server) Run() error {

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	auth.RegisterAuthServiceServer(s.engine, authhandler.New(&s.service))

	return s.engine.Serve(lis)
}

func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
