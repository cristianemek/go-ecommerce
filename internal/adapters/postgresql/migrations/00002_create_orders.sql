-- +goose Up
SELECT 'up SQL query';

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INTEGER NOT NULL,
    price_cents INTEGER NOT NULL,
    CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE
);

CREATE INDEX idx_customer_id ON orders(customer_id);
CREATE INDEX idx_order_id ON order_items(order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP INDEX IF EXISTS idx_order_id;
DROP INDEX IF EXISTS idx_customer_id;

DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd