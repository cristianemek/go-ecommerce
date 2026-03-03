package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/cristianemek/go-ecommerce/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNotStock = errors.New("product not in stock")
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {

	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer_id is required")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least one item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	//! importante usar el repo con la transacción para que todas las operaciones se hagan dentro de la misma transacción, y si algo falla, se haga rollback de todo lo que se haya hecho dentro de esa transacción
	qtx := s.repo.WithTx(tx)

	//crear orden
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, fmt.Errorf("failed to create order: %w", err)
	}

	//mirar si existe el producto
	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound //devuelvo un error personalizado por si quiero en el handler devolver x tipo de error segun el codigo de error, o manejarlo de otra forma
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNotStock
		}

		//crear orden_item
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, fmt.Errorf("failed to create order items: %w", err)
		}

		//al crear la orden, actualizo el stock
		_, err = qtx.UpdateProductStock(ctx, repo.UpdateProductStockParams{
			ID:       item.ProductID,
			Quantity: item.Quantity,
		})

		//si el error es pgx.ErrNoRows, en la transacicon se fue el stock, asi evitamos race conditions
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.Order{}, ErrProductNotStock
		}

		if err != nil {
			return repo.Order{}, fmt.Errorf("failed to update product stock: %w", err)
		}
	}

	tx.Commit(ctx)

	return order, nil
}
