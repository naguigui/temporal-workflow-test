package main

import (
	"log"

	inMemoryAdapter "order-service/adapter/in-memory"
	"order-service/internal/activity"
	wf "order-service/internal/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	repository := inMemoryAdapter.New()
	activity := activity.NewOrderActivity(repository)
	workflow := wf.NewOrderWorkflow(activity)

	w := worker.New(c, wf.CreateOrderTaskQueue.String(), worker.Options{})
	w.RegisterWorkflow(workflow.HandleOrderPurchase)
	w.RegisterActivity(activity.CreateOrder)
	w.RegisterActivity(activity.ReserveItem)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
