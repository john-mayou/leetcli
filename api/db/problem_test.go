package db_test

import (
	"testing"

	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/require"
)

func TestProblemCRUD(t *testing.T) {
	client := testutil.SetupTestClient(t)

	problem := testutil.FakeProblem()

	// === Create ===
	problem, err := client.CreateProblem(problem)
	require.NoError(t, err)

	// === Find ===
	foundProblem, err := client.FindProblemByID(problem.ID)
	require.NoError(t, err)
	assertProblemEqual(t, problem, foundProblem)
	problem = foundProblem

	// === Update ===
	problem.Slug = "newslug"
	require.NoError(t, client.UpdateProblem(problem))
	updatedProblem, err := client.FindProblemByID(problem.ID)
	require.NoError(t, err)
	require.Equal(t, "newslug", updatedProblem.Slug)

	// === Delete ===
	require.NoError(t, client.DeleteProblem(problem.ID))
	deletedProblem, err := client.FindProblemByID(problem.ID)
	require.NoError(t, err)
	require.False(t, deletedProblem.DeletedAt.IsZero())
}

func assertProblemEqual(t *testing.T, want, got *model.Problem) {
	t.Helper()

	require.Equal(t, want.ID, got.ID)
	require.Equal(t, want.Slug, got.Slug)
}
