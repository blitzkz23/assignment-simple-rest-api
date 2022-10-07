package entity

type Item struct {
	ID          uint   `json:"item_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_ID    uint   `json:"order_id"`
}
