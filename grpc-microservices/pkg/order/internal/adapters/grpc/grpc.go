package grpc

import (
	"context"

	"github.com/krisctl/grpc-microservices/golang/order"
	"github.com/krisctl/grpc-microservices/pkg/order/internal/application/core/domain"
)

func (ga GrpcAdapter) Create(ctx context.Context, createOrderReq *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range createOrderReq.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			Quantity:    orderItem.Quantity,
			UnitPrice:   orderItem.UnitPrice,
		})
	}
	newOrder := domain.NewOrder(createOrderReq.UserId, orderItems)
	result, err := ga.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
