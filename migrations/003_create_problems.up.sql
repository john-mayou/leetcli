CREATE TABLE problems (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  slug TEXT UNIQUE NOT NULL
);
SELECT manage_updated_at('problems');

CREATE INDEX index_problems_deleted_at ON problems(deleted_at) WHERE deleted_at IS NULL;