package model

import "time"

type ProblemSubmissionStatus string

const (
	ProblemSubmissionStutusPending  ProblemSubmissionStatus = "pending"
	ProblemSubmissionStatusAccepted ProblemSubmissionStatus = "accepted"
	ProblemSubmissionStatusRejected ProblemSubmissionStatus = "rejected"
	ProblemSubmissionStatusError    ProblemSubmissionStatus = "error"
)

type ProblemSubmission struct {
	ID         string                  `db:"id" json:"id"`
	CreatedAt  time.Time               `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time               `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time              `db:"deleted_at" json:"deleted_at"`
	ProblemID  string                  `db:"problem_id" json:"problem_id"`
	UserID     string                  `db:"user_id" json:"user_id"`
	Status     ProblemSubmissionStatus `db:"status" json:"status"`
	Code       string                  `db:"code" json:"code"`
	ExecTimeMs int                     `db:"exec_time_ms" json:"exec_time_ms"`
}
