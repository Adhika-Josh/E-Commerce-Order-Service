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
