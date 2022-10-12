package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type NewOrderRequest struct {
	CustomerName string `json:"customer_name" valid:"required~customer_name cannot be empty"`
}

type EmbeddedItemsRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type NewOrderItemsRequest struct {
	CustomerName string                 `json:"customer_name" valid:"required~customer_name cannot be empty"`
	Items        []EmbeddedItemsRequest `json:"items" valid:"required~items cannot be empty"`
}

type EmbeddedItemsResponse struct {
	ID          int    `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type OrderItemsResponse struct {
	ID           int                     `json:"order_id"`
	CustomerName string                  `json:"customer_name"`
	OrderedAt    time.Time               `json:"ordered_at"`
	Items        []EmbeddedItemsResponse `json:"items"`
}

type NewOrderResponse struct {
	ID           int                    `json:"id"`
	CustomerName string                 `json:"customer_name"`
	Ordered_at   time.Time              `json:"ordered_at"`
	Items        []EmbeddedItemsRequest `json:"items"`
}

func (o *NewOrderRequest) Validate() error {
	_, err := govalidator.ValidateStruct(o)
	if err != nil {
		return err
	}
	return nil
}

func (o *NewOrderItemsRequest) Validate() error {
	_, err := govalidator.ValidateStruct(o)
	if err != nil {
		return err
	}
	return nil
}
