package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"upstox-sdk-go/interfaces"
	"upstox-sdk-go/models"
	"upstox-sdk-go/viewmodels"
)

// OrderController ...
type OrderController struct {
	OrderService interfaces.IOrderService
	Client       *models.ClientModel
}

// GetOrderHistory ... Get user order history
func (controller *OrderController) GetOrderHistory() {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.OrderHistoryModel{}

	err := controller.OrderService.GetOrderHistory(headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Order History ---- ", data)
	} else {
		fmt.Println("Order History ---- ", err)
	}
}

// GetOrderDetails ... Get user order details
func (controller *OrderController) GetOrderDetails(orderID string) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.OrderDetailsModel{}

	err := controller.OrderService.GetOrderDetails(orderID, headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Order Details ---- ", data)
	} else {
		fmt.Println("Order Details ---- ", err)
	}
}

// GetTradeBook ... Get user trade book
func (controller *OrderController) GetTradeBook() {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.TradeBookModel{}

	err := controller.OrderService.GetTradeBook(headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Trade Book ---- ", data)
	} else {
		fmt.Println("Trade Book ---- ", err)
	}
}

// GetTradeHistory ... Get user trade history
func (controller *OrderController) GetTradeHistory(orderID string) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.TradeHistoryModel{}

	err := controller.OrderService.GetTradeHistory(orderID, headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Trade History ---- ", data)
	} else {
		fmt.Println("Trade Histry ---- ", err)
	}
}

// PlaceOrder ... place user order
func (controller *OrderController) PlaceOrder(params *viewmodels.PlaceOrderModel) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.PlaceOrderModel{}

	err := controller.OrderService.PlaceOrder(params, headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Place Order ---- ", data)
	} else {
		fmt.Println("Place Order ---- ", err)
	}
}

// ModifyOrder ... place user order
func (controller *OrderController) ModifyOrder(params *viewmodels.ModifyOrderModel) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.ModifyOrderModel{}

	err := controller.OrderService.ModifyOrder(data.OrderID, params, headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Modify Order ---- ", data)
	} else {
		fmt.Println("Modify Order ---- ", err)
	}
}

// CancelOrder ... cancel user order
func (controller *OrderController) CancelOrder(orders []string) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.CancelOrderModel{}

	err := controller.OrderService.CancelOrder(strings.Join(orders, ","), headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Cancel Order ---- ", data)
	} else {
		fmt.Println("Cancel Order ---- ", err)
	}
}
