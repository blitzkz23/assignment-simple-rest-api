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
	insertItemQuery = `
		INSERT INTO items (item_code, description, quantity, order_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, item_code, description, quantity, order_id;
	`
	retrieveOrderItemQuery = `
		SELECT o.id, o.customer_name, o.ordered_at, i.id, i.item_code, i.description, i.quantity
		FROM orders as o
		LEFT JOIN items as i ON o.id = i.order_id
		ORDER BY o.id ASC;
	`
	updateOrderQuery = `
		UPDATE orders
		SET customer_name = $2
		WHERE id = $1
	`
	updateItemQuery = `
		UPDATE items
		SET item_code = $1, description = $2, quantity = $3
		WHERE item_code = $1
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

func (o *orderPG) CreateOrder(orderPayload *entity.Order) error {
	var order entity.Order
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	// ! Insert data order ke database
	row := tx.QueryRow(insertOrderQuery, orderPayload.CustomerName)
	err2 := row.Scan(&order.ID, &order.CustomerName, &order.OrderedAt)
	if err2 != nil {
		tx.Rollback()
		return err
	}

	// ! Insert data item ke database
	for _, value := range orderPayload.Items {
		_, err = tx.Exec(insertItemQuery, value.ItemCode, value.Description, value.Quantity, &order.ID)

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

func (o *orderPG) GetOrderItems() ([]*dto.OrderItemsResponse, error) {
	var orderItems = []*dto.OrderItemsResponse{}

	rows, err := o.db.Query(retrieveOrderItemQuery)
	if err != nil {
		return nil, err
	}

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

		if err != nil {
			return nil, err
		}

		dto := orderItem.ToOrderItemsDTO()
		orderItems = append(orderItems, &dto)
	}

	return orderItems, nil
}

func (o *orderPG) UpdateOrderItems(orderId int, orderPayload *entity.Order) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	// ! Update tabel order
	_, err2 := tx.Exec(updateOrderQuery, orderId, orderPayload.CustomerName)
	if err2 != nil {
		tx.Rollback()
		return err
	}

	// ! Update tabel item
	for _, value := range orderPayload.Items {
		_, err = tx.Exec(updateItemQuery, value.ItemCode, value.Description, value.Quantity)

		fmt.Println("current value update payload repository", value)
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
