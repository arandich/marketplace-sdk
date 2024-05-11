// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: marketplace-api.sql

package orders

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    external_id, client_id, order_id, status
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT DO NOTHING RETURNING status
`

type CreateOrderParams struct {
	ExternalID string
	ClientID   string
	OrderID    int64
	Status     pgtype.Text
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, createOrder,
		arg.ExternalID,
		arg.ClientID,
		arg.OrderID,
		arg.Status,
	)
	var status pgtype.Text
	err := row.Scan(&status)
	return status, err
}

const getOrderByExternalId = `-- name: GetOrderByExternalId :one
SELECT external_id, client_id, order_id, status FROM orders
WHERE external_id = $1
`

type GetOrderByExternalIdRow struct {
	ExternalID string
	ClientID   string
	OrderID    int64
	Status     pgtype.Text
}

func (q *Queries) GetOrderByExternalId(ctx context.Context, externalID string) (GetOrderByExternalIdRow, error) {
	row := q.db.QueryRow(ctx, getOrderByExternalId, externalID)
	var i GetOrderByExternalIdRow
	err := row.Scan(
		&i.ExternalID,
		&i.ClientID,
		&i.OrderID,
		&i.Status,
	)
	return i, err
}

const updateOrderStatusAndSetOrderIDByExternalId = `-- name: UpdateOrderStatusAndSetOrderIDByExternalId :exec
UPDATE orders
SET status = $1, order_id = $2
WHERE external_id = $3
`

type UpdateOrderStatusAndSetOrderIDByExternalIdParams struct {
	Status     pgtype.Text
	OrderID    int64
	ExternalID string
}

func (q *Queries) UpdateOrderStatusAndSetOrderIDByExternalId(ctx context.Context, arg UpdateOrderStatusAndSetOrderIDByExternalIdParams) error {
	_, err := q.db.Exec(ctx, updateOrderStatusAndSetOrderIDByExternalId, arg.Status, arg.OrderID, arg.ExternalID)
	return err
}
