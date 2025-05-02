DROP INDEX IF EXISTS index_users_auth;
CREATE INDEX index_users_auth ON users(github_id);

DROP INDEX IF EXISTS index_users_deleted_at;
CREATE INDEX index_users_deleted_at ON users(deleted_at);

DROP INDEX IF EXISTS index_problems_deleted_at;
CREATE INDEX index_problems_deleted_at ON problems(deleted_at);