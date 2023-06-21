package request

type CreateOrderRequest struct {
	ProductId int64 `json:"productid"`
	Quantity  int64 `json:"quantity"`
}
