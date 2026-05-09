package grpc

import (
	"myapp/pkg/interceptor"
	"myapp/pkg/richerror"
	"net"
	"userapp/internal/delivery/grpc/userhandler"
	userservice "userapp/internal/service"
	"userapp/internal/validator"

	"myapp/api/gen/user"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type Server struct {
	Validator validator.Validator
	service   userservice.Service
	address   string
	engine    *grpc.Server
}

func NewServer(V validator.Validator, s userservice.Service, address string) *Server {

	engine := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcprometheus.UnaryServerInterceptor,
			interceptor.UnaryErrorInterceptor(),
		),
	)

	grpcprometheus.Register(engine)

	return &Server{
		Validator: V,
		service:   s,
		address:   address,
		engine:    engine,
	}
}

func (s *Server) Run() error {
	const op = "grpc.server.Run"

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	user.RegisterUserServiceServer(s.engine, userhandler.New(s.service, s.Validator))

	return s.engine.Serve(lis)
}
func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
