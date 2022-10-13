package service

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
	"assignment-simple-rest-api/repository/order_repository"
)

// ! Order Service Interface
type OrderService interface {
	InsertOrderItems(orderPayload *dto.NewOrderItemsRequest) (*dto.NewOrderItemsRequest, error)
	GetAllOrderItems() ([]*dto.OrderItemsResponse, error)
	UpdateOrderItems(orderId int, orderPayload *dto.NewOrderItemsRequest) (*dto.NewOrderItemsRequest, error)
	DeleteOrderByID(orderID int) (int64, error)
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

func (o *orderService) InsertOrderItems(orderPayload *dto.NewOrderItemsRequest) (*dto.NewOrderItemsRequest, error) {
	// ! Service untuk insert data order dan items ke database
	if err := orderPayload.Validate(); err != nil {
		return nil, err
	}

	orderRequest := &entity.Order{
		CustomerName: orderPayload.CustomerName,
	}

	for _, item := range orderPayload.Items {
		itemRequest := &entity.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		orderRequest.Items = append(orderRequest.Items, *itemRequest)
	}

	err := o.repo.CreateOrder(orderRequest)
	if err != nil {
		return nil, err
	}

	return orderPayload, nil
}

func (o *orderService) GetAllOrderItems() ([]*dto.OrderItemsResponse, error) {
	// ! Service untuk mengambil data order items dari repository
	orderItems, err := o.repo.GetOrderItems()
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (o *orderService) UpdateOrderItems(orderId int, orderPayload *dto.NewOrderItemsRequest) (*dto.NewOrderItemsRequest, error) {
	// ! Service untuk update data order dan items ke database
	if err := orderPayload.Validate(); err != nil {
		return nil, err
	}

	orderRequest := &entity.Order{
		CustomerName: orderPayload.CustomerName,
	}

	for _, item := range orderPayload.Items {
		itemRequest := &entity.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		orderRequest.Items = append(orderRequest.Items, *itemRequest)
	}

	err := o.repo.UpdateOrderItems(orderId, orderRequest)
	if err != nil {
		return nil, err
	}

	return orderPayload, nil
}

func (o *orderService) DeleteOrderByID(orderID int) (int64, error) {
	deletedOrder, err := o.repo.DeleteOrder(orderID)
	if err != nil {
		return 0, err
	}

	return deletedOrder, nil
}
