// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package id

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Hold struct {
	HoldID    int64
	Amount    int64
	UserID    string
	Status    string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type User struct {
	UserID    string
	Balance   int64
	Username  string
	Password  string
	CreatedAt pgtype.Timestamp
}
