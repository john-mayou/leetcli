package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/john-mayou/leetcli/handler"
	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentUser(t *testing.T) {
	h := handler.NewTestHandler(&handler.HandlerOpts{DBClient: testutil.SetupTestClient(t)})
	user := testutil.FakeUser()
	user, err := h.DBClient.CreateUser(user)
	require.NoError(t, err)

	t.Run("user found", func(t *testing.T) {
		ctx := handler.CtxWithUserID(context.Background(), user.ID)
		req := httptest.NewRequestWithContext(ctx, "GET", "/users/me", nil)
		r := httptest.NewRecorder()

		h.GetCurrentUser(r, req)

		require.Equal(t, http.StatusOK, r.Code)
		require.Contains(t, r.Body.String(), `"id":"`+user.ID+`"`)
	})
	t.Run("user not found", func(t *testing.T) {
		ctx := handler.CtxWithUserID(context.Background(), "randomid")
		req := httptest.NewRequestWithContext(ctx, "GET", "/users/me", nil)
		r := httptest.NewRecorder()

		h.GetCurrentUser(r, req)

		require.Equal(t, http.StatusNotFound, r.Code)
	})
}
