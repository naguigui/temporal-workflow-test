package app

import (
	"order-service/internal/app/order"
	"order-service/internal/domain"

	"go.temporal.io/sdk/client"
)

type Application struct {
	OrderHandler order.Handler
}

func New(temporalClient client.Client, workflow domain.OrderWorkflow) *Application {
	return &Application{
		OrderHandler: *order.NewHandler(temporalClient, workflow),
	}
}
