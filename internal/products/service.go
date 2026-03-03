package products

import (
	"context"
	"fmt"

	repo "github.com/cristianemek/go-ecommerce/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)

}

func (s *svc) GetProductByID(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *svc) CreateProduct(ctx context.Context, tempProduct CreateProductParams) (repo.Product, error) {
	if len(tempProduct.Name) == 0 {
		return repo.Product{}, fmt.Errorf("product name cannot be empty")
	}

	if tempProduct.PriceInCents <= 0 {
		return repo.Product{}, fmt.Errorf("product price must be greater than 0")
	}

	if tempProduct.Quantity < 0 {
		return repo.Product{}, fmt.Errorf("product quantity cannot be negative")
	}

	product, err := s.repo.CreateProduct(ctx, repo.CreateProductParams{
		Name:         tempProduct.Name,
		PriceInCents: tempProduct.PriceInCents,
		Quantity:     tempProduct.Quantity,
	})

	if err != nil {
		return repo.Product{}, err
	}

	return product, nil
}
