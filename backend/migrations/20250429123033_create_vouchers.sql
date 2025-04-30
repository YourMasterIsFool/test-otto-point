-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE vouchers(
	id SERIAL PRIMARY KEY,
	point float not null,
	brand_id int not null,
	created_at timestamp default current_timestamp,
	deleted_at timestamp default null,
	foreign key(brand_id) REFERENCES brands(id) on delete cascade
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS vouchers;
-- +goose StatementEnd
