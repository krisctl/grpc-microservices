package ports

import "github.com/krisctl/grpc-microservices/order/internal/application/core/domain"

type DbPort interface {
	Get(id string) (domain.Order, error)
	Save(domain.Order) error
}
