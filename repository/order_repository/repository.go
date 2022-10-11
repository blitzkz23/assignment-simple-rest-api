package order_repository

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
)

type OrderRepository interface {
	InsertOrder(orderPayload *entity.Order) (*entity.Order, error)
	CreateOrder(orderPayload *entity.Order, itemPayLoad []*entity.Item) error
	GetAllOrders() ([]*entity.Order, error)
	GetOrderItems() ([]*dto.OrderItemsResponse, error)
	UpdateOrderById(orderID int) (*entity.Order, error)
	DeleteOrder(orderID int) (int64, error)
}
