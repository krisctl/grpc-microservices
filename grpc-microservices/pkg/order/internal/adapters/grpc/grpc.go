package grpc

import (
	"github.com/krisctl/grpc-microservices/golang/order"
	"github.com/krisctl/grpc-microservices/pkg/order/internal/ports"
)

type GrpcAdapter struct {
	api ports.APIPort
	order.UnimplementedOrderServer
}
