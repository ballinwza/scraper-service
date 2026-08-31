package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ballinwza/scraper-service/config"
	"github.com/ballinwza/scraper-service/internal/delivery/di"
	"google.golang.org/grpc"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("⚙️ App initialized in [%s] mode", cfg.Environment)

	serverAddr := fmt.Sprintf(":%s", cfg.Port)

	// Start Service
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	di.InjectionGRPC(grpcServer)

	log.Printf("🚀 Server running on %s", serverAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
