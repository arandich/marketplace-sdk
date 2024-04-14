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

-- name: CreateOrder :exec
INSERT INTO orders
(id, user_id, good_ids, status)
VALUES ($1, $2, $3, $4);

-- name: UpdateOrder :exec
UPDATE orders
SET status = $2
WHERE id = $1;

-- name: GetOrder :one
SELECT
  id, user_id, good_ids, status, created_at, updated_at
FROM orders
WHERE id = $1;

-- name: GetOrdersByUser :many
SELECT
  id, user_id, good_ids, status, created_at, updated_at
FROM orders
WHERE user_id = $1;
