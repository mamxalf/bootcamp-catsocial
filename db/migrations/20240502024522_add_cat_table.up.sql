BEGIN;

CREATE TABLE IF NOT EXISTS cats (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(), -- uuid v4
    user_id UUID NOT NULL,
    name varchar NOT NULL,
    race varchar NOT NULL,
    sex boolean NOT NULL DEFAULT FALSE,
    age int NOT NULL DEFAULT 1,
    descriptions TEXT,
    images_url TEXT [],
    has_matched boolean NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_cats_updated_at BEFORE UPDATE ON cats FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- alter table
ALTER TABLE cats ADD FOREIGN KEY (user_id) REFERENCES users (id);

COMMIT;