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

// Controller ...
type Controller struct {
	*controllers.ConnectController
	*controllers.UserController
}

// ServiceContainer ... Injecting Dependencies
func ServiceContainer() Controller {
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

	userService := &services.UserService{
		HTTPService: httpService,
	}

	userController := &controllers.UserController{
		UserService: userService,
		Client: clientModel,
	}

	return	Controller{
		connectController,
		userController,
	}
}