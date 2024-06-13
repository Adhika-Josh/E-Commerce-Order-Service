package validate_order

import (
	"e-commerce-order-service/model"
	"e-commerce-order-service/validator"

	"github.com/gin-gonic/gin"
)

func ValidatePlaceOrder(c *gin.Context) (req model.OrderRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	if len(req.Products) > 10 {
		return req, model.Errors{Error: "Cannot order more than 10 products", Type: "validation_error"}
	}
	return req, custErr
}
func ValidateEditOrder(c *gin.Context) (req model.EditOrderRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, model.Errors{Error: "Invalid request", Type: "validation_error"}
	}
	return req, model.Errors{}
}

func ValidateChangeOrderStatus(c *gin.Context) (req model.ChangeOrderStatusRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, model.Errors{Error: "Invalid request", Type: "validation_error"}
	}
	return req, model.Errors{}
}