package request

type CreateProductRequest struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"proce"`
}
