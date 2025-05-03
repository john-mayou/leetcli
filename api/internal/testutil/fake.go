package testutil

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/john-mayou/leetcli/internal/sandbox"
	"github.com/john-mayou/leetcli/model"
)

var counter int

func ResetFakeCounter() {
	counter = 0
}

func FakeUser() *model.User {
	return &model.User{
		ID:       counterUUID(),
		GithubID: "github_" + counterStr(),
		Username: "testuser_" + counterStr(),
		Email:    "test_" + counterStr() + "@example.com",
	}
}

func FakeProblem() *model.Problem {
	return &model.Problem{
		ID:   counterUUID(),
		Slug: "testslug_" + counterStr(),
	}
}

func FakeProblemMeta() *sandbox.ProblemMeta {
	return &sandbox.ProblemMeta{
		Number:     1,
		Title:      "testtitle_" + counterStr(),
		Slug:       "testslug_" + counterStr(),
		Difficulty: "easy",
		Prompt:     "testprompt_" + counterStr(),
		Input:      "testinput_" + counterStr(),
		Expected:   "testexpected_" + counterStr(),
	}
}

func FakeProblemSubmission(problemID, userID string) *model.ProblemSubmission {
	return &model.ProblemSubmission{
		ID:         counterUUID(),
		ProblemID:  problemID,
		UserID:     userID,
		Status:     model.ProblemSubmissionStutusPending,
		Code:       "testcode_" + counterStr(),
		Output:     "testoutput_" + strconv.Itoa(counter),
		ExecTimeMs: 1,
	}
}

func counterStr() string {
	counter++
	return strconv.Itoa(counter)
}

func counterUUID() string {
	return uuid.NewMD5(uuid.NameSpaceOID, []byte("id-"+counterStr())).String()
}
