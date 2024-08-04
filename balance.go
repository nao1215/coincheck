package coincheck

import (
	"context"
	"net/http"
)

// GetAccountsBalanceResponse represents the response from the GetAccountsBalance method.
type GetAccountsBalanceResponse struct {
	// Success is true if the request was successful.
	Success bool `json:"success"`
	// JPY is the balance of JPY.
	JPY string `json:"jpy"`
	// BTC is the balance of BTC.
	BTC string `json:"btc"`
	// JPYReserved is amount of JPY for unsettled buying order
	JPYReserved string `json:"jpy_reserved"`
	// BTCReserved is amount of BTC for unsettled selling order
	BTCReserved string `json:"btc_reserved"`
	// JPYLendInUse is JPY amount you are applying for lending (We don't allow you to loan JPY.)
	JPYLendInUse string `json:"jpy_lend_in_use"`
	// BTCLendInUse is BTC Amount you are applying for lending (We don't allow you to loan BTC.)
	BTCLendInUse string `json:"btc_lend_in_use"`
	// JPYLent is JPY lending amount (Currently, we don't allow you to loan JPY.)
	JPYLent string `json:"jpy_lent"`
	// BTCLent is BTC lending amount (Currently, we don't allow you to loan BTC.)
	BTCLent string `json:"btc_lent"`
	// JPYDebt is JPY borrowing amount
	JPYDebt string `json:"jpy_debt"`
	// BTCDebt is BTC borrowing amount
	BTCDebt string `json:"btc_debt"`
	// JPYTsumitate is JPY reserving amount
	JPYTsumitate string `json:"jpy_tsumitate"`
	// BTCTsumitate is BTC reserving amount
	BTCTsumitate string `json:"btc_tsumitate"`
}

// GetAccountsBalance returns the balance of the account.
// API: GET /api/accounts/balance
// Visibility: Private
// It doesn't include jpy_reserved use unsettled orders (it's GET /api/exchange/orders/opens) in jpy btc.
func (c *Client) GetAccountsBalance(ctx context.Context) (*GetAccountsBalanceResponse, error) {
	req, err := c.createRequest(ctx, createRequestInput{
		method:  http.MethodGet,
		path:    "/api/accounts/balance",
		private: true,
	})
	if err != nil {
		return nil, err
	}

	var output GetAccountsBalanceResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
