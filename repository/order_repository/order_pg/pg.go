package order_pg

import (
	"assignment-simple-rest-api/entity"
	"database/sql"
)

type orderPG struct {
	db *sql.DB
}

// * Factory function yang mengembalikan orderPg dengan inject db.
func NewOrderPg(db *sql.DB) *orderPG {
	return &orderPG{db}
}

// * Implement functions from repository interface.
func (o *orderPG) InsertOrder(orderPayload *entity.Order) (*entity.Order, error) {
	var order entity.Order

	return &order, nil
}

func (o *orderPG) GetAllOrder() ([]*entity.Order, error) {
	var orders []*entity.Order

	return orders, nil
}

func (o *orderPG) UpdateOrderById(orderID int) (*entity.Order, error) {
	var order entity.Order

	return &order, nil
}

func (o *orderPG) DeleteOrder(orderID int) error {
	return error(nil)
}
