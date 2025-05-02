package testutil

import (
	"github.com/google/uuid"
	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/model"
)

func FakeUser() *model.User {
	return &model.User{
		ID:       uuid.NewString(),
		GithubID: "github_" + uuid.NewString(),
		Username: "testuser_" + uuid.NewString(),
		Email:    "test_" + uuid.NewString() + "@example.com",
	}
}

func FakeProblem() *model.Problem {
	return &model.Problem{
		ID:   uuid.NewString(),
		Slug: "testslug_" + uuid.NewString(),
	}
}

func FakeProblemMeta() *sandbox.ProblemMeta {
	return &sandbox.ProblemMeta{
		Number:     1,
		Title:      "testtitle_" + uuid.NewString(),
		Slug:       "testslug_" + uuid.NewString(),
		Difficulty: "easy",
		Prompt:     "testprompt_" + uuid.NewString(),
		Input:      "testinput" + uuid.NewString(),
		Expected:   "testexpected" + uuid.NewString(),
	}
}

func FakeProblemSubmission(problemID, userID string) *model.ProblemSubmission {
	return &model.ProblemSubmission{
		ID:         uuid.NewString(),
		ProblemID:  problemID,
		UserID:     userID,
		Status:     model.ProblemSubmissionStutusPending,
		Code:       "testcode_" + uuid.NewString(),
		Output:     "testoutput_" + uuid.NewString(),
		ExecTimeMs: 1,
	}
}
