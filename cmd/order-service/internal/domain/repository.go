package domain

import "context"

type Repository interface {
	CreateOrder(ctx context.Context, customerID string, itemID string) (string, error)
	CreateReservation(ctx context.Context, customerID string, itemID string) error
}
