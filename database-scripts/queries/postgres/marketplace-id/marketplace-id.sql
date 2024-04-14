-- name: CreateUser :exec
INSERT INTO users (
  user_id, balance, username, password
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT DO NOTHING;

-- name: GetUser :one
SELECT user_id, balance, username, password, created_at FROM users
WHERE user_id = $1 LIMIT 1;

-- name: UpdateUserBalance :exec
UPDATE users
SET balance = $2
WHERE user_id = $1;

-- name: CreateHold :exec
INSERT INTO holds (
  hold_id, amount, user_id, status
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT DO NOTHING;

-- name: GetHold :one
SELECT hold_id, amount, user_id, status, created_at FROM holds
WHERE hold_id = $1 LIMIT 1;

-- name: GetHolds :many
SELECT hold_id, amount, user_id, status, created_at FROM holds
WHERE user_id = $1;

-- name: UpdateHoldStatus :exec
UPDATE holds
SET status = $2
WHERE hold_id = $1;