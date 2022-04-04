package order

import (
	"context"
	"fmt"
	"order-service/internal/domain"
	wf "order-service/internal/workflow"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

type Handler struct {
	temporalClient client.Client
	workflow       domain.OrderWorkflow
}

func NewHandler(client client.Client, workflow domain.OrderWorkflow) *Handler {
	return &Handler{
		temporalClient: client,
		workflow:       workflow,
	}
}

func (h *Handler) PurchaseItem(ctx context.Context, args domain.OrderItemArgs) error {
	options := client.StartWorkflowOptions{
		ID:        uuid.NewString(),
		TaskQueue: wf.CreateOrderTaskQueue.String(),
	}

	_, err := h.temporalClient.ExecuteWorkflow(context.Background(), options, h.workflow.HandleOrderPurchase, args)
	if err != nil {
		return fmt.Errorf("error starting purchase item workflow: %w", err)
	}
	return nil
}
