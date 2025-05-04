package db_test

import (
	"testing"

	"github.com/john-mayou/leetcli/db"
	"github.com/john-mayou/leetcli/internal/testutil"
	"github.com/john-mayou/leetcli/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProblemSubmissionCRUD(t *testing.T) {
	client := testutil.SetupTestClient(t)

	submission := buildFakeProblemSubmission(t, client)

	// === Create ===
	submission, err := client.CreateProblemSubmission(submission)
	require.NoError(t, err)

	// === Find ===
	foundSubmission, err := client.FindProblemSubmissionByID(submission.ID)
	require.NoError(t, err)
	assertProblemSubmissionEqual(t, submission, foundSubmission)
	submission = foundSubmission

	// === Update ===
	submission.Code = "newcode"
	require.NoError(t, client.UpdateProblemSubmission(submission))
	updatedSubmission, err := client.FindProblemSubmissionByID(submission.ID)
	require.NoError(t, err)
	require.Equal(t, "newcode", updatedSubmission.Code)

	// === Delete ===
	require.NoError(t, client.DeleteProblemSubmission(submission.ID))
	deletedSubmission, err := client.FindProblemSubmissionByID(submission.ID)
	require.NoError(t, err)
	require.False(t, deletedSubmission.DeletedAt.IsZero())
}

func TestListProblemSubmissions(t *testing.T) {
	client := testutil.SetupTestClient(t)

	submissionA, err := client.CreateProblemSubmission(buildFakeProblemSubmission(t, client))
	require.NoError(t, err)
	submissionB, err := client.CreateProblemSubmission(buildFakeProblemSubmission(t, client))
	require.NoError(t, err)

	submissions, err := client.ListProblemSubmissions()
	require.NoError(t, err)
	require.Len(t, submissions, 2)

	ids := []string{submissions[0].ID, submissions[1].ID}
	require.Contains(t, ids, submissionA.ID)
	require.Contains(t, ids, submissionB.ID)
}

func buildFakeProblemSubmission(t *testing.T, client *db.Client) *model.ProblemSubmission {
	t.Helper()

	user, err := client.CreateUser(testutil.FakeUser())
	require.NoError(t, err)
	problem, err := client.CreateProblem(testutil.FakeProblem())
	require.NoError(t, err)
	return testutil.FakeProblemSubmission(problem.ID, user.ID)
}

func assertProblemSubmissionEqual(t *testing.T, expected, actual *model.ProblemSubmission) {
	t.Helper()

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.ProblemID, actual.ProblemID)
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.Code, actual.Code)
	assert.Equal(t, expected.ExecTimeMs, actual.ExecTimeMs)
}
