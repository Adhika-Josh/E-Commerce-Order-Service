package order_controller

import (
	"e-commerce-order-service/app"
	"e-commerce-order-service/service/order_service"
	"e-commerce-order-service/validator"
	"e-commerce-order-service/validator/validate_order"
	"net/http"

	"github.com/gin-gonic/gin"
)

var o order_service.OrderInterface = order_service.OrderImpl{
	DB: app.ConnectDB(),
}

func PlaceOrder(c *gin.Context) {
	req, err := validate_order.ValidatePlaceOrder(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := o.PlaceOrder(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
