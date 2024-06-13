package order_service

import (
	"e-commerce-order-service/internal_api/dto"
	"e-commerce-order-service/model"
	"fmt"
	"strconv"

	"e-commerce-order-service/internal_api"

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
	}
	if premiumCount == 3 {
		totalAmount *= 0.9
	}

}
