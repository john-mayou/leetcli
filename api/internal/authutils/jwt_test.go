package authutils_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/john-mayou/leetcli/config"
	"github.com/john-mayou/leetcli/internal/authutils"
	"github.com/stretchr/testify/require"
)

func TestGenerateJWT(t *testing.T) {
	now := time.Unix(1000, 0)
	cfg, err := config.LoadConfig()
	cfg.JWTSecret = "testsecret"
	require.NoError(t, err)

	// generate
	jwt, err := authutils.GenerateJWT(cfg, now, "userid")
	require.NoError(t, err)

	// validate
	claims, err := authutils.ValidateJWT(cfg, now, jwt)
	require.NoError(t, err)

	require.Equal(t, "userid", claims.UserID)
	require.Equal(t, now.Add(30*24*time.Hour).Unix(), claims.RegisteredClaims.ExpiresAt.Unix())
}

func TestValidateJWT(t *testing.T) {
	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	now := time.Unix(1000, 0)

	cases := map[string]struct {
		claims      jwt.MapClaims
		secret      string
		expectError bool
	}{
		"valid token": {
			claims: jwt.MapClaims{
				"user_id": "userid",
				"exp":     now.Add(1 * time.Hour).Unix(),
			},
			secret:      "testsecret",
			expectError: false,
		},
		"expired token": {
			claims: jwt.MapClaims{
				"user_id": "userid",
				"exp":     now.Add(-1 * time.Hour).Unix(),
			},
			secret:      "testsecret",
			expectError: true,
		},
		"invalid signature": {
			claims: jwt.MapClaims{
				"user_id": "userid",
				"exp":     now.Add(1 * time.Hour).Unix(),
			},
			secret:      "invalidsecret",
			expectError: true,
		},
		"missing user_id": {
			claims: jwt.MapClaims{
				"exp": now.Add(1 * time.Hour).Unix(),
			},
			secret:      "testsecret",
			expectError: true,
		},
		"missing exp": {
			claims: jwt.MapClaims{
				"user_id": "userid",
			},
			secret:      "testsecret",
			expectError: true,
		},
	}

	for tcName, tc := range cases {
		t.Run(tcName, func(t *testing.T) {
			cfg.JWTSecret = "testsecret"
			token := generateTokenFromClaims(t, cfg, tc.claims)

			cfg.JWTSecret = tc.secret
			claims, err := authutils.ValidateJWT(cfg, now, token)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NotNil(t, claims, "expected validate to return claims: %v", err)
				require.Equal(t, tc.claims["user_id"], claims.UserID)
				require.Equal(t, tc.claims["exp"], claims.RegisteredClaims.ExpiresAt.Unix())
			}
		})
	}
}

func generateTokenFromClaims(t *testing.T, cfg *config.Config, claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(cfg.JWTSecret))
	require.NoError(t, err)
	return tokenStr
}
