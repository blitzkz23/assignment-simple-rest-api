package order_repository

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
)

type OrderItems struct {
	Order entity.Order
	Items entity.Item
}

func (o OrderItems) ToOrderItemsDTO() dto.OrderItemsResponse {
	orderItemTemp := dto.OrderItemsResponse{
		ID:           o.Order.ID,
		CustomerName: o.Order.CustomerName,
		OrderedAt:    o.Order.OrderedAt,
		Items: []dto.EmbeddedItems{
			{
				ID:          o.Items.ID,
				ItemCode:    o.Items.ItemCode,
				Description: o.Items.Description,
				Quantity:    o.Items.Quantity,
			},
		},
	}

	return orderItemTemp
}
