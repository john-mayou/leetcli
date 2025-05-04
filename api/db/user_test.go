package db_test

import (
	"testing"

	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/assert"
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

func TestFindUserByGithubID(t *testing.T) {
	client := testutil.SetupTestClient(t)

	user := testutil.FakeUser()
	user.GithubID = "123"

	// create
	user, err := client.CreateUser(user)
	require.NoError(t, err)

	// find
	foundUser, err := client.FindUserByGithubID("123")
	require.NoError(t, err)
	assertUserEqual(t, user, foundUser)
}

func TestListUsers(t *testing.T) {
	client := testutil.SetupTestClient(t)

	userA, err := client.CreateUser(testutil.FakeUser())
	require.NoError(t, err)
	userB, err := client.CreateUser(testutil.FakeUser())
	require.NoError(t, err)

	users, err := client.ListUsers()
	require.NoError(t, err)
	require.Len(t, users, 2)

	ids := []string{users[0].ID, users[1].ID}
	require.Contains(t, ids, userA.ID)
	require.Contains(t, ids, userB.ID)
}

func assertUserEqual(t *testing.T, expected, actual *model.User) {
	t.Helper()

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.GithubID, actual.GithubID)
	assert.Equal(t, expected.Username, actual.Username)
	assert.Equal(t, expected.Email, actual.Email)
}
