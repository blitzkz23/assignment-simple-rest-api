package entity

import "time"

type Order struct {
	ID            uint      `json:"order_id"`
	Customer_Name string    `json:"customer_name"`
	Ordered_At    time.Time `json:"ordered_at"`
	Items         []Item    `json:"items"`
}
