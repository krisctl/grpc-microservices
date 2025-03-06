package grpc

import (
	"github.com/krisctl/grpc-microservices/order/internal/ports"
	"github.com/myusername/grpc-microservices/golang/order"
)

type GrpcAdapter struct {
	api ports.APIPort
	order.UnimplementedOrderServer
}
