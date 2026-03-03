package products

import (
	"context"

	repo "github.com/cristianemek/go-ecommerce/internal/adapters/postgresql/sqlc"
)

type CreateProductParams struct {
	Name         string `json:"name"`
	PriceInCents int32  `json:"price_in_cents"`
	Quantity     int32  `json:"quantity"`
}

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (repo.Product, error)
	CreateProduct(ctx context.Context, tempProduct CreateProductParams) (repo.Product, error)
}
