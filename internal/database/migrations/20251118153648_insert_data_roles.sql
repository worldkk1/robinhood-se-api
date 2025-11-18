-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (id, name) VALUES ('ae4c58a6-101a-4b0b-a63e-e187d1920c7e', 'admin');
INSERT INTO roles (id, name) VALUES ('756b0c5c-e9ff-4e91-aa21-49c0dfdc653c', 'user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
DELETE FROM roles;
-- +goose StatementEnd
