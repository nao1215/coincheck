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

func TestClient_GetBalance(t *testing.T) {
	t.Run("GetBalance returns the balance", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if diff := cmp.Diff(wantMethod, r.Method); diff != "" {
				printDiff(t, diff)
			}

			wantEndpoint := "/api/accounts/balance"
			if diff := cmp.Diff(wantEndpoint, r.URL.Path); diff != "" {
				printDiff(t, diff)
			}

			result := GetAccountsBalanceResponse{
				Success:      true,
				JPY:          "0.8401",
				BTC:          "7.75052654",
				JPYReserved:  "3000.0",
				BTCReserved:  "3.5002",
				JPYLendInUse: "1.1",
				BTCLendInUse: "0.3",
				JPYLent:      "0",
				BTCLent:      "1.2",
				JPYDebt:      "0",
				BTCDebt:      "0",
				JPYTsumitate: "10000.0",
				BTCTsumitate: "0.43034",
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		client, err := NewClient(
			WithBaseURL(testServer.URL),
			WithCredentials("api_key", "api_secret"),
		)
		if err != nil {
			t.Fatal(err)
		}

		got, err := client.GetAccountsBalance(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		want := &GetAccountsBalanceResponse{
			Success:      true,
			JPY:          "0.8401",
			BTC:          "7.75052654",
			JPYReserved:  "3000.0",
			BTCReserved:  "3.5002",
			JPYLendInUse: "1.1",
			BTCLendInUse: "0.3",
			JPYLent:      "0",
			BTCLent:      "1.2",
			JPYDebt:      "0",
			BTCDebt:      "0",
			JPYTsumitate: "10000.0",
			BTCTsumitate: "0.43034",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("GetBalance returns an error if the server returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		client, err := NewClient(
			WithBaseURL(testServer.URL),
			WithCredentials("api_key", "api_secret"),
		)
		if err != nil {
			t.Fatal(err)
		}

		_, err = client.GetAccountsBalance(context.Background())
		if err == nil {
			t.Fatal("GetBalance did not return an error")
		}
	})
	t.Run("GetBankAccounts returns an error if client does not set credentials", func(t *testing.T) {
		// Create a new client
		client, err := NewClient(WithBaseURL("https://example.com"))
		if err != nil {
			t.Fatal(err)
		}

		// Start testing
		_, err = client.GetAccountsBalance(context.Background())
		if !errors.Is(err, ErrNoCredentials) {
			t.Errorf("error is not ErrNoCredentials: %v", err)
		}
	})
}
