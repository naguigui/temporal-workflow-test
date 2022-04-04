package ports

type CreateOrderRequest struct {
	CustomerID string `json:"customerID"`
	ItemID     string `json:"itemID"`
	Quantity   uint32 `json:"quantity"`
}
