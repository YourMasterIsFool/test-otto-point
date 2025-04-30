-- +goose Up
-- +goose StatementBegin
ALTER TABLE transaction_details
ADD COLUMN transaction_id int NOT NULL;

ALTER TABLE transaction_details
ADD CONSTRAINT fk_transaction_details_transaction
FOREIGN KEY (transaction_id) REFERENCES transactions(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transaction_details
DROP CONSTRAINT IF EXISTS fk_transaction_details_transaction;

ALTER TABLE transaction_details
DROP COLUMN IF EXISTS transaction_id;
-- +goose StatementEnd
