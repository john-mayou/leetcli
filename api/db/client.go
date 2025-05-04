package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/john-mayou/leetcli/model"
)

type DBClient interface {
	// user
	CreateUser(user *model.User) (*model.User, error)
	FindUserByID(id string) (*model.User, error)
	FindUserByGithubID(id string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
	ListUsers() ([]*model.User, error)

	// problem
	CreateProblem(problem *model.Problem) (*model.Problem, error)
	FindProblemByID(id string) (*model.Problem, error)
	FindProblemBySlug(slug string) (*model.Problem, error)
	UpdateProblem(problem *model.Problem) error
	DeleteProblem(is string) error
	ListProblems() ([]*model.Problem, error)

	// problem submission
	CreateProblemSubmission(ps *model.ProblemSubmission) (*model.ProblemSubmission, error)
	FindProblemSubmissionByID(id string) (*model.ProblemSubmission, error)
	UpdateProblemSubmission(ps *model.ProblemSubmission) error
	DeleteProblemSubmission(id string) error
	ListProblemSubmissions() ([]*model.ProblemSubmission, error)
}

type Client struct {
	DB *sqlx.DB
}

func NewClient(db *sqlx.DB) *Client {
	return &Client{DB: db}
}
