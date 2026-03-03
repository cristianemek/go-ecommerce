package orders

import (
	"context"

	repo "github.com/cristianemek/go-ecommerce/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) ([]repo.Order, error) {
	return

}
