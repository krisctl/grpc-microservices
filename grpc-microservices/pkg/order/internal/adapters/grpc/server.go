package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/krisctl/grpc-microservices/golang/order"
	"github.com/krisctl/grpc-microservices/pkg/order/config"
	"github.com/krisctl/grpc-microservices/pkg/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcAdapter struct {
	// for what purpose do we depend on APIPort interface? - To call PlaceOrder function
	api ports.APIPort
	order.UnimplementedOrderServer
	// who is responsible for injecting this port on which grpc server is started - main.go
	port int
}

func NewGrpcAdapter(api ports.APIPort, port int) *GrpcAdapter {
	return &GrpcAdapter{
		api:  api,
		port: port,
	}
}

func (ga GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", ga.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", ga.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, ga)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on assigned port")
	}
}
