package dto

type GetProductByIdResponse struct {
	ItemPID      string `json:"item_pid"`
	ItemName     string `json:"item_name"`
	ItemCategory string `json:"item_category"`
	ItemQuantity string `json:"item_quantity"`
	ItemPrice    string `json:"item_price"`
}
type GetProductByIDRequest struct {
	ID string `json:"id"`
}
type UpdateProductRequest struct {
	ItemPID      string `json:"item_pid"`
	ItemQuantity string `json:"item_quantity"`
}
type UpdateProductResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	ItemPID string `json:"item_pid"`
}
