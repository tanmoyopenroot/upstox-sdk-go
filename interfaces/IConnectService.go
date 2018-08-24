package interfaces

// IConnectService ... Interface for ConnectService
type IConnectService interface {
	GetAccessToken(params interface{}, headers map[string][]string, token interface{}) (error)
}
