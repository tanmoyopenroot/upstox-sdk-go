package main

import (
	"net/http"
	"time"
	"upstox-sdk-go/services"
	"upstox-sdk-go/models"
	"upstox-sdk-go/controllers"
)


// Const ... for http client
const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// ServiceContainer ... Injecting Dependencies
func ServiceContainer() *controllers.ConnectController {
	client := &http.Client{
		Transport: &http.Transport{
				MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	clientModel := &models.ClientModel{}
	httpService := &services.HTTPService{
		HTTPClient: client,
	}
	connectService := &services.ConnectService{
		HTTPService: httpService,
	}
	connectController := &controllers.ConnectController{
		ConnectService: connectService,
		Client: clientModel,
	}

	return connectController
}