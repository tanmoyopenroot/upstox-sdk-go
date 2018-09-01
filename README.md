# upstox SDK GO

GO client for Upstox Developer APIs. Write your own strategies to buy and sell stocks with NSE, BSE & MCX Market exchanges.

## View Models

Profile  ... Response Model

```sh
type Profile struct {
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Data      struct {
		ClientID         string   `json:"client_id"`
		BankName         string   `json:"bank_name"`
		BankAccount      string   `json:"bank_account"`
		BankName2        string   `json:"bank_name2,omitempty"`
		BankAccount2     string   `json:"bank_account2,omitempty"`
		BankName3        string   `json:"bank_name3,omitempty"`
		BankAccount3     string   `json:"bank_account3,omitempty"`
		BankName4        string   `json:"bank_name4,omitempty"`
		BankAccount4     string   `json:"bank_account4,omitempty"`
		ProductsEnabled  []string `json:"products_enabled"`
		Email            string   `json:"email"`
		Phone            string   `json:"phone"`
		Name             string   `json:"name"`
		IsActive         bool     `json:"is_active"`
		ExchangesEnabled []string `json:"exchanges_enabled"`
	} `json:"data"`
}
```

Balance  ... Response Model

```sh
type Balance struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Data      struct {
		Equity struct {
			NotionalCash    int     `json:"notional_cash"`
			ExposureMargin  int     `json:"exposure_margin"`
			SpanMargin      int     `json:"span_margin"`
			AvailableMargin float64 `json:"available_margin"`
			UsedMargin      int     `json:"used_margin"`
			AdhocMargin     int     `json:"adhoc_margin"`
			PayinAmount     int     `json:"payin_amount"`
		} `json:"equity"`
		Commodity struct {
			UsedMargin      int `json:"used_margin"`
			ExposureMargin  int `json:"exposure_margin"`
			SpanMargin      int `json:"span_margin"`
			AvailableMargin int `json:"available_margin"`
			NotionalCash    int `json:"notional_cash"`
			LedgerBalance   int `json:"ledger_balance"`
			AdhocMargin     int `json:"adhoc_margin"`
			PayinAmount     int `json:"payin_amount"`
		} `json:"commodity"`
	} `json:"data"`
}
```

Position  ... Response Model

```sh
type Position struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange         string  `json:"exchange"`
		Product          string  `json:"product"`
		Symbol           string  `json:"symbol"`
		Token            int     `json:"token"`
		BuyAmount        int     `json:"buy_amount"`
		SellAmount       int     `json:"sell_amount"`
		BuyQuantity      int     `json:"buy_quantity"`
		SellQuantity     int     `json:"sell_quantity"`
		CfBuyAmount      int     `json:"cf_buy_amount"`
		CfSellAmount     float64 `json:"cf_sell_amount"`
		CfBuyQuantity    int     `json:"cf_buy_quantity"`
		CfSellQuantity   int     `json:"cf_sell_quantity"`
		AvgBuyPrice      string  `json:"avg_buy_price"`
		AvgSellPrice     float64 `json:"avg_sell_price"`
		NetQuantity      int     `json:"net_quantity"`
		ClosePrice       float64 `json:"close_price"`
		LastTradedPrice  int     `json:"last_traded_price"`
		RealizedProfit   string  `json:"realized_profit"`
		UnrealizedProfit float64 `json:"unrealized_profit"`
		CfAvgPrice       string  `json:"cf_avg_price"`
	} `json:"data,omitempty"`
}
```

Holdings  ... Response Model

```sh
type Holdings struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Instrument []struct {
			Exchange string `json:"exchange"`
			Symbol   string `json:"symbol"`
			Token    int    `json:"token"`
		} `json:"instrument"`
		Product         string `json:"product"`
		CollateralType  string `json:"collateral_type"`
		CncUsedQuantity int    `json:"cnc_used_quantity"`
		Quantity        int    `json:"quantity"`
		CollateralQty   int    `json:"collateral_qty"`
		Haircut         int    `json:"haircut"`
		AvgPrice        string `json:"avg_price"`
	} `json:"data,omitempty"`
}
```

MasterContract  ... Response Model

```sh
type MasterContract struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    string      `json:"status"`
	Timestamp time.Time   `json:"timestamp"`
}
```

OrderHistoryModel  ... Response Model

```sh
type OrderHistoryModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange          string `json:"exchange"`
		Token             int    `json:"token"`
		Symbol            string `json:"symbol"`
		Product           string `json:"product"`
		OrderType         string `json:"order_type"`
		Duration          string `json:"duration"`
		Price             int    `json:"price"`
		TriggerPrice      int    `json:"trigger_price"`
		Quantity          int    `json:"quantity"`
		DisclosedQuantity int    `json:"disclosed_quantity"`
		TransactionType   string `json:"transaction_type"`
		AveragePrice      int    `json:"average_price"`
		TradedQuantity    int    `json:"traded_quantity"`
		Message           string `json:"message"`
		ExchangeOrderID   string `json:"exchange_order_id"`
		ParentOrderID     string `json:"parent_order_id"`
		OrderID           string `json:"order_id"`
		ExchangeTime      string `json:"exchange_time"`
		TimeInMicro       string `json:"time_in_micro"`
		Status            string `json:"status"`
		IsAmo             bool   `json:"is_amo"`
		ValidDate         string `json:"valid_date"`
		OrderRequestID    string `json:"order_request_id"`
	} `json:"data,omitempty"`
}
```

OrderDetailsModel  ... Response Model

```sh
type OrderDetailsModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange          string  `json:"exchange"`
		Token             int     `json:"token"`
		Symbol            string  `json:"symbol"`
		Product           string  `json:"product"`
		OrderType         string  `json:"order_type"`
		Duration          string  `json:"duration"`
		Price             int     `json:"price"`
		TriggerPrice      int     `json:"trigger_price"`
		Quantity          int     `json:"quantity"`
		DisclosedQuantity int     `json:"disclosed_quantity"`
		TransactionType   string  `json:"transaction_type"`
		AveragePrice      float64 `json:"average_price"`
		TradedQuantity    int     `json:"traded_quantity"`
		Message           string  `json:"message"`
		ExchangeOrderID   string  `json:"exchange_order_id"`
		ParentOrderID     string  `json:"parent_order_id"`
		OrderID           string  `json:"order_id"`
		ExchangeTime      string  `json:"exchange_time"`
		TimeInMicro       string  `json:"time_in_micro"`
		Status            string  `json:"status"`
		IsAmo             bool    `json:"is_amo"`
		ValidDate         string  `json:"valid_date"`
		FillLeg           string  `json:"fill_leg"`
		OrderRequestID    string  `json:"order_request_id"`
		Report            string  `json:"report"`
		Text              string  `json:"text"`
	} `json:"data,omitempty"`
}
```

TradeBookModel  ... Response Model

```sh
type TradeBookModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange        string  `json:"exchange"`
		Token           int     `json:"token"`
		Symbol          string  `json:"symbol"`
		Product         string  `json:"product"`
		OrderType       string  `json:"order_type"`
		TransactionType string  `json:"transaction_type"`
		TradedQuantity  int     `json:"traded_quantity"`
		ExchangeOrderID string  `json:"exchange_order_id"`
		OrderID         string  `json:"order_id"`
		ExchangeTime    string  `json:"exchange_time"`
		TimeInMicro     string  `json:"time_in_micro"`
		TradedPrice     float64 `json:"traded_price"`
		TradeID         string  `json:"trade_id"`
	} `json:"data"`
}
```

TradeHistoryModel  ... Response Model

```sh
type TradeHistoryModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange        string `json:"exchange"`
		Token           int    `json:"token"`
		Symbol          string `json:"symbol"`
		Product         string `json:"product"`
		OrderType       string `json:"order_type"`
		TransactionType string `json:"transaction_type"`
		TradedQuantity  int    `json:"traded_quantity"`
		ExchangeOrderID string `json:"exchange_order_id"`
		OrderID         string `json:"order_id"`
		ExchangeTime    string `json:"exchange_time"`
		TimeInMicro     string `json:"time_in_micro"`
		TradedPrice     int    `json:"traded_price"`
		TradeID         string `json:"trade_id"`
	} `json:"data"`
}
```

PlaceOrderModel ...  Request and Response Model

```sh
type PlaceOrderModel struct {
	TransactionType   string    `json:"transaction_type,omitempty"`
	Exchange          string    `json:"exchange,omitempty"`
	Symbol            string    `json:"symbol,omitempty"`
	Quantity          int       `json:"quantity,omitempty"`
	OrderType         string    `json:"order_type,omitempty"`
	Product           string    `json:"product,omitempty"`
	TriggerPrice      float64   `json:"trigger_price,omitempty"`
	Squareoff         float64   `json:"squareoff,omitempty"`
	IsAmo             bool      `json:"is_amo,omitempty"`
	DisclosedQuantity int       `json:"disclosed_quantity,omitempty"`
	Stoploss          float64   `json:"stoploss,omitempty"`
	Duration          string    `json:"duration,omitempty"`
	TrailingTicks     int       `json:"trailing_ticks,omitempty"`
	Price             float64   `json:"price,omitempty"`
	Code              int       `json:"code,omitempty"`
	Status            string    `json:"status,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	Message           string    `json:"message,omitempty"`
	Data              struct {
		Exchange          string  `json:"exchange,omitempty"`
		Token             int     `json:"token,omitempty"`
		Symbol            string  `json:"symbol,omitempty"`
		Product           string  `json:"product,omitempty"`
		OrderType         string  `json:"order_type,omitempty"`
		Duration          string  `json:"duration,omitempty"`
		Price             float64 `json:"price,omitempty"`
		TriggerPrice      float64 `json:"trigger_price,omitempty"`
		Quantity          int     `json:"quantity,omitempty"`
		DisclosedQuantity int     `json:"disclosed_quantity,omitempty"`
		TransactionType   string  `json:"transaction_type,omitempty"`
		AveragePrice      float64 `json:"average_price,omitempty"`
		TradedQuantity    int     `json:"traded_quantity,omitempty"`
		Message           string  `json:"message,omitempty"`
		ExchangeOrderID   string  `json:"exchange_order_id,omitempty"`
		ParentOrderID     string  `json:"parent_order_id,omitempty"`
		OrderID           string  `json:"order_id,omitempty"`
		ExchangeTime      string  `json:"exchange_time,omitempty"`
		TimeInMicro       string  `json:"time_in_micro,omitempty"`
		Status            string  `json:"status,omitempty"`
		IsAmo             bool    `json:"is_amo,omitempty"`
		ValidDate         string  `json:"valid_date,omitempty"`
		OrderRequestID    string  `json:"order_request_id,omitempty"`
	} `json:"data,omitempty"`
}
```

ModifyOrderModel ...  Request and Response Model

```sh
type ModifyOrderModel struct {
	OrderID           string    `json:"order_id,omitempty"`
	Quantity          int       `json:"quantity,omitempty"`
	OrderType         string    `json:"order_type,omitempty"`
	Price             float64   `json:"price,omitempty"`
	TriggerPrice      float64   `json:"trigger_price,omitempty"`
	DisclosedQuantity int       `json:"disclosed_quantity,omitempty"`
	Code              int       `json:"code,omitempty"`
	Status            string    `json:"status,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	Message           string    `json:"message,omitempty"`
	Data              struct {
		Exchange          string  `json:"exchange,omitempty"`
		Token             int     `json:"token,omitempty"`
		Symbol            string  `json:"symbol,omitempty"`
		Product           string  `json:"product,omitempty"`
		OrderType         string  `json:"order_type,omitempty"`
		Duration          string  `json:"duration,omitempty"`
		Price             float64 `json:"price,omitempty"`
		TriggerPrice      float64 `json:"trigger_price,omitempty"`
		Quantity          int     `json:"quantity,omitempty"`
		DisclosedQuantity int     `json:"disclosed_quantity,omitempty"`
		TransactionType   string  `json:"transaction_type,omitempty"`
		AveragePrice      float64 `json:"average_price,omitempty"`
		TradedQuantity    int     `json:"traded_quantity,omitempty"`
		Message           string  `json:"message,omitempty"`
		ExchangeOrderID   string  `json:"exchange_order_id,omitempty"`
		ParentOrderID     string  `json:"parent_order_id,omitempty"`
		OrderID           string  `json:"order_id,omitempty"`
		ExchangeTime      string  `json:"exchange_time,omitempty"`
		TimeInMicro       string  `json:"time_in_micro,omitempty"`
		Status            string  `json:"status,omitempty"`
		IsAmo             bool    `json:"is_amo,omitempty"`
		ValidDate         string  `json:"valid_date,omitempty"`
		OrderRequestID    string  `json:"order_request_id,omitempty"`
	} `json:"data,omitempty"`
}
```

CancelOrderModel ... Response Model

```sh
type CancelOrderModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      string    `json:"data"`
}
```

LiveFeed  ... Response Model

```sh
type LiveFeed struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      struct {
		Timestamp    int64   `json:"timestamp"`
		Exchange     string  `json:"exchange"`
		Symbol       string  `json:"symbol"`
		Ltp          float64 `json:"ltp"`
		Open         int     `json:"open"`
		High         float64 `json:"high"`
		Low          float64 `json:"low"`
		Close        float64 `json:"close"`
		Vtt          int     `json:"vtt"`
		Atp          float64 `json:"atp"`
		Oi           string  `json:"oi"`
		SpotPrice    float64 `json:"spot_price"`
		TotalBuyQty  int     `json:"total_buy_qty"`
		TotalSellQty int     `json:"total_sell_qty"`
		LowerCircuit float64 `json:"lower_circuit"`
		UpperCircuit float64 `json:"upper_circuit"`
		Bids         []struct {
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
			Orders   int     `json:"orders"`
		} `json:"bids"`
		Asks []struct {
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
			Orders   int     `json:"orders"`
		} `json:"asks"`
		Ltt int64 `json:"ltt"`
	} `json:"data"`
}
```

Subscribe  ... Response Model

```sh
type Subscribe struct {
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
	Data      struct {
		Exchange string   `json:"exchange"`
		Type     string   `json:"type"`
		Success  bool     `json:"success"`
		Symbol   []string `json:"symbol"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
```

Unsubscribe  ... Response Model

```sh
type Unsubscribe struct {
	Message string `json:"message"`
	Data    struct {
		Symbol   string `json:"symbol"`
		Exchange string `json:"exchange"`
		Type     string `json:"type"`
		Success  bool   `json:"success"`
	} `json:"data"`
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
```

Historical  ... Response Model

```sh
type Historical struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Data      []struct {
		Timestamp int64   `json:"timestamp"`
		Open      float64 `json:"open"`
		High      float64 `json:"high"`
		Low       float64 `json:"low"`
		Close     float64 `json:"close"`
		Volume    int     `json:"volume"`
		Cp        float64 `json:"cp"`
	} `json:"data"`
}
```