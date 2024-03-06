-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION track_changes_on_password()
	RETURNS trigger
	LANGUAGE 'plpgsql'
	COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
BEGIN
	NEW.password_last_change_at = NOW();
RETURN NEW;
END;
$BODY$;

CREATE OR REPLACE FUNCTION track_changes_on_users()
    RETURNS trigger
	LANGUAGE 'plpgsql'
	COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$BODY$;

CREATE TRIGGER users_password_change
    BEFORE UPDATE OF password ON users
    FOR EACH ROW EXECUTE PROCEDURE track_changes_on_password();

CREATE TRIGGER users_change
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE PROCEDURE track_changes_on_users();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS users_password_change ON users;
DROP TRIGGER IF EXISTS users_change ON users;
DROP FUNCTION track_changes_on_password();
DROP FUNCTION track_changes_on_users();
-- +goose StatementEnd
