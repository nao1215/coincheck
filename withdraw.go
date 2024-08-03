package coincheck

import (
	"context"
	"net/http"
)

// GetBankAccountsResponse represents the response from the GetBankAccounts API.
type GetBankAccountsResponse struct {
	// Success is a boolean value that indicates the success of the API call.
	Success bool `json:"success"`
	// Data is a list of bank accounts.
	Data []BankAccount
}

// BankAccount represents a bank account.
type BankAccount struct {
	// ID is the bank account ID.
	ID int `json:"id"`
	// BankName is the bank name.
	BankName string `json:"bank_name"`
	// BranchName is the branch name.
	BranchName string `json:"branch_name"`
	// BankAccountType is the bank account type.
	BankAccountType string `json:"bank_account_type"`
	// Number is the bank account number.
	Number string `json:"number"`
	// Name is the bank account name.
	Name string `json:"name"`
}

// GetBankAccounts returns a list of bank account you registered (withdrawal).
// API: GET /api/bank_accounts
// Visibility: Private
func (c *Client) GetBankAccounts(ctx context.Context) (*GetBankAccountsResponse, error) {
	req, err := c.createRequest(ctx, createRequestInput{
		method: http.MethodGet,
		path:   "/api/bank_accounts",
	})
	if err != nil {
		return nil, err
	}
	if err := c.setAuthHeaders(req, ""); err != nil {
		return nil, err
	}

	var output GetBankAccountsResponse
	if err := c.do(req, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
