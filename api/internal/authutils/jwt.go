package authutils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/john-mayou/leetcli/config"
)

const MONTH = 30 * 24 * time.Hour

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(cfg *config.Config, now time.Time, userId string) (string, error) {
	claims := TokenClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(MONTH)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateJWT(cfg *config.Config, now time.Time, tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	}, jwt.WithTimeFunc(func() time.Time { return now }))
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claim structure: %v", token.Claims)
	}
	if claims.ExpiresAt == nil {
		return nil, errors.New("missing valid exp")
	}
	if claims.UserID == "" {
		return nil, errors.New("missing valid user_id")
	}

	return claims, nil
}
