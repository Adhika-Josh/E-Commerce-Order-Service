package internal_api

import (
	"bytes"
	"e-commerce-order-service/internal_api/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductsById(c *gin.Context, request dto.GetProductByIDRequest) (dto.GetProductByIdResponse, error) {
	var res dto.GetProductByIdResponse

	url := "http://localhost:8080/product-service/v1/get/" + request.ID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	var productResponse dto.GetProductByIdResponse
	if err := json.Unmarshal(body, &productResponse); err != nil {
		return res, err
	}
	return productResponse, nil

}

func UpdateProduct(c *gin.Context, request dto.UpdateProductRequest) (dto.UpdateProductResponse, error) {
	var res dto.UpdateProductResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return res, err
	}
	reqBody := bytes.NewBuffer(jsonReq)
	url := "http://localhost:8080/product-service/v1/update-product"
	req, err := http.NewRequest(http.MethodPatch, url, reqBody)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	var productResponse dto.UpdateProductResponse
	if err := json.Unmarshal(body, &productResponse); err != nil {
		return res, err
	}
	return productResponse, nil

}
