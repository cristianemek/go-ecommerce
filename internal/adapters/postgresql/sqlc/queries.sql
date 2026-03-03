-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;

-- name: CreateOrder :exec
INSERT INTO orders (customer_id) VALUES ($1) RETURNING id;