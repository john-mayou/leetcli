package db

import (
	"github.com/john-mayou/leetcli/model"
)

func (c *Client) CreateProblem(problem *model.Problem) (*model.Problem, error) {
	query := `
		INSERT INTO problems (id, slug)
		VALUES (:id, :slug);
	`
	_, err := c.DB.NamedExec(query, problem)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (c *Client) FindProblemByID(id string) (*model.Problem, error) {
	query := `
		SELECT *
		FROM problems
		WHERE id = $1;
	`
	var problem model.Problem
	err := c.DB.Get(&problem, query, id)
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (c *Client) UpdateProblem(problem *model.Problem) error {
	query := `
		UPDATE problems
		SET
			slug = :slug
		WHERE id = :id;
	`
	_, err := c.DB.NamedExec(query, problem)
	return err
}

func (c *Client) DeleteProblem(id string) error {
	query := `
		UPDATE problems
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL;
	`
	_, err := c.DB.Exec(query, id)
	return err
}

func (c *Client) ListProblems() ([]*model.Problem, error) {
	query := `
		SELECT *
		FROM problems
		WHERE deleted_at IS NULL
	`
	var problems []*model.Problem
	err := c.DB.Select(&problems, query)
	if err != nil {
		return nil, err
	}
	return problems, nil
}
