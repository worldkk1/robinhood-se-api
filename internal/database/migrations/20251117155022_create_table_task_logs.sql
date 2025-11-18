-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_logs (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  task_id UUID NOT NULL,
  title TEXT,
  description TEXT,
  status task_status,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task_logs;
-- +goose StatementEnd
