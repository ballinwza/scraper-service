package di

import (
	"log"
	"os"

	pb "github.com/ballinwza/scraper-service/internal/delivery/gen_grpc"
	"github.com/ballinwza/scraper-service/internal/delivery/service"
	"google.golang.org/grpc"
)

func InjectionGRPC(grpcServer *grpc.Server) {
	customLogger := log.New(os.Stdout, "[gRPC-Server] ", log.LstdFlags)

	greeterService := service.NewGreeterServer(customLogger)

	pb.RegisterGreeterServiceServer(grpcServer, greeterService)
}
