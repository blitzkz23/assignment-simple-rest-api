package order_repository

import "assignment-simple-rest-api/entity"

type OrderRepository interface {
	InsertOrder(orderPayload *entity.Order) (*entity.Order, error)
	GetAllOrder() ([]*entity.Order, error)
	UpdateOrderById(orderID int) (*entity.Order, error)
	DeleteOrder(orderID int) error
}
