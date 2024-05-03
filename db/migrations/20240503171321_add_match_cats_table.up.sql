BEGIN;

CREATE TABLE IF NOT EXISTS matches (
                                    id SERIAL PRIMARY KEY,
                                    issued_user_id UUID NOT NULL,
                                    match_cat_d UUID NOT NULL,
                                    user_cat_id UUID NOT NULL,
                                    message TEXT,
                                    is_approved boolean NOT NULL DEFAULT FALSE,

                                    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- create trigger trigger for automatically set updated_at on row update
CREATE TRIGGER set_matches_updated_at BEFORE UPDATE ON matches FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- alter table
ALTER TABLE matches ADD FOREIGN KEY (issued_user_id) REFERENCES users (id);
ALTER TABLE matches ADD FOREIGN KEY (match_cat_d) REFERENCES cats (id);
ALTER TABLE matches ADD FOREIGN KEY (user_cat_id) REFERENCES cats (id);

COMMIT;