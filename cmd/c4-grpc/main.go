package main

import (
	pb "github.com/FernandoCagale/c4-type/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hproto "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	app, err := SetupApplication()
	if err != nil {
		panic("Erro to start application")
	}

	lis, err := net.Listen("tcp", os.Getenv("GRPC"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotifyServer(s, app)

	// Server Health Check
	health := health.NewServer()
	health.SetServingStatus("", hproto.HealthCheckResponse_SERVING)
	hproto.RegisterHealthServer(s, health)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
