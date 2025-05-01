CREATE OR REPLACE FUNCTION manage_updated_at(_tbl regclass) RETURNS VOID AS $$
BEGIN
  EXECUTE format('CREATE TRIGGER set_updated_at BEFORE UPDATE ON %s
                  FOR EACH ROW EXECUTE PROCEDURE set_updated_at()', _tbl);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
BEGIN
  IF (
    NEW IS DISTINCT FROM OLD AND
    NEW.updated_at IS NOT DISTINCT FROM OLD.updated_at
  ) THEN
    NEW.updated_at := current_timestamp;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  github_id TEXT UNIQUE NOT NULL,
  username TEXT NOT NULL,
  email TEXT NOT NULL
);
SELECT manage_updated_at('users');