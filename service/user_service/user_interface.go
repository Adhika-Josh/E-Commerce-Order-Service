package user_service

import (
	"e-commerce-order-service/model"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	CreateUser(c *gin.Context, req model.CreateUserRequest) (model.CreateUserResponse, model.Errors)
	LoginUser(c *gin.Context, req model.LoginUserRequest) (model.LoginUserResponse, model.Errors)
	DeleteUser(c *gin.Context, req model.DeleteUserRequest) (model.DeleteUserResponse, model.Errors)
	UpdateUser(c *gin.Context, req model.UpdateUserRequest) (model.UpdateUserResponse, model.Errors)
	GetUserByID(c *gin.Context, id string) (model.GetUserDetailsResponse, model.Errors)
}
