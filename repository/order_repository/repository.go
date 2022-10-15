package order_repository

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
)

type OrderRepository interface {
	CreateOrder(orderPayload *entity.Order) error
	GetOrderItems() ([]*dto.OrderItemsResponse, error)
	UpdateOrderItems(orderId int, orderPayload *entity.Order) error
	DeleteOrder(orderID int) (int64, error)
}
