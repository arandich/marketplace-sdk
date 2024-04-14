//go:generate easyjson -all $GOFILE
package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID       string `json:"user_id"`
	jwt.RegisteredClaims
}
