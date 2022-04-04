package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	inMemoryAdapter "order-service/adapter/in-memory"
	"order-service/internal/activity"
	"order-service/internal/app"
	"order-service/internal/workflow"
	"order-service/ports"

	"go.temporal.io/sdk/client"
)

func main() {
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer temporalClient.Close()

	repo := inMemoryAdapter.New()
	activity := activity.NewOrderActivity(repo)
	workflow := workflow.NewOrderWorkflow(activity)

	app := app.New(temporalClient, workflow)
	httpServer := ports.SetupHTTPServer(app)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	var wait time.Duration
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	httpServer.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
