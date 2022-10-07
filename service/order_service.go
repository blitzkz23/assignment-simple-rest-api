package service

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
	"assignment-simple-rest-api/repository/order_repository"
)

type OrderService interface {
	InsertOrder(*dto.NewOrderRequest) (*dto.NewOrderResponse, error)
	GetAllOrder() ([]*entity.Order, error)
}

type orderService struct {
	repo order_repository.OrderRepository
}

func NewOrderService(repo order_repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (o *orderService) InsertOrder(orderPayload *dto.NewOrderRequest) (*dto.NewOrderResponse, error) {
	var orderResponse dto.NewOrderResponse

	return &orderResponse, nil
}

func (o *orderService) GetAllOrder() ([]*entity.Order, error) {
	var orders []*entity.Order

	return orders, nil
}
