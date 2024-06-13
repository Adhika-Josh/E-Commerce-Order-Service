package validate_user

import (
	"e-commerce-order-service/model"
	"e-commerce-order-service/validator"

	"github.com/gin-gonic/gin"
)

func ValidateCreateUser(c *gin.Context) (req model.CreateUserRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	if !validator.IsValidPincode(req.Address.Pincode) {
		custErr.Error = "Invalid pincode format"
		return req, custErr
	}

	if !validator.IsValidDOB(req.DOB) {
		custErr.Error = "Invalid DOB format. Use YYYY-MM-DD"
		return req, custErr
	}

	if !validator.IsValidEmail(req.Email) {
		custErr.Error = "Invalid email format"
		return req, custErr
	}
	return req, custErr
}
func ValidateLoginUser(c *gin.Context) (req model.LoginUserRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, custErr
}
func ValidateDeleteUser(c *gin.Context) (req model.DeleteUserRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, custErr
}
func ValidateUpdateUser(c *gin.Context) (req model.UpdateUserRequest, custErr model.Errors) {
	custErr = validator.ValidateUnknownParams(&req, c)
	if custErr.Error != "" {
		return req, custErr
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, validator.GetRequestUnableToBindZwError()
	}
	return req, custErr
}
