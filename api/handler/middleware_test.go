package handler_test

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/john-mayou/leetcli/config"
	"github.com/john-mayou/leetcli/handler"
	"github.com/john-mayou/leetcli/internal/authutils"
	"github.com/stretchr/testify/require"
)

func TestLoggingMiddleare(t *testing.T) {
	var buf bytes.Buffer
	handler := handler.NewTestHandler(&handler.HandlerOpts{Logger: log.New(&buf, "", 0)})

	server := handler.LoggingMiddlware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/test-path", nil))

	require.Contains(t, buf.String(), "GET /test-path")
}

func TestAuthMiddleare(t *testing.T) {
	cfg, err := config.LoadConfig()
	require.NoError(t, err)
	cfg.JWTSecret = "testsecret"
	now := time.Unix(1000, 0)

	cases := map[string]struct {
		setAuth        func(r *http.Request)
		expectedCode   int
		expectedUserID string
	}{
		"token valid": {
			setAuth: func(r *http.Request) {
				jwt, err := authutils.GenerateJWT(cfg, now, "userid")
				require.NoError(t, err)
				r.Header.Set("Authorization", "Bearer "+jwt)
			},
			expectedCode:   200,
			expectedUserID: "userid",
		},
		"token invalid": {
			setAuth: func(r *http.Request) {
				r.Header.Set("Authorization", "Bearer random")
			},
			expectedCode:   401,
			expectedUserID: "",
		},
		"authorization blank": {
			setAuth: func(r *http.Request) {
				r.Header.Set("Authorization", "")
			},
			expectedCode:   401,
			expectedUserID: "",
		},
		"authorization missing": {
			setAuth:        func(r *http.Request) {},
			expectedCode:   401,
			expectedUserID: "",
		},
	}

	for tcName, tc := range cases {
		t.Run(tcName, func(t *testing.T) {
			h := handler.NewTestHandler(&handler.HandlerOpts{Config: cfg, Now: func() time.Time { return now }})

			req := httptest.NewRequest("GET", "/test-path", nil)
			tc.setAuth(req)
			r := httptest.NewRecorder()

			var handlerCtx context.Context
			server := h.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				handlerCtx = r.Context()
			}))
			server.ServeHTTP(r, req)

			require.Equal(t, tc.expectedCode, r.Code)
			userID := ""
			if handlerCtx != nil {
				userID, _ = handlerCtx.Value(handler.UserIDKey()).(string)
			}
			require.Equal(t, tc.expectedUserID, userID)
		})
	}
}
