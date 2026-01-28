package grpc

import (
	"boilerplate-go/internal/infra/grpc/proto/health"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HealthServer struct {
	health.UnimplementedHealthServer
}

func (s *HealthServer) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "ok"}, nil
}

type GRPCService struct {
	server *grpc.Server
	port   string
}

func NewGRPCService(port string) *GRPCService {
	s := grpc.NewServer()
	health.RegisterHealthServer(s, &HealthServer{})
	reflection.Register(s)

	return &GRPCService{
		server: s,
		port:   port,
	}
}

func (s *GRPCService) Start() error {
	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", s.port, err)
	}

	fmt.Printf("gRPC server running on port %s...\n", s.port)
	return s.server.Serve(lis)
}

func (s *GRPCService) Stop() {
	fmt.Println("Stopping gRPC server...")
	s.server.GracefulStop()
}
