package ports

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"order-service/internal/domain"

	"order-service/internal/app"

	"github.com/gorilla/mux"
)

func SetupHTTPServer(application *app.Application) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/orders", getOrdersHandler(application)).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8081",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	return srv
}

func getOrdersHandler(application *app.Application) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateOrderRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		args := domain.OrderItemArgs{
			CustomerID: req.CustomerID,
			ItemID:     req.ItemID,
			Quantity:   req.Quantity,
		}
		if err := application.OrderHandler.PurchaseItem(context.Background(), args); err != nil {
			// Check app errors and log
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
