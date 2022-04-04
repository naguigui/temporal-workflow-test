package domain

import (
	"context"

	"go.temporal.io/sdk/workflow"
)

type OrderActivy interface {
	CreateOrder(ctx context.Context, args CreateOrderArgs) (string, error)
	ReserveItem(ctx context.Context, args ReserveItemArgs) error
}

type OrderWorkflow interface {
	HandleOrderPurchase(ctx workflow.Context, args OrderItemArgs) error
}

type OrderItemArgs struct {
	CustomerID string
	ItemID     string
	Quantity   uint32
}

type CreateOrderArgs struct {
	CustomerID string
	ItemID     string
}

type ReserveItemArgs struct {
	CustomerID    string
	ItemID        string
	Quantity      uint32
	ReservationID string
}
