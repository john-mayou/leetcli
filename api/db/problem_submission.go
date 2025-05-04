package db

import (
	"github.com/john-mayou/leetcli/model"
)

func (c *Client) CreateProblemSubmission(ps *model.ProblemSubmission) (*model.ProblemSubmission, error) {
	query := `
		INSERT INTO problem_submissions (id, problem_id, user_id, status, code, exec_time_ms)
		VALUES (:id, :problem_id, :user_id, :status, :code, :exec_time_ms);
	`
	_, err := c.DB.NamedExec(query, ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *Client) FindProblemSubmissionByID(id string) (*model.ProblemSubmission, error) {
	query := `
		SELECT *
		FROM problem_submissions
		WHERE id = $1;
	`
	var ps model.ProblemSubmission
	err := c.DB.Get(&ps, query, id)
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

func (c *Client) UpdateProblemSubmission(ps *model.ProblemSubmission) error {
	query := `
		UPDATE problem_submissions
		SET
			problem_id = :problem_id,
			user_id = :user_id,
			status = :status,
			code = :code,
			exec_time_ms = :exec_time_ms
		WHERE id = :id;
	`
	_, err := c.DB.NamedExec(query, ps)
	return err
}

func (c *Client) DeleteProblemSubmission(id string) error {
	query := `
		UPDATE problem_submissions
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL;
	`
	_, err := c.DB.Exec(query, id)
	return err
}

func (c *Client) ListProblemSubmissions() ([]*model.ProblemSubmission, error) {
	query := `
		SELECT *
		FROM problem_submissions
		WHERE deleted_at IS NULL
	`
	var submissions []*model.ProblemSubmission
	err := c.DB.Select(&submissions, query)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}
