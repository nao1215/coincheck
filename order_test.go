package coincheck

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClientGetOrderBooksResponse(t *testing.T) {
	t.Run("GetOrderBooksResponse returns a list of order books", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/order_books"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			result := GetOrderBooksResponse{
				Asks: []SellOrderStatus{
					{"27330", "2.25"},
					{"27340", "0.1"},
				},
				Bids: []BuyOrderStatus{
					{"27320", "0.2"},
				},
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		// Create a new client
		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		got, err := client.GetOrderBooks(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		want := &GetOrderBooksResponse{
			Asks: []SellOrderStatus{
				{"27330", "2.25"},
				{"27340", "0.1"},
			},
			Bids: []BuyOrderStatus{
				{"27320", "0.2"},
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("GetOrderBooksResponse returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer testServer.Close()

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		if _, err = client.GetOrderBooks(context.Background()); err == nil {
			t.Error("want error, but got nil")
		}
	})
}
