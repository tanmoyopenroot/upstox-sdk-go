package interfaces

// IHTTPService ... Interface fot HTTP
type IHTTPService interface {
	PostJSON(url string, params interface{}, headers map[string][]string, token interface{}) (error)
}
