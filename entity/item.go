package entity

type Item struct {
	ID          int    `json:"item_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_ID    int    `json:"order_id"`
}
