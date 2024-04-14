-- name: CreateGood :exec
INSERT INTO goods (
     id, user_id, name, description, price, status, image_url)
VALUES (
     $1, $2, $3, $4, $5, $6, $7
);

-- name: UpdateGood :exec
UPDATE goods
SET
    name = $2,
    description = $3,
    price = $4,
    status = $5,
    image_url = $6
WHERE id = $1;

-- name: DeleteGood :exec
DELETE FROM goods
WHERE id = $1;

-- name: GetGood :one
SELECT id, user_id, name, description, price, status, image_url, created_at, updated_at FROM goods
WHERE id = $1
LIMIT 1;

-- name: ListGoods :many
SELECT id, user_id, name, description, price, status, image_url, created_at, updated_at FROM goods
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;


