// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: marketplace-id.sql

package id

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createHold = `-- name: CreateHold :exec
INSERT INTO holds (
  hold_id, amount, user_id, status
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT DO NOTHING
`

type CreateHoldParams struct {
	HoldID int64
	Amount int64
	UserID string
	Status string
}

func (q *Queries) CreateHold(ctx context.Context, arg CreateHoldParams) error {
	_, err := q.db.Exec(ctx, createHold,
		arg.HoldID,
		arg.Amount,
		arg.UserID,
		arg.Status,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
  user_id, balance, username, password
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT DO NOTHING
`

type CreateUserParams struct {
	UserID   string
	Balance  int64
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.UserID,
		arg.Balance,
		arg.Username,
		arg.Password,
	)
	return err
}

const getHold = `-- name: GetHold :one
SELECT hold_id, amount, user_id, status, created_at FROM holds
WHERE hold_id = $1 LIMIT 1
`

type GetHoldRow struct {
	HoldID    int64
	Amount    int64
	UserID    string
	Status    string
	CreatedAt pgtype.Timestamp
}

func (q *Queries) GetHold(ctx context.Context, holdID int64) (GetHoldRow, error) {
	row := q.db.QueryRow(ctx, getHold, holdID)
	var i GetHoldRow
	err := row.Scan(
		&i.HoldID,
		&i.Amount,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getHolds = `-- name: GetHolds :many
SELECT hold_id, amount, user_id, status, created_at FROM holds
WHERE user_id = $1
`

type GetHoldsRow struct {
	HoldID    int64
	Amount    int64
	UserID    string
	Status    string
	CreatedAt pgtype.Timestamp
}

func (q *Queries) GetHolds(ctx context.Context, userID string) ([]GetHoldsRow, error) {
	rows, err := q.db.Query(ctx, getHolds, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetHoldsRow
	for rows.Next() {
		var i GetHoldsRow
		if err := rows.Scan(
			&i.HoldID,
			&i.Amount,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
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

const getUser = `-- name: GetUser :one
SELECT user_id, balance, username, password, created_at FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Balance,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const updateHoldStatus = `-- name: UpdateHoldStatus :exec
UPDATE holds
SET status = $2
WHERE hold_id = $1
`

type UpdateHoldStatusParams struct {
	HoldID int64
	Status string
}

func (q *Queries) UpdateHoldStatus(ctx context.Context, arg UpdateHoldStatusParams) error {
	_, err := q.db.Exec(ctx, updateHoldStatus, arg.HoldID, arg.Status)
	return err
}

const updateUserBalance = `-- name: UpdateUserBalance :exec
UPDATE users
SET balance = $2
WHERE user_id = $1
`

type UpdateUserBalanceParams struct {
	UserID  string
	Balance int64
}

func (q *Queries) UpdateUserBalance(ctx context.Context, arg UpdateUserBalanceParams) error {
	_, err := q.db.Exec(ctx, updateUserBalance, arg.UserID, arg.Balance)
	return err
}
