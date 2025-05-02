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
	UserID     string
	Expiration time.Time
}

func GenerateJWT(cfg *config.Config, now time.Time, userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"expiration": now.Add(MONTH).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateJWT(cfg *config.Config, now time.Time, tokenStr string) (*TokenClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}

	// user_id
	userId, ok := claims["user_id"]
	if !ok {
		return nil, fmt.Errorf("expected to find 'user_id' in claims")
	}
	userIdString, ok := userId.(string)
	if !ok || userIdString == "" {
		return nil, fmt.Errorf("expected to find 'user_id' that is a non empty string")
	}

	// expiration
	expirationVal, ok := claims["expiration"]
	if !ok {
		return nil, fmt.Errorf("expected to find 'expiration' in claims: %v", claims)
	}
	expirationInt, ok := expirationVal.(float64)
	if !ok {
		return nil, fmt.Errorf("'expiration' claim is not a number: %v", expirationVal)
	}
	expirationUnix := int64(expirationInt)
	if expirationUnix < now.Unix() {
		return nil, errors.New("expired jwt token")
	}

	return &TokenClaims{
		UserID:     userIdString,
		Expiration: time.Unix(expirationUnix, 0),
	}, nil
}
