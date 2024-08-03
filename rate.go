package coincheck

import (
	"context"
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
