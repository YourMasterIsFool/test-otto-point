-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE transactions(
	id SERIAL PRIMARY KEY,
	total_point float not null,
	user_id int not null,
	foreign key(user_id) references users(id) on delete cascade,
	brand_name varchar(100) not null,
	brand_id int not null,
	foreign key(brand_id) references brands(id) on delete cascade

);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
