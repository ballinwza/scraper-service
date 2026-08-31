package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ballinwza/scraper-service/internal/delivery/gen_grpc"
)

// (Optional) นิยาม Interface สำหรับ External Dependencies เพื่อทำ Mock ใน Unit Test
type Logger interface {
	Printf(format string, v ...any)
}

// GreeterServer ทำหน้าที่เป็น Service Layer และ Implement gRPC Server Interface
type GreeterServer struct {
	pb.UnimplementedGreeterServiceServer
	logger Logger // Dependency ที่ถูก Inject เข้ามา
}

// NewGreeterServer คือ Constructor Function สำหรับทำ Dependency Injection
func NewGreeterServer(logger Logger) *GreeterServer {
	return &GreeterServer{
		logger: logger,
	}
}

// Implement SayHello RPC Method
func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.logger.Printf("Received request for name: %s", req.GetName())

	replyMessage := fmt.Sprintf("Hello %s! Executed via Service Layer.", req.GetName())

	return &pb.HelloResponse{
		Message:   replyMessage,
		Timestamp: time.Now().Unix(),
	}, nil
}
