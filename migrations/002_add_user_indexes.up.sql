CREATE INDEX index_users_auth ON users(github_id) WHERE deleted_at IS NULL;
CREATE INDEX index_users_deleted_at ON users(deleted_at) WHERE deleted_at IS NULL;