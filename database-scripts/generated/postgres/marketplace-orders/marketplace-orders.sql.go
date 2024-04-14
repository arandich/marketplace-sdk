// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: marketplace-orders.sql

package orders

import (
	"context"
)

const createOrder = `-- name: CreateOrder :exec
/*
 CREATE TABLE IF NOT EXISTS orders (
  id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  good_ids varchar(255) [] NOT NULL,
  status varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
 */

INSERT INTO orders
(id, user_id, good_ids, status)
VALUES ($1, $2, $3, $4)
`

type CreateOrderParams struct {
	ID      string
	UserID  string
	GoodIds []string
	Status  string
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.Exec(ctx, createOrder,
		arg.ID,
		arg.UserID,
		arg.GoodIds,
		arg.Status,
	)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT
  id, user_id, good_ids, status, created_at, updated_at
FROM orders
WHERE id = $1
`

func (q *Queries) GetOrder(ctx context.Context, id string) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.GoodIds,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrdersByUser = `-- name: GetOrdersByUser :many
SELECT
  id, user_id, good_ids, status, created_at, updated_at
FROM orders
WHERE user_id = $1
`

func (q *Queries) GetOrdersByUser(ctx context.Context, userID string) ([]Order, error) {
	rows, err := q.db.Query(ctx, getOrdersByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.GoodIds,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE orders
SET status = $2
WHERE id = $1
`

type UpdateOrderParams struct {
	ID     string
	Status string
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.Exec(ctx, updateOrder, arg.ID, arg.Status)
	return err
}
