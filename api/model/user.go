package model

import "time"

type User struct {
	ID        string     `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	GithubID  string     `db:"github_id" json:"github_id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
}
