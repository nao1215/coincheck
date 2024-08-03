package coincheck

import (
	"context"
	"net/http"
)

// GetTradesInput represents the input parameter for GetTrades
type GetTradesInput struct {
	// Pair is the pair of the currency. e.g. btc_jpy.
	Pair Pair
}

// GetTradesResponse represents the output from GetTrades
type GetTradesResponse struct {
	// Success is a boolean value that indicates the success of the API call.
	Success bool `json:"success"`
	// Pagination is the pagination of the data.
	Pagination Pagination `json:"pagination"`
	// Data is a list of trades.
	Data []Trade `json:"data"`
}

// Trade represents a trade.
type Trade struct {
	// ID is the trade ID.
	ID int `json:"id"`
	// Amount is the amount of the trade.
	Amount float64 `json:"amount"`
	// Rate is the rate of the trade.
	Rate float64 `json:"rate"`
	// Pair is the pair of the currency.
	Pair Pair `json:"pair"`
	// OrderType is the order type.
	OrderType OrderType `json:"order_type"`
	// CreatedAt is the creation time of the trade.
	CreatedAt string `json:"created_at"`
}

// GetTrades returns a list of trades (order transactions).
// API: GET /api/trades
// Visibility: Public
func (c *Client) GetTrades(ctx context.Context, input GetTradesInput) (*GetTradesResponse, error) {
	req, err := c.createRequest(ctx, createRequestInput{
		method: http.MethodGet,
		path:   "/api/trades",
		queryParam: map[string]string{
			"pair": string(input.Pair),
		},
	})
	if err != nil {
		return nil, err
	}

	var output GetTradesResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
