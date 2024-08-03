package coincheck

import (
	"context"
	"net/http"
)

// Pair represents the pair of the currency.
type Pair string

// String returns the string representation of the pair.
func (p Pair) String() string {
	return string(p)
}

const (
	// PairBTCJPY is the pair of Bitcoin and Japanese Yen.
	PairBTCJPY Pair = "btc_jpy"
	// PairETCJPY is the pair of Ethereum Classic and Japanese Yen.
	PairETCJPY Pair = "etc_jpy"
	// PairLskJPY is the pair of Lisk and Japanese Yen.
	PairLskJPY Pair = "lsk_jpy"
	// PairMonaJPY is the pair of MonaCoin and Japanese Yen.
	PairMonaJPY Pair = "mona_jpy"
	// PairPltJPY is the pair of Palette Token and Japanese Yen.
	PairPltJPY Pair = "plt_jpy"
	// PairFnctJPY is the pair of FiNANCiE and Japanese Yen.
	PairFnctJPY Pair = "fnct_jpy"
	// PairDaiJPY is the pair of DAI and Japanese Yen.
	PairDaiJPY Pair = "dai_jpy"
	// PairWbtcJPY is the pair of Wrapped Bitcoin and Japanese Yen.
	PairWbtcJPY Pair = "wbtc_jpy"
	// PairBrilJPY is the pair of Brilliantcrypto and Japanese Yen.
	PairBrilJPY Pair = "bril_jpy"
)

// GetTickerInput represents the input parameter for GetTicker.
type GetTickerInput struct {
	// Pair is the pair of the currency. e.g. btc_jpy.
	Pair Pair
}

// GetTickerResponse represents the output from GetTicker.
type GetTickerResponse struct {
	// Last is latest quote.
	Last float64 `json:"last"`
	// Bid is current highest buying order.
	Bid float64 `json:"bid"`
	// Ask is current lowest selling order.
	Ask float64 `json:"ask"`
	// High is highest price in last 24 hours.
	High float64 `json:"high"`
	// Low is lowest price in last 24 hours.
	Low float64 `json:"low"`
	// Volume is trading Volume in last 24 hours.
	Volume float64 `json:"volume"`
	// Timestamp is current time. It's Unix Timestamp.
	Timestamp float64 `json:"timestamp"`
}

// GetTicker check latest ticker information.
// API: GET /api/ticker
// Visibility: Public
// If pair is not specified, you can get the information of btc_jpy.
func (c *Client) GetTicker(ctx context.Context, input GetTickerInput) (*GetTickerResponse, error) {
	req, err := c.createRequest(ctx, createRequestInput{
		method: http.MethodGet,
		path:   "/api/ticker",
		queryParam: map[string]string{
			"pair": string(input.Pair),
		},
	})
	if err != nil {
		return nil, err
	}

	var output GetTickerResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
