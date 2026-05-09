package grpc

import (
	authhandler "authapp/internal/delivery/grpc/authhamdler"
	authservice "authapp/internal/service"
	"fmt"
	"myapp/api/gen/auth"
	"net"
	"net/http"

	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	service     authservice.Service
	address     string
	engine      *grpc.Server
	metricsPort int
}

func NewServer(s *authservice.Service, address string, MetricsPort int) *Server {

	srv := &Server{
		service: *s,
		address: address,
		engine: grpc.NewServer(
			grpc.UnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
			grpc.StreamInterceptor(grpcprometheus.StreamServerInterceptor),
		),
		metricsPort: MetricsPort,
	}

	grpcprometheus.EnableHandlingTimeHistogram()

	return srv
}

func (s *Server) Run() error {

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "ok")
		})
		_ = http.ListenAndServe(fmt.Sprintf("authservice:%d", s.metricsPort), mux)
	}()

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s.engine, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	auth.RegisterAuthServiceServer(s.engine, authhandler.New(&s.service))

	grpcprometheus.Register(s.engine)

	return s.engine.Serve(lis)
}

func (s *Server) GracefulStop() {
	if s.engine != nil {
		s.engine.GracefulStop()
	}
}
