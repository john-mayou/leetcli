package db_test

import (
	"testing"

	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/require"
)

func TestUserCRUD(t *testing.T) {
	client := testutil.SetupTestClient(t)

	user := testutil.FakeUser()

	// === Create ===
	user, err := client.CreateUser(user)
	require.NoError(t, err)

	// === Find ===
	foundUser, err := client.FindUserByID(user.ID)
	require.NoError(t, err)
	assertUserEqual(t, user, foundUser)
	user = foundUser

	// === Update ===
	user.Username = "updateduser"
	require.NoError(t, client.UpdateUser(user))
	updatedUser, err := client.FindUserByID(user.ID)
	require.NoError(t, err)
	require.Equal(t, "updateduser", updatedUser.Username)

	// === Delete ===
	require.NoError(t, client.DeleteUser(user.ID))
	deletedUser, err := client.FindUserByID(user.ID)
	require.NoError(t, err)
	require.False(t, deletedUser.DeletedAt.IsZero())
}

func assertUserEqual(t *testing.T, want, got *model.User) {
	t.Helper()

	require.Equal(t, want.ID, got.ID)
	require.Equal(t, want.GithubID, got.GithubID)
	require.Equal(t, want.Username, got.Username)
	require.Equal(t, want.Email, got.Email)
}
