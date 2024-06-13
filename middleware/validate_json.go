package middleware

import (
	"e-commerce-order-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.Errors{
				Error: "Request content type must be application/json",
				Type:  "bad_request_error",
			})
			return
		}
		c.Next()
	}
}
