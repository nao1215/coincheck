package coincheck

import (
	"context"
	"net/http"
)

// OrderType represents the order type.
type OrderType string

const (
	// OrderTypeBuy is the order type of buy.
	OrderTypeBuy OrderType = "buy"
	// OrderTypeSell is the order type of sell.
	OrderTypeSell OrderType = "sell"
)

// SellOrderStatus represents the sell order status.
// It has the only two elements.
// e.g.  [ "27330", "1.25" ]
type SellOrderStatus []string

// BuyOrderStatus represents the buy order status.
// It has the only two elements.
// e.g.  [ "12100", "0.12" ]
type BuyOrderStatus []string

// GetOrderBooksResponse represents the structure of the API response.
type GetOrderBooksResponse struct {
	// Asks is the sell order status.
	Asks []SellOrderStatus `json:"asks"`
	// Bids is the buy order status.
	Bids []BuyOrderStatus `json:"bids"`
}

// GetOrderBooks fetch order book information.
// API: GET /api/order_books
// Visibility: Public
func (c *Client) GetOrderBooks(ctx context.Context) (*GetOrderBooksResponse, error) {
	req, err := c.createRequest(ctx, createRequestInput{
		method: http.MethodGet,
		path:   "/api/order_books",
	})
	if err != nil {
		return nil, err
	}

	var output GetOrderBooksResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
