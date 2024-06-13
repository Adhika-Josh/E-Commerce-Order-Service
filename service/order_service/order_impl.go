package order_service

import (
	"e-commerce-order-service/internal_api/dto"
	"e-commerce-order-service/model"
	"fmt"
	"strconv"

	"e-commerce-order-service/internal_api"
	"e-commerce-order-service/service/order_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderImpl struct {
	DB *gorm.DB
}

func (o OrderImpl) PlaceOrder(c *gin.Context, req model.OrderRequest) (model.OrderResponse, model.Errors) {
	var res model.OrderResponse
	var totalAmount float64
	var premiumCount int
	var orderDetails []model.OrderDetails

	for _, v := range req.Products {
		product, err := internal_api.GetProductsById(c, dto.GetProductByIDRequest{
			ID: v.ProductID,
		})
		if err != nil {
			return res, model.Errors{
				Error: "failed to get response from the third party api",
				Type:  "external_service_error",
			}
		}
		if product.ItemQuantity < v.Quantity {
			return res, model.Errors{
				Error: fmt.Sprintf("insffient quantity for the item %v", v.ProductName),
				Type:  "quantity_erro",
			}
		}
		availableQty, err := strconv.Atoi(product.ItemQuantity)
		if err != nil {
			return res, model.Errors{
				Error: fmt.Sprintf("failed to convert %v to integer", product.ItemQuantity),
				Type:  "internal_package_error",
			}
		}
		requestedQuantity, err := strconv.Atoi(v.Quantity)
		if err != nil {
			return res, model.Errors{
				Error: fmt.Sprintf("failed to convert %v to integer", v.Quantity),
				Type:  "internal_package_error",
			}
		}
		remainingQuantity := strconv.Itoa(availableQty - requestedQuantity)
		_, err = internal_api.UpdateProduct(c, dto.UpdateProductRequest{
			ItemPID:      product.ItemPID,
			ItemQuantity: remainingQuantity,
		})
		if err != nil {
			return res, model.Errors{
				Error: "failed to update product quantity",
				Type:  "internal_server_error",
			}
		}
		itemPrice, err := strconv.ParseFloat(product.ItemQuantity, 64)
		if err != nil {
			return res, model.Errors{
				Error: fmt.Sprintf("failed to convert %v to float64", product.ItemPrice),
				Type:  "internal_package_error",
			}
		}
		totalAmount += itemPrice
		if product.ItemCategory == "Premium" {
			premiumCount++
		}
		orderDetails = append(orderDetails, model.OrderDetails{
			ItemPid:      product.ItemPID,
			ItemName:     product.ItemName,
			ItemQuantity: requestedQuantity,
			ItemPrice:    itemPrice,
			ItemCategory: product.ItemCategory,
		})
	}
	if premiumCount == 3 {
		totalAmount *= 0.9
	}
	customerDetails, err := order_service.GetUserByID(c,req.UserPID)
	if err != nil {
		return res, model.Errors{
			Error: "failed to get customer details",
			Type:  "internal_server_error",
		}
	}
	

customerDetails, err := GetCustomerDetails(req.CustomerPID)
	if err != nil {
		return res, model.Errors{
			Error: "failed to get customer details",
			Type:  "internal_server_error",
		}
	}

	customerDetailsJSON, err := json.Marshal(customerDetails)
	if err != nil {
		return res, model.Errors{
			Error: "failed to marshal customer details",
			Type:  "internal_server_error",
		}
	}

	orderDetailsJSON, err := json.Marshal(orderDetails)
	if err != nil {
		return res, model.Errors{
			Error: "failed to marshal order details",
			Type:  "internal_server_error",
		}
	}
	newOrder := entity.OrderDetails{
		OrderPID:        "ORD_" + utils.GenerateRandString(6),
		CustomerDetails: string(customerDetailsJSON),
		ItemDetails:     string(orderDetailsJSON),
		OrderStatus:     "Placed",
		TotalAmount:     totalAmount,
	}

	if err := u.DB.Create(&newOrder).Error; err != nil {
		return res, model.Errors{
			Error: "failed to create order",
			Type:  "internal_server_error",
		}
	}

	res.Status = http.StatusOK
	res.Message = "Order placed successfully"
	res.OrderID = newOrder.OrderPID
	res.TotalAmount = totalAmount
	return res, model.Errors{}
}

func GetCustomerDetails(customerPID string) (model.CustomerDetails, error) {
	var customer entity.CustomerDetails
	if err := u.DB.Where("customer_pid = ?", customerPID).First(&customer).Error; err != nil {
		return model.CustomerDetails{}, err
	}

	var customerDetails model.CustomerDetails
	if err := json.Unmarshal([]byte(customer.CustomerDetails), &customerDetails); err != nil {
		return model.CustomerDetails{}, err
	}
	return customerDetails, nil
}
func (o OrderImpl) EditOrder(ctx context.Context, orderID string, req model.EditOrderRequest) (model.OrderResponse, model.Errors) {
	var res model.OrderResponse
	var order entity.OrderDetails
	err := o.DB.Where("order_pid",&order)
	if err != nil {
		return res, model.Errors{Error: "Order not found", Type: "record_not_found"}
	}

	order = req.UpdateOrderDetails(order)
	err = o.OrderRepo.UpdateOrder(ctx, order)
	if err != nil {
		return res, model.Errors{Error: "Failed to update order", Type: "internal_server_error"}
	}

	res = order.ToOrderResponse()
	return res, model.Errors{}
}
func (req model.EditOrderRequest) UpdateOrderDetails(order entity.OrderDetails) entity.OrderDetails {
	customerDetailsBytes, _ := json.Marshal(req.CustomerDetails)
	itemDetailsBytes, _ := json.Marshal(req.ItemDetails)

	order.CustomerDetails = string(customerDetailsBytes)
	order.ItemDetails = string(itemDetailsBytes)
	order.OrderStatus = req.OrderStatus
	order.UpdatedAt = time.Now()

	return order
}
func (order entity.OrderDetails) ToOrderResponse() model.OrderResponse {
	var customerDetails model.CustomerDetails
	var itemDetails model.ItemDetails

	_ = json.Unmarshal([]byte(order.CustomerDetails), &customerDetails)
	_ = json.Unmarshal([]byte(order.ItemDetails), &itemDetails)

	return model.OrderResponse{
		OrderPID:        order.OrderPID,
		CustomerDetails: customerDetails,
		ItemDetails:     itemDetails,
		OrderStatus:     order.OrderStatus,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
}
func (o OrderImpl) DeleteOrder(ctx context.Context, orderID string) (model.DeleteOrderResponse, model.Errors) {
	err:= r.DB.WithContext(ctx).Where("order_id = ?", orderID).Delete(&entity.Order{}).Error
	if err != nil {
		return res, model.Errors{
			Error: "failed to get customer details",
			Type:  "internal_server_error",
		}
	}
	res:= model.DeleteOrderResponse{
		Status:http.StatusOK,
		Message:"Order deleted successfully"
	}
	return res,model.Errors{}
}

func (o OrderImpl) ChangeOrderStatus(ctx context.Context, orderID string, req model.ChangeOrderStatusRequest) (model.OrderResponse, model.Errors) {
	var res model.OrderResponse
	var order entity.OrderDetails
	if err := o.DB.Where("order_pid = ?", orderID).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, model.Errors{Error: "Order not found", Type: "record_not_found"}
		}
		return res, model.Errors{Error: "Failed to retrieve order", Type: "internal_server_error"}
	}

	order.OrderStatus = req.OrderStatus
	if err := o.DB.Save(&order).Error; err != nil {
		return res, model.Errors{Error: "Failed to update order status", Type: "internal_server_error"}
	}
	res = order.ToOrderResponse()
	return res, model.Errors{}
}

func (order *entity.OrderDetails) ToOrderResponse() model.OrderResponse {
	var customerDetails model.CustomerDetails
	var itemDetails model.ItemDetails

	_ = json.Unmarshal([]byte(order.CustomerDetails), &customerDetails)
	_ = json.Unmarshal([]byte(order.ItemDetails), &itemDetails)

	return model.OrderResponse{
		OrderPID:        order.OrderPID,
		CustomerDetails: customerDetails,
		ItemDetails:     itemDetails,
		OrderStatus:     order.OrderStatus,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
}