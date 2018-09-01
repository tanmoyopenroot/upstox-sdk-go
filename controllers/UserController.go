package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"upstox-sdk-go/interfaces"
	"upstox-sdk-go/models"
	"upstox-sdk-go/viewmodels"
)

// UserController ...
type UserController struct {
	UserService interfaces.IUserService
	Client      *models.ClientModel
}

// GetProfile ... Get user profile
func (controller *UserController) GetProfile() {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Profile{}

	err := controller.UserService.GetProfile(headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Profile ---- ", data)
	} else {
		fmt.Println("Profile ---- ", err)
	}
}

// GetBalance ... Get user balance
func (controller *UserController) GetBalance(kind string) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Balance{}

	err := controller.UserService.GetBalance(kind, headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Balance ---- ", data)
	} else {
		fmt.Println("Balance ---- ", err)
	}
}

// GetPositions ... Get user positions
func (controller *UserController) GetPositions() {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Position{}

	err := controller.UserService.GetPositions(headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Position ---- ", data)
	} else {
		fmt.Println("Position ---- ", err)
	}
}

// GetHoldings ... Get user holdings
func (controller *UserController) GetHoldings() {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Holdings{}

	err := controller.UserService.GetHoldings(headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Holding ---- ", data)
	} else {
		fmt.Println("Holding ---- ", err)
	}
}

// GetMasterContracts ... Get user Master Contracts
func (controller *UserController) GetMasterContracts(exchange string, symbol string, token string) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	params := url.Values{}
	if symbol != "" {
		params.Set("symbol", symbol)
	}
	if token != "" {
		params.Set("token", token)
	}

	data := &viewmodels.MasterContract{}

	err := controller.UserService.GetMasterContracts(exchange, params, headers, data)
	if err == nil && data.Code == 200 {
		fmt.Println("Master ---- ", data)
	} else {
		fmt.Println("Master ---- ", err)
	}
}
