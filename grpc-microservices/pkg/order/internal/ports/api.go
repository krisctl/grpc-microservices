package ports

import "github.com/krisctl/grpc-microservices/pkg/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
