-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE transaction_details(
	id SERIAL PRIMARY KEY,
	point float not null,
	voucher_id int not null,
	foreign key(voucher_id) references vouchers(id) on delete cascade,
	brand_name varchar(100) not null,
	brand_id int not null,
	foreign key(brand_id) references brands(id) on delete cascade
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS transaction_details;
-- +goose StatementEnd
