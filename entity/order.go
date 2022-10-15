package entity

import (
	"assignment-simple-rest-api/dto"
	"time"
)

type Order struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}

func (o *Order) NewOrderResponseDTO() *dto.NewOrderResponse {
	return &dto.NewOrderResponse{
		ID:           o.ID,
		CustomerName: o.CustomerName,
		Ordered_at:   o.OrderedAt,
	}
}
