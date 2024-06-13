package user_controller

import (
	"e-commerce-order-service/app"
	"e-commerce-order-service/service/user_service"
	"e-commerce-order-service/validator"
	"e-commerce-order-service/validator/validate_user"
	"net/http"

	"github.com/gin-gonic/gin"
)

var u user_service.UserInterface = user_service.UserImpl{
	DB: app.ConnectDB(),
}

func CreateUser(c *gin.Context) {
	req, err := validate_user.ValidateCreateUser(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := u.CreateUser(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
func LoginUser(c *gin.Context) {
	req, err := validate_user.ValidateLoginUser(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := u.LoginUser(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)

}
func DeleteUser(c *gin.Context) {
	req, err := validate_user.ValidateDeleteUser(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := u.DeleteUser(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
func UpdateUser(c *gin.Context) {
	req, err := validate_user.ValidateUpdateUser(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	apiRes, err := u.UpdateUser(c, req)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	validator.ReturnJsonStruct(c, apiRes)
}
