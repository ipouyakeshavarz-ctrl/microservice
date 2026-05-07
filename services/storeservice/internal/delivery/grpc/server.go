package grpc

import (
	"myapp/api/gen/store"
	"myapp/pkg/interceptor"
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
	engine    *grpc.Server
}

func NewServer(V storevalidator.Validator, s storeservice.Service, address string) *Server {
	return &Server{
		Validator: V,
		service:   s,
		address:   address,
		engine: grpc.NewServer(
			grpc.UnaryInterceptor(interceptor.UnaryErrorInterceptor())),
	}
}

func (s *Server) Run() error {
	const op = "grpc.server.run"

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	store.RegisterStoreServiceServer(s.engine, storehandler.New(s.service, s.Validator))

	return s.engine.Serve(lis)
}
func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
