package order_service

import (
	"e-commerce-order-service/model"

	"github.com/gin-gonic/gin"
)

type OrderInterface interface {
	PlaceOrder(c *gin.Context, req model.OrderRequest) (model.OrderResponse, model.Errors)
}
