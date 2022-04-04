.PHONY: run-dev
run-dev: 
	make run-order-service
	make run-order-worker

.PHONY: run-order-service
run-order-service:
	go run cmd/web/main.go

.PHONY: run-order-worker
run-order-worker:
	go run cmd/worker/main.go