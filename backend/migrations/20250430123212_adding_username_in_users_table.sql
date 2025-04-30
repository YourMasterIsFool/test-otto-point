-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
alter table users
add column username varchar(100) not null;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
alter table users
drop column username;
-- +goose StatementEnd
