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
}

type Client struct {
	DB *sqlx.DB
}

func NewClient(db *sqlx.DB) *Client {
	return &Client{DB: db}
}
