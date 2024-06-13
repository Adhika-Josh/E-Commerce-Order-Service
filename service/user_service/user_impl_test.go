package user_service

import (
	"e-commerce-order-service/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var c *gin.Context

func TestCreateUserNewUser(t *testing.T) {
	user := model.CreateUserRequest{
		Name:    "Test",
		DOB:     "1998-04-23",
		Age:     "26",
		PhoneNo: "9756342187",
		Address: model.Address{
			AddressLine1: "test",
			AddressLine2: "test",
			AddressLine3: "test",
			City:         "test",
			Country:      "test",
			District:     "test",
			State:        "test",
			Pincode:      "670987",
		},
		UserName: "Test12",
		Password: "Test@123",
	}
	res, err := UserImpl{}.CreateUser(c, user)
	assert.NoError(t, err.Message)
	assert.NotEmpty(t, res)
	expectecRes := model.CreateUserResponse{
		Status:  200,
		Message: "User created successfully",
		UserPID: res.UserPID,
	}
	assert.Equal(t, expectecRes, res)
}
