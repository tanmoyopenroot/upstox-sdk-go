package services

import (
	"fmt"
	"upstox-sdk-go/endpoints"
	"upstox-sdk-go/interfaces"
)

// ConnectService ... Service for connection with client
type ConnectService struct {
	HTTPService interfaces.IHTTPService
}

// GetAccessToken ...
func (service *ConnectService) GetAccessToken(params interface{}, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIAccessToken, endpoints.URIhost)
	err := service.HTTPService.POST(url, params, headers, data)
	return err
}
