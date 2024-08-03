package coincheck

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClientGetBankAccounts(t *testing.T) {
	t.Run("GetBankAccounts returns a list of bank accounts", func(t *testing.T) {
		// Create a new test server
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if diff := cmp.Diff(wantMethod, r.Method); diff != "" {
				printDiff(t, diff)
			}

			wantEndpoint := "/api/bank_accounts"
			if diff := cmp.Diff(wantEndpoint, r.URL.Path); diff != "" {
				printDiff(t, diff)
			}

			result := GetBankAccountsResponse{
				Success: true,
				Data: []BankAccount{
					{
						ID:              1,
						BankName:        "bank_name_1",
						BranchName:      "branch_name_1",
						BankAccountType: "bank_account_type_1",
						Number:          "number_1",
						Name:            "name_1",
					},
					{
						ID:              2,
						BankName:        "bank_name_2",
						BranchName:      "branch_name_2",
						BankAccountType: "bank_account_type_2",
						Number:          "number_2",
						Name:            "name_2",
					},
				},
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		// Create a new client
		client, err := NewClient(
			WithBaseURL(testServer.URL),
			WithCredentials("api_key", "api_secret"),
		)
		if err != nil {
			t.Fatal(err)
		}

		// Start testing
		got, err := client.GetBankAccounts(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		// Check the result
		want := &GetBankAccountsResponse{
			Success: true,
			Data: []BankAccount{
				{
					ID:              1,
					BankName:        "bank_name_1",
					BranchName:      "branch_name_1",
					BankAccountType: "bank_account_type_1",
					Number:          "number_1",
					Name:            "name_1",
				},
				{
					ID:              2,
					BankName:        "bank_name_2",
					BranchName:      "branch_name_2",
					BankAccountType: "bank_account_type_2",
					Number:          "number_2",
					Name:            "name_2",
				},
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("GetBankAccounts returns an error if the server return internal server error", func(t *testing.T) {
		// Create a new test server
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		// Create a new client
		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		// Start testing
		if _, err = client.GetBankAccounts(context.Background()); err == nil {
			t.Error("want error, but got nil")
		}
	})

	t.Run("GetBankAccounts returns an error if client does not set credentials", func(t *testing.T) {
		// Create a new client
		client, err := NewClient(WithBaseURL("https://example.com"))
		if err != nil {
			t.Fatal(err)
		}

		// Start testing
		_, err = client.GetBankAccounts(context.Background())
		if !errors.Is(err, ErrNoCredentials) {
			t.Errorf("error is not ErrNoCredentials: %v", err)
		}
	})
}
