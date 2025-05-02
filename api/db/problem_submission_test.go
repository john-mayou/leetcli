package db_test

import (
	"testing"

	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/require"
)

func TestProblemSubmissionCRUD(t *testing.T) {
	client := testutil.SetupTestClient(t)

	user, err := client.CreateUser(testutil.FakeUser())
	require.NoError(t, err)
	problem, err := client.CreateProblem(testutil.FakeProblem())
	require.NoError(t, err)
	ps := testutil.FakeProblemSubmission(problem.ID, user.ID)

	// === Create ===
	ps, err = client.CreateProblemSubmission(ps)
	require.NoError(t, err)

	// === Find ===
	foundPs, err := client.FindProblemSubmissionByID(ps.ID)
	require.NoError(t, err)
	assertProblemSubmissionEqual(t, ps, foundPs)
	ps = foundPs

	// === Update ===
	ps.Code = "newcode"
	require.NoError(t, client.UpdateProblemSubmission(ps))
	updatedPs, err := client.FindProblemSubmissionByID(ps.ID)
	require.NoError(t, err)
	require.Equal(t, "newcode", updatedPs.Code)

	// === Delete ===
	require.NoError(t, client.DeleteProblemSubmission(ps.ID))
	deletedPs, err := client.FindProblemSubmissionByID(ps.ID)
	require.NoError(t, err)
	require.False(t, deletedPs.DeletedAt.IsZero())
}

func assertProblemSubmissionEqual(t *testing.T, want, got *model.ProblemSubmission) {
	t.Helper()

	require.Equal(t, want.ID, got.ID)
	require.Equal(t, want.ProblemID, got.ProblemID)
	require.Equal(t, want.UserID, got.UserID)
	require.Equal(t, want.Status, got.Status)
	require.Equal(t, want.Code, got.Code)
	require.Equal(t, want.Output, got.Output)
	require.Equal(t, want.ExecTimeMs, got.ExecTimeMs)
}
