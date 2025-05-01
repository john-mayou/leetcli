package testutil

import (
	"github.com/google/uuid"
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
