package app

import (
	"e-commerce-order-service/model/entity"
	"time"

	"gorm.io/gorm"
)

func seedOrderDetails(db *gorm.DB) error {
	orders := []entity.OrderDetails{
		{
			OrderPID:        "ORD_123ABC",
			CustomerDetails: `{"cust_pid": "CUST_001", "cust_name": "John Doe", "cust_mob": "9876543210", "cust_address": {"address_line_1": "123 Main St", "address_line_2": "Apt 101", "pincode": "560001", "city": "Bangalore", "state": "Karnataka", "country": "India"}}`,
			ItemDetails:     `{"item_pid": "ITEM_001", "item_name": "Product 1", "item_quantity": 2, "item_price": 50.00, "item_category": "Premium"}`,
			OrderStatus:     "Placed",
		},
		{
			OrderPID:        "ORD_456DEF",
			CustomerDetails: `{"cust_pid": "CUST_002", "cust_name": "Jane Doe", "cust_mob": "9876543211", "cust_address": {"address_line_1": "456 Elm St", "address_line_2": "Apt 202", "pincode": "560002", "city": "Mumbai", "state": "Maharashtra", "country": "India"}}`,
			ItemDetails:     `{"item_pid": "ITEM_002", "item_name": "Product 2", "item_quantity": 1, "item_price": 30.00, "item_category": "Regular"}`,
			OrderStatus:     "Dispatched",
		},
	}

	for _, o := range orders {
		err := db.Create(&o).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func seedCustomerDetails(db *gorm.DB) error {
	users := []entity.CustomerDetails{
		{
			CustomerPID: "CUS_AD2334",
			CustomerDetails: ` {
				"name": "John Doe",
				"dob": "1990-01-01",
				"age": "34",
				"phone_no": "1234567890",
				"address": {
				  "address_line_1": "123 Main St",
				  "address_line_2": "Apt 4B",
				  "address_line_3": "",
				  "city": "Anytown",
				  "district": "Central",
				  "state": "State",
				  "country": "Country",
				  "pincode": "123456"
				},
				"email": "john.doe@example.com"
			  }`,
			UserName:  "John123",
			Password:  "John123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
		{
			CustomerPID: "CUS_TY567R",
			CustomerDetails: `{
				"name": "Jane Smith",
				"dob": "1985-05-15",
				"age": "39",
				"phone_no": "0987654321",
				"address": {
				  "address_line_1": "456 Another St",
				  "address_line_2": "",
				  "address_line_3": "",
				  "city": "Othertown",
				  "district": "North",
				  "state": "State",
				  "country": "Country",
				  "pincode": "654321"
				},
				"email": "jane.smith@example.com"
			}`,
			UserName:  "Jane123",
			Password:  "Jane123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
	}
	for _, o := range users {
		err := db.Create(&o).Error
		if err != nil {
			return err
		}
	}
	return nil

}
