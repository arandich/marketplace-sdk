-- name: GetOrderByExternalId :one
SELECT external_id, client_id, order_id, status FROM orders
WHERE external_id = $1;

-- name: CreateOrder :one
INSERT INTO orders (
    external_id, client_id, order_id, status
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT DO NOTHING RETURNING status;

-- name: UpdateOrderStatusAndSetOrderIDByExternalId :exec
UPDATE orders
SET status = $1, order_id = $2
WHERE external_id = $3;