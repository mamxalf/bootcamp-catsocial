BEGIN;

DROP TABLE IF EXISTS cats CASCADE;
DROP TRIGGER IF EXISTS set_cats_updated_at ON cats CASCADE;

COMMIT;