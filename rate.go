package coincheck

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// GetRateInput represents the input for the GetStandardRate function.
type GetRateInput struct {
	// Pair is the pair of the currency. e.g. btc_jpy.
	Pair Pair
}

// GetRateResponse represents the output from GetStandardRate.
type GetRateResponse struct {
	// Rate is the standard rate.
	Rate string `json:"rate"`
}

// GetRate returns the standard rate.
// API: GET /api/rate/[pair]
// Visibility: Public
// If you set don't set the pair, this function will return the standard rate of BTC/JPY.
func (c *Client) GetRate(ctx context.Context, input GetRateInput) (*GetRateResponse, error) {
	pair := PairBTCJPY
	if input.Pair != "" {
		pair = input.Pair
	}

	req, err := c.createRequest(ctx, createRequestInput{
		method: http.MethodGet,
		path:   "/api/rate/" + pair.String(),
	})
	if err != nil {
		return nil, err
	}

	var output GetRateResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

// GetExchangeOrdersRateInput represents the input for the GetExchangeOrdersRate function.
// Either price or amount must be specified as a parameter.
// When both Price and Amount were specified, a status code of 400 was returned in the response.
type GetExchangeOrdersRateInput struct {
	// OrderType is the order type. Order type（"sell" or "buy"）
	OrderType OrderType
	// Pair is the pair of the currency. e.g. btc_jpy.
	// Specify a currency pair to trade. btc_jpy, etc_jpy, lsk_jpy, mona_jpy, plt_jpy, fnct_jpy, dai_jpy, wbtc_jpy, bril_jpy are now available.
	Pair Pair
	// Price is the price of the order. e.g. 30000.
	Price *float64
	// Amount is the amount of the order. e.g. 0.1.
	Amount *float64
}

// GetExchangeOrdersRateResponse represents the output from GetExchangeOrdersRate.
type GetExchangeOrdersRateResponse struct {
	// Success is a boolean value that indicates the success of the API call.
	Success bool `json:"success"`
	// Rate is the rate of the order.
	Rate string `json:"rate"`
	// Price is the price of the order.
	Price string `json:"price"`
	// Amount is the amount of the order.
	Amount string `json:"amount"`
}

// GetExchangeOrdersRate calculate the rate from the order of the exchange.
// API: GET /api/exchange/orders/rate
// Visibility: Public
func (c *Client) GetExchangeOrdersRate(ctx context.Context, input GetExchangeOrdersRateInput) (*GetExchangeOrdersRateResponse, error) {
	queryParam := map[string]string{
		"order_type": input.OrderType.String(),
		"pair":       input.Pair.String(),
	}
	if input.Price == nil && input.Amount == nil {
		return nil, withPrefixError(errors.New("either price or amount must be specified as a parameter"))
	}
	if input.Price != nil {
		queryParam["price"] = fmt.Sprintf("%f", *input.Price)
	} else {
		queryParam["amount"] = fmt.Sprintf("%f", *input.Amount)
	}

	req, err := c.createRequest(ctx, createRequestInput{
		method:     http.MethodGet,
		path:       "/api/exchange/orders/rate",
		queryParam: queryParam,
	})
	if err != nil {
		return nil, err
	}

	var output GetExchangeOrdersRateResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
