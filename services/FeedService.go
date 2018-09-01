package services

import (
	"fmt"
	"net/url"
	"upstox-sdk-go/endpoints"
	"upstox-sdk-go/interfaces"
)

// FeedService ... Services available for client
type FeedService struct {
	HTTPService interfaces.IHTTPService
}

// GetLiveFeed ...
func (service *FeedService) GetLiveFeed(params url.Values, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URILiveFeed, endpoints.URIhost, params.Get("exchange"), params.Get("symbol"), params.Get("requestType"))
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// SubscribeFeed ...
func (service *FeedService) SubscribeFeed(params url.Values, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URILiveFeedSubscribe, endpoints.URIhost, params.Get("requestType"), params.Get("exchange"), params.Get("symbol"))
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// UnsubscribeFeed ...
func (service *FeedService) UnsubscribeFeed(params url.Values, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URILiveFeedSubscribe, endpoints.URIhost, params.Get("requestType"), params.Get("exchange"), params.Get("symbol"))
	err := service.HTTPService.GET(url, headers, data)
	return err
}
