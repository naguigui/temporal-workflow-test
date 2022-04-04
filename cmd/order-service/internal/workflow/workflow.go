package workflow

import (
	"fmt"
	"order-service/internal/activity"
	"order-service/internal/domain"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func NewOrderWorkflow(activity *activity.OrderActivity) *OrderWorkflow {
	return &OrderWorkflow{
		activity: activity,
	}
}

type OrderWorkflow struct {
	activity domain.OrderActivy
}

func (o *OrderWorkflow) HandleOrderPurchase(ctx workflow.Context, args domain.OrderItemArgs) error {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    5,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	createOrderArgs := domain.CreateOrderArgs{
		CustomerID: args.CustomerID,
		ItemID:     args.ItemID,
	}

	var orderID string
	err := workflow.ExecuteActivity(ctx, o.activity.CreateOrder, createOrderArgs).Get(ctx, &orderID)
	if err != nil {
		return fmt.Errorf("failed to execute CreateOrder activity: %w", err)
	}

	reserveItemArgs := domain.ReserveItemArgs{
		CustomerID:    args.CustomerID,
		ItemID:        args.ItemID,
		Quantity:      args.Quantity,
		ReservationID: orderID,
	}
	err = workflow.ExecuteActivity(ctx, o.activity.ReserveItem, reserveItemArgs).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to execute ReserveItem activity: %w", err)
	}
	return nil
}
