-- +goose Up
-- +goose StatementBegin
CREATE TYPE task_status AS ENUM ('to_do', 'in_progress', 'done');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS task_status;
-- +goose StatementEnd
