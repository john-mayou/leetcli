package db

import (
	"database/sql"
	"errors"

	"github.com/john-mayou/leetcli/model"
)

func (c *Client) CreateUser(user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (id, github_id, username, email)
		VALUES (:id, :github_id, :username, :email);
	`
	_, err := c.DB.NamedExec(query, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) FindUserByID(id string) (*model.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE id = $1;
	`
	var user model.User
	err := c.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) FindUserByGithubID(id string) (*model.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE github_id = $1
	`
	var user model.User
	err := c.DB.Get(&user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (c *Client) UpdateUser(user *model.User) error {
	query := `
		UPDATE users
		SET
			github_id = :github_id,
			username = :username,
			email = :email
		WHERE id = :id;
	`
	_, err := c.DB.NamedExec(query, user)
	return err
}

func (c *Client) DeleteUser(id string) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL;
	`
	_, err := c.DB.Exec(query, id)
	return err
}

func (c *Client) ListUsers() ([]*model.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE deleted_at IS NULL
	`
	var users []*model.User
	err := c.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
