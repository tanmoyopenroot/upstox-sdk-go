package services

import (
	"fmt"
	"upstox-sdk-go/interfaces"
	"upstox-sdk-go/endpoints"
)

// ConnectService ... Service for connection with client
type ConnectService struct {
	HTTPService interfaces.IHTTPService
}

// GetAccessToken ...
func (service *ConnectService) GetAccessToken(params interface{}, headers map[string][]string, data interface{}) (error) {	
	url  := fmt.Sprintf(endpoints.URIAccessToken, endpoints.URIhost)
	err := service.HTTPService.PostJSON(url, params, headers, data)
	return err 
}
