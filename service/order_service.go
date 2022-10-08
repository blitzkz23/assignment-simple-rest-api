package service

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
	"assignment-simple-rest-api/repository/order_repository"
)

// ! Order Service Interface
type OrderService interface {
	InsertOrder(*dto.NewOrderRequest) (*dto.NewOrderResponse, error)
	GetAllOrders() ([]*entity.Order, error)
}

// ! Order Service Implementation
type orderService struct {
	repo order_repository.OrderRepository
}

// ! Factory function yang mengembalikan orderPg dengan inject repo.
func NewOrderService(repo order_repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (o *orderService) InsertOrder(orderPayload *dto.NewOrderRequest) (*dto.NewOrderResponse, error) {
	// ! Service untuk insert data order ke database
	if err := orderPayload.Validate(); err != nil {
		return nil, err
	}

	orderRequest := &entity.Order{
		CustomerName: orderPayload.CustomerName,
	}

	newOrder, err := o.repo.InsertOrder(orderRequest)
	if err != nil {
		return nil, err
	}

	return newOrder.NewOrderResponseDTO(), nil
}

func (o *orderService) GetAllOrders() ([]*entity.Order, error) {
	// ! Service untuk mengambil data order dari repository
	orders, err := o.repo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
