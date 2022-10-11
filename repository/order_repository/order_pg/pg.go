package order_pg

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/entity"
	"assignment-simple-rest-api/repository/order_repository"
	"database/sql"
	"fmt"
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
	retrieveOrderItemQuery = `
		SELECT o.id, o.customer_name, o.ordered_at, i.id, i.item_code, i.description, i.quantity
		FROM orders as o
		LEFT JOIN items as i ON o.id = i.order_id
		ORDER BY o.id ASC;
	`
	updateOrderQuery = `

	`
	deleteOrderQuery = `
		DELETE FROM orders
		WHERE id = $1
	`
)

type orderPG struct {
	db *sql.DB
}

// ! Factory function yang mengembalikan orderPg dengan inject db.
func NewOrderPg(db *sql.DB) *orderPG {
	return &orderPG{db}
}

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

func (o *orderPG) CreateOrder(orderPayload *entity.Order, itemPayload []*entity.Item) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	for _, value := range itemPayload {
		_, err = tx.Exec(insertOrderQuery, orderPayload.CustomerName, value.ID, value.ItemCode, value.Description, value.Quantity)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
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

func (o *orderPG) GetOrderItems() ([]*dto.OrderItemsResponse, error) {
	var orderItems = []*dto.OrderItemsResponse{}

	rows, err := o.db.Query(retrieveOrderItemQuery)
	if err != nil {
		return nil, err
	}
	// fmt.Println("current row", &rows)

	for rows.Next() {
		var orderItem order_repository.OrderItems
		err = rows.Scan(
			&orderItem.Order.ID,
			&orderItem.Order.CustomerName,
			&orderItem.Order.OrderedAt,
			&orderItem.Items.ID,
			&orderItem.Items.ItemCode,
			&orderItem.Items.Description,
			&orderItem.Items.Quantity,
		)
		// fmt.Println("current orderItem", orderItem)

		if err != nil {
			return nil, err
		}

		dto := orderItem.ToOrderItemsDTO()
		fmt.Println("current dto", dto)
		orderItems = append(orderItems, &dto)
	}

	return orderItems, nil
}

func (o *orderPG) UpdateOrderById(orderID int) (*entity.Order, error) {
	var order entity.Order

	return &order, nil
}

func (o *orderPG) DeleteOrder(orderID int) (int64, error) {
	res, err := o.db.Exec(deleteOrderQuery, orderID)

	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}
