package main

import (
	"e-commerce-order-service/app"
	"e-commerce-order-service/middleware"

	"e-commerce-order-service/controller/order_controller"
	"e-commerce-order-service/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func main() {
	app.ConnectDB()
	r := gin.Default()
	r.Use(middleware.PanicHandling())
	r.Use(middleware.ValidateJSONMiddleware())
	user := r.Group("/order-service/v1/user")
	{
		user.POST("/create", user_controller.CreateUser)
		user.POST("/login", user_controller.LoginUser)
		user.DELETE("/delete", user_controller.DeleteUser)
		user.PATCH("/update", user_controller.UpdateUser)
	}
	order := r.Group("/order-service/v1")
	{
		order.POST("/create", order_controller.PlaceOrder)
	}
	r.Run("localhost:8081")

}
