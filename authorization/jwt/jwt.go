package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/arandich/marketplace-sdk/model"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

// Authorization.
var (
	jwtPrivateKey = []byte(`ce10f7a4ffaecb509598adb0acb75a86eb9975919ea2e50256589917bdfb3f5b`)
	jwtSubject    = "authorization"
)

func CreateJWTToken(ctx context.Context, claims model.Claims) (string, model.Claims, error) {
	if err := validateIssuer(claims.RegisteredClaims.Issuer); err != nil {
		return "", model.Claims{}, err
	}

	// IAT (Issued At).
	issuedAt := time.Now().Local()

	// EXP (Expires At).
	expiresAtt := issuedAt.Add(time.Hour * 24)

	// Set claims data.
	claims.RegisteredClaims.Subject = jwtSubject
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAtt)
	claims.RegisteredClaims.IssuedAt = jwt.NewNumericDate(issuedAt)

	// Create token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Sign token.
	jwtToken, err := token.SignedString(jwtPrivateKey)
	if err != nil {
		return "", model.Claims{}, err
	}

	if err = setAuthorizationCookie(ctx, jwtToken); err != nil {
		return "", model.Claims{}, err
	}

	return jwtToken, claims, nil
}

func ValidateJWTToken(ctx context.Context, jwtToken string) (model.Claims, error) {
	opts := []jwt.ParserOption{
		jwt.WithoutClaimsValidation(),
	}

	var claims model.Claims
	token, err := jwt.ParseWithClaims(jwtToken, &claims, keyFunc, opts...)
	if err != nil {
		return model.Claims{}, errors.New("wrong jwt token provided")
	}

	if !token.Valid {
		return model.Claims{}, errors.New("jwt token is invalid")
	}

	if err = ValidateJWTClaims(claims); err != nil {
		return model.Claims{}, err
	}

	return claims, nil
}

func ValidateJWTClaims(claims model.Claims) error {
	if err := validateIssuer(claims.RegisteredClaims.Issuer); err != nil {
		return err
	}

	if claims.RegisteredClaims.Subject != jwtSubject {
		return jwt.ErrTokenInvalidSubject
	}

	if claims.RegisteredClaims.ExpiresAt == nil ||
		claims.RegisteredClaims.ExpiresAt.Before(time.Now().Local()) {
		return jwt.ErrTokenExpired
	}

	return nil
}

func setAuthorizationCookie(ctx context.Context, newJwtToken string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("no ctx provided")
	}

	md.Set("Authorization", newJwtToken)

	if err := grpc.SetHeader(ctx, md); err != nil {
		return err
	}

	return nil
}

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected token signing algorithm: %s", token.Method.Alg())
	}
	return jwtPrivateKey, nil
}
