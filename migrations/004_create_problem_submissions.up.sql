CREATE TYPE problem_submission_status AS ENUM ('pending', 'accepted', 'rejected', 'error');

CREATE TABLE problem_submissions (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  problem_id UUID REFERENCES problems(id) NOT NULL,
  user_id UUID REFERENCES users(id) NOT NULL,
  status problem_submission_status NOT NULL,
  code TEXT NOT NULL,
  output TEXT NOT NULL,
  exec_time_ms INTEGER NOT NULL
);
SELECT manage_updated_at('problem_submissions');

CREATE INDEX index_problem_submissions_deleted_at ON problem_submissions(deleted_at);
CREATE INDEX index_problem_submissions_problem_id_and_user_id ON problem_submissions(problem_id, user_id);