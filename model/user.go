package model

type CreateUserRequest struct {
	Name     string  `json:"name" binding:"required"`
	DOB      string  `json:"dob" binding:"required"`
	Age      string  `json:"age" binding:"required"`
	PhoneNo  string  `json:"phone_no" binding:"required"`
	Address  Address `json:"address" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	UserName string  `json:"user_name" binding:"required"`
	Password string  `json:"password" binding:"required"`
}
type Address struct {
	AddressLine1 string `json:"address_line_1" binding:"required"`
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	City         string `json:"city" binding:"required"`
	District     string `json:"district" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	Pincode      string `json:"pincode" binding:"required"`
}

type CreateUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	UserPID string `json:"user_pid"`
}

type LoginUserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	UserPID string `json:"user_pid"`
}

type DeleteUserRequest struct {
	UserPID string `json:"user_pid" binding:"required"`
}
type DeleteUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	UserPID string `json:"user_pid"`
}
type UpdateUserRequest struct {
	UserPID     string      `json:"user_pid" binding:"required"`
	UserDetails UserDetails `json:"user_details"`
}
type UserDetails struct {
	Name     string        `json:"name"`
	DOB      string        `json:"dob"`
	Age      string        `json:"age"`
	PhoneNo  string        `json:"phone_no"`
	Address  UpdateAddress `json:"address"`
	Email    string        `json:"email"`
	UserName string        `json:"user_name"`
	Password string        `json:"password"`
}
type UpdateUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	UserPID string `json:"user_pid"`
}
type UpdateAddress struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	City         string `json:"city"`
	District     string `json:"district"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode"`
}

type CustomerDetails struct {
	Name     string  `json:"name"`
	DOB      string  `json:"dob"`
	Age      string  `json:"age"`
	PhoneNo  string  `json:"phone_no"`
	Address  Address `json:"address"`
	Email    string  `json:"email"`
}

type GetUserDetailsResponse struct {
	CustomerPID     string          `json:"customer_pid"`
	CustomerDetails CustomerDetails `json:"customer_details"`
	UserName        string          `json:"user_name"`
	Password        string          `json:"password"`
}