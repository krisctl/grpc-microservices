package db

import (
	"fmt"

	"github.com/krisctl/grpc-microservices/order/internal/application/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// It's odd that I have to duplicate the similar models here and in domain,order
// but looks like it is good for separation of concerns and for not doing ORM
// pollution in my core business logic
type OrderItem struct {
	gorm.Model
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
	OrderId     int32   `json:"order_id"` // back reference to Order model
}

type Order struct {
	gorm.Model
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"` // Reference to OrderItem model
}

type DbAdapter struct {
	db *gorm.DB // how do I configure this to save data to a sqlite table?
}

func NewDbAdapter(dataSourceUrl string) (*DbAdapter, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &DbAdapter{db: db}, nil
}

func (dba DbAdapter) Get(id string) (domain.Order, error) {
	var orderEntity Order
	res := dba.db.First(&orderEntity, id) // Puts the record from Db into orderEntity
	if res.Error != nil {
		return domain.Order{}, res.Error
	}
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity})
	}
	// Converting from gorm model to domain model
	order := domain.Order{
		ID:         int64(orderEntity.CustomerID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano()}
	return order, nil
}

func (dba DbAdapter) Save(domain.Order) error {
	// implementation here
	return nil
}
