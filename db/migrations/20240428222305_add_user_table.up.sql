BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    name varchar NOT NULL,
    email varchar NOT NULL UNIQUE,
    password varchar NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

CREATE TABLE IF NOT EXISTS user_sessions (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    user_id UUID NOT NULL,

    access_token text NOT NULL, -- encrypted
    refresh_token text NOT NULL, -- encrypted

    is_active boolean NOT NULL DEFAULT false,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_user_sessions_updated_at BEFORE UPDATE ON user_sessions FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- alter table
ALTER TABLE user_sessions ADD FOREIGN KEY (user_id) REFERENCES users (id);

COMMIT;