package order_pg

import (
	"assignment-simple-rest-api/entity"
	"database/sql"
)

const (
	// * Queries for order
	insertOrderQuery = `
		INSERT INTO orders (customer_name)
		VALUES ($1)
		RETURNING id, customer_name, ordered_at;
	`
	retrieveAllOrderQuery = `
		SELECT id, customer_name, ordered_at
		FROM orders
	`
)

type orderPG struct {
	db *sql.DB
}

// ! Factory function yang mengembalikan orderPg dengan inject db.
func NewOrderPg(db *sql.DB) *orderPG {
	return &orderPG{db}
}

// * Implement functions from repository interface.
func (o *orderPG) InsertOrder(orderPayload *entity.Order) (*entity.Order, error) {
	// ! Insert data order ke database
	var order entity.Order

	row := o.db.QueryRow(insertOrderQuery, orderPayload.CustomerName)
	err := row.Scan(&order.ID, &order.CustomerName, &order.OrderedAt)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *orderPG) GetAllOrders() ([]*entity.Order, error) {
	// ! Ambil semua data order dari database
	var orders []*entity.Order

	rows, err := o.db.Query(retrieveAllOrderQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order

		err := rows.Scan(&order.ID, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (o *orderPG) UpdateOrderById(orderID int) (*entity.Order, error) {
	var order entity.Order

	return &order, nil
}

func (o *orderPG) DeleteOrder(orderID int) error {
	return error(nil)
}
