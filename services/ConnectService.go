package services

import (
	"upstox-sdk-go/interfaces"
	"upstox-sdk-go/endpoints"
)

// ConnectService ... Service for connection with client
type ConnectService struct {
	HTTPService interfaces.IHTTPService
}

// GetAccessToken ...
func (service *ConnectService) GetAccessToken(params interface{}, headers map[string][]string, token interface{}) (error) {	
	url := endpoints.URIhost + endpoints.URIAccessToken
	err := service.HTTPService.PostJSON(url, params, headers, token)
	return err 
}
