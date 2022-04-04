package activity

import (
	"context"
	"fmt"
	"order-service/internal/domain"
)

func NewOrderActivity(repo domain.Repository) *OrderActivity {
	return &OrderActivity{
		repository: repo,
	}
}

type OrderActivity struct {
	repository domain.Repository
}

func (o *OrderActivity) CreateOrder(ctx context.Context, args domain.CreateOrderArgs) (string, error) {
	orderID, err := o.repository.CreateOrder(ctx, args.CustomerID, args.ItemID)
	if err != nil {
		return "", fmt.Errorf("repository.CreateOrder: failed creating order: %w", err)
	}

	return orderID, nil
}

func (o *OrderActivity) ReserveItem(ctx context.Context, args domain.ReserveItemArgs) error {
	fmt.Println("reserving item. too lazy to implement repository")
	return nil
}
