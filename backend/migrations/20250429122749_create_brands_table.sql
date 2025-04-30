-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE brands(
	id SERIAL PRIMARY KEY,
	name varchar(100) not null,
	created_at timestamp default current_timestamp,
	deleted_at timestamp default null
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS brands;
-- +goose StatementEnd
