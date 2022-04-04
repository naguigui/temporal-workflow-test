package in_memory

import (
	"context"

	"github.com/google/uuid"
)

func New() *Repository {
	return &Repository{
		orders: []*Order{},
	}
}

type Repository struct {
	orders []*Order
}

type Order struct {
	ID         string
	CustomerID string
	ItemID     string
}

func (r *Repository) CreateOrder(ctx context.Context, customerID string, itemID string) (string, error) {
	order := &Order{
		ID:         uuid.NewString(),
		CustomerID: customerID,
		ItemID:     itemID,
	}
	r.orders = append(r.orders, order)

	return order.ID, nil
}

func (r *Repository) CreateReservation(ctx context.Context, customerID string, itemID string) error {
	return nil
}
