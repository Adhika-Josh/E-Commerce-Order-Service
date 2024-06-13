package model

type OrderRequest struct {
	UserPID  string         `json:"user_pid" binding:"required"`
	Products []ProductOrder `json:"products" binding:"required,dive"`
}

type ProductOrder struct {
	ProductID   string `json:"product_id" binding:"required"`
	ProductName string `json:"product_name" binding:"required"`
	Quantity    string `json:"quantity" binding:"required,min=1"`
}

type OrderResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	OrderID string  `json:"order_id,omitempty"`
	Total   float64 `json:"total,omitempty"`
}

type OrderDetails struct {
	ItemPid      string  `json:"item_pid"`
	ItemName     string  `json:"item_name"`
	ItemQuantity int     `json:"item_quantity"`
	ItemPrice    float64 `json:"item_price"`
	ItemCategory string  `json:"item_category"`
}

type EditOrderRequest struct {
	ItemDetails string `json:"item_details"`
}

type ChangeOrderStatusRequest struct {
	OrderStatus string `json:"order_status"`
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
	Message string `json:"message"`
}
type DeleteOrderResponse struct{
	Status  int     `json:"status"`
	Message string  `json:"message"`
}