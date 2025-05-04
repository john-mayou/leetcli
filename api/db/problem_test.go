package db_test

import (
	"testing"

	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/assert"
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

func TestListProblems(t *testing.T) {
	client := testutil.SetupTestClient(t)

	problemA, err := client.CreateProblem(testutil.FakeProblem())
	require.NoError(t, err)
	problemB, err := client.CreateProblem(testutil.FakeProblem())
	require.NoError(t, err)

	problems, err := client.ListProblems()
	require.NoError(t, err)
	require.Len(t, problems, 2)

	ids := []string{problems[0].ID, problems[1].ID}
	require.Contains(t, ids, problemA.ID)
	require.Contains(t, ids, problemB.ID)
}

func assertProblemEqual(t *testing.T, expected, actual *model.Problem) {
	t.Helper()

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Slug, actual.Slug)
}
