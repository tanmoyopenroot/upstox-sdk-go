package endpoints

// API Endpoints
const (
	URIhost string = "https://api.upstox.com"

	URIAuthorise string = "/index/dialog/authorize?apiKey=%s&redirect_uri=%s&response_type=code"
	URIAccessToken string = "/index/oauth/token"
	URIProfile string = "/index/profile"
	URIMasterContract string = "/index/master-contract/%s"
	URILogout string = "/index/logout"

	URIHoldings string = "/live/profile/holdings"
	URIBalance string = "/live/profile/balance"
	URIPositions string = "/live/profile/positions"
	URIPlaceOrder string = "/live/orders"
	URIGetOrders string = "/live/orders"
	URIGetOrdersInfo string = "/live/orders/%s"
	URIAModifyOrders string = "/live/orders/%s/%s"
	URICancelOrder string = "/live/orders/%s"
	URICancelAllOrders string = "/live/orders"
	URITradesInfo string = "/live/orders/%s/trades"
	URITradeBook string = "/live/trade-book"
	URILiveFeed string = "/live/feed/now/%s/%s/%s"
	URILiveFeedSubscribe string = "/live/feed/sub/%s/%s"
	URILiveFeedUnSubscribe string = "/live/feed/unsub/%s/%s"
	URIOHLC string = "/historical/ohlc/%s/%s"
	URISocketParams string = "/live/socket-params"

	URISocketEndpoint string = "wss://ws-api.upstox.com"
)
