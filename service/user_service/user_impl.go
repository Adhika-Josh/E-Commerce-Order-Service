package user_service

import (
	"e-commerce-order-service/model"
	"e-commerce-order-service/model/entity"
	"e-commerce-order-service/utils"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserImpl struct {
	DB *gorm.DB
}

func (u UserImpl) CreateUser(c *gin.Context, req model.CreateUserRequest) (model.CreateUserResponse, model.Errors) {
	var res model.CreateUserResponse
	var user entity.CustomerDetails
	var custErr model.Errors
	if err := u.DB.Where("user_name = ?", req.UserName).First(&user).Error; err == nil {
		return res, model.Errors{
			Error: "Username already exists",
			Type:  "uplicate_entry_error",
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	jsonStr, err := json.Marshal(req)
	if err != nil {
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	newUser := entity.CustomerDetails{
		CustomerPID:     "CUST_" + utils.GenerateRandString(6),
		CustomerDetails: string(jsonStr),
		CreatedAt:       time.Now(),
	}
	if err := u.DB.Create(&newUser).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to create user",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "User created successfully"
	res.UserPID = newUser.CustomerPID

	return res, custErr
}
func (u UserImpl) LoginUser(c *gin.Context, req model.LoginUserRequest) (model.LoginUserResponse, model.Errors) {
	var res model.LoginUserResponse
	var custErr model.Errors
	var user entity.CustomerDetails

	if err := u.DB.Where("user_name = ?", req.UserName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "Invalid username or password",
				Type:  "invalid_credentials_error",
			}
		}
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	if user.UserName != req.UserName || user.Password != req.Password {
		return res, model.Errors{
			Error: "Invalid username or password",
			Type:  "invalid_credentials_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "Login successful"
	res.UserPID = user.CustomerPID
	return res, custErr
}
func (u UserImpl) DeleteUser(c *gin.Context, req model.DeleteUserRequest) (model.DeleteUserResponse, model.Errors) {
	var user entity.CustomerDetails
	var res model.DeleteUserResponse
	var custErr model.Errors
	if err := u.DB.Where("customer_pid = ?", req.UserPID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "no_user_found",
				Type:  "invalid_credentials_error",
			}
		}
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	if err := u.DB.Delete(&user).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to delete user",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "User deleted successfully"
	res.UserPID = user.CustomerPID
	return res, custErr
}
func (u UserImpl) UpdateUser(c *gin.Context, req model.UpdateUserRequest) (model.UpdateUserResponse, model.Errors) {
	var user entity.CustomerDetails
	var res model.UpdateUserResponse
	var custErr model.Errors

	if err := u.DB.Where("cutsomer_pid = ?", req.UserPID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "No record found for the given Item_PID",
				Type:  "record_not_found",
			}
		}
	}
	updateUser := model.UserDetails{
		Name:    req.UserDetails.Name,
		DOB:     req.UserDetails.DOB,
		Age:     req.UserDetails.Age,
		PhoneNo: req.UserDetails.PhoneNo,
		Address: model.UpdateAddress{
			AddressLine1: req.UserDetails.Address.AddressLine1,
			AddressLine2: req.UserDetails.Address.AddressLine2,
			AddressLine3: req.UserDetails.Address.AddressLine3,
			City:         req.UserDetails.Address.City,
			Country:      req.UserDetails.Address.Country,
			State:        req.UserDetails.Address.State,
			District:     req.UserDetails.Address.District,
			Pincode:      req.UserDetails.Address.Pincode,
		},
		Email: req.UserDetails.Email,
	}
	jsonUserData, err := json.Marshal(updateUser)
	if err != nil {
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}
	user.CustomerPID = req.UserPID
	user.CustomerDetails = string(jsonUserData)
	user.UserName = req.UserDetails.UserName
	user.Password = req.UserDetails.Password
	user.UpdatedAt = time.Now()
	if err := u.DB.Save(&user).Error; err != nil {
		return res, model.Errors{
			Error: "Failed to update product",
			Type:  "internal_server_error",
		}
	}
	res.Status = http.StatusOK
	res.Message = "Admin Details updated successfully"
	res.UserPID = user.CustomerPID
	return res, custErr
}
func (u UserImpl) GetUserByID(c *gin.Context, id string) (model.GetUserDetailsResponse, model.Errors){
    var customer entity.CustomerDetails
	var res model.GetUserDetailsResponse

	if err := u.DB.Where("customer_pid = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{
				Error: "No record found for the given User_PID",
				Type:  "record_not_found",
			}
		}
		return res, model.Errors{
			Error: err.Error(),
			Type:  "internal_server_error",
		}
	}

	var customerDetails model.CustomerDetails
	if err := json.Unmarshal([]byte(customer.CustomerDetails), &customerDetails); err != nil {
		return res, model.Errors{
			Error: "Failed to parse customer details",
			Type:  "internal_server_error",
		}
	}

	res = model.GetUserDetailsResponse{
		CustomerPID:     product.CustomerPID,
		CustomerDetails: customerDetails,
		UserName:        product.UserName,
		Password:        product.Password,
	}
	return res, model.Errors{}

}
