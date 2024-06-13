package order_service

import (
	"e-commerce-order-service/model"

	"github.com/gin-gonic/gin"
)

type OrderInterface interface {
	PlaceOrder(c *gin.Context, req model.OrderRequest) (model.OrderResponse, model.Errors)
	EditOrder(ctx context.Context, orderID string, req model.EditOrderRequest) (model.OrderResponse, model.Errors)
	DeleteOrder(ctx context.Context, orderID string) (model.DeleteOrderResponse, model.Errors)
	ChangeOrderStatus(ctx context.Context, orderID string, req model.ChangeOrderStatusRequest) (model.OrderResponse, model.Errors)
}
