package coincheck

import (
	"context"
	"net/http"
)

// GetExchangeStatusInput represents the input parameter for the GetExchangeStatus method.
type GetExchangeStatusInput struct {
	// Pair represents the pair of the currency.
	// If pair is not specified, information on all tradable pairs is returned.
	// If pair is specified, information is returned only for that particular currency.
	Pair *Pair
}

// GetExchangeStatusResponse represents the output from the GetExchangeStatus method.
type GetExchangeStatusResponse struct {
	ExchangeStatus []ExchangeStatus `json:"exchange_status"`
}

// ExchangeStatus represents the exchange status.
type ExchangeStatus struct {
	// Pair is the pair of the currency.
	Pair Pair `json:"pair"`
	// Status is the exchange status (available, itayose, stop).
	Status ExchangeStatusAvailability `json:"status"`
	// Timestamp is the time of the status.
	Timestamp float64 `json:"timestamp"`
	// Availability response whether limit orders ( order ), market orders ( market_order ), or cancel orders ( cancel ) can be placed.
	Availability Availability `json:"availability"`
}

// ExchangeStatusAvailability represents the availability of the exchange.
type ExchangeStatusAvailability string

const (
	// ExchangeStatusAvailabilityAvailable means the availability of available.
	ExchangeStatusAvailabilityAvailable ExchangeStatusAvailability = "available"
	// ExchangeStatusAvailabilityItayose is the availability of itayose.
	// Itayose (Itayose method) is a method of determining the transaction price
	// before orders (bids) are executed by recording them on an order book called "ita."
	// Following the price priority principle, market orders are given priority first.
	// Then, orders are matched starting from the lowest sell orders and highest buy orders.
	// When the quantities of buy and sell orders match, a trade is executed at that price.
	ExchangeStatusAvailabilityItayose ExchangeStatusAvailability = "itayose"
	// ExchangeStatusAvailabilityStop is the availability of stop.
	ExchangeStatusAvailabilityStop ExchangeStatusAvailability = "stop"
)

// Availability represents the availability of the exchange.
type Availability struct {
	// Order is the availability of limit orders.
	Order bool `json:"order"`
	// MarketOrder is the availability of market orders.
	MarketOrder bool `json:"market_order"`
	// Cancel is the availability of cancel orders.
	Cancel bool `json:"cancel"`
}

// GetExchangeStatus get retrieving the status of the exchange.
// API: GET /api/exchange_status
// Visibility: Public
// If GetExchangeStatusInput.Pair is not specified, information on all tradable pairs is returned.
func (c *Client) GetExchangeStatus(ctx context.Context, input GetExchangeStatusInput) (*GetExchangeStatusResponse, error) {
	queryParam := map[string]string{}
	if input.Pair != nil {
		queryParam["pair"] = input.Pair.String()
	}

	req, err := c.createRequest(ctx, createRequestInput{
		method:     http.MethodGet,
		path:       "/api/exchange_status",
		queryParam: queryParam,
	})
	if err != nil {
		return nil, err
	}

	var output GetExchangeStatusResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
