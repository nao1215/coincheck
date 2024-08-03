package coincheck

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shogo82148/pointer"
)

func TestClient_GetStandardRate(t *testing.T) {
	t.Run("GetStandardRate returns the standard rate", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/rate/etc_jpy"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			result := GetRateResponse{
				Rate: "1000000",
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		got, err := client.GetRate(context.Background(), GetRateInput{Pair: PairETCJPY})
		if err != nil {
			t.Fatal(err)
		}
		want := &GetRateResponse{
			Rate: "1000000",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("GetStandardRate returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		if _, err = client.GetRate(context.Background(), GetRateInput{Pair: PairETCJPY}); err == nil {
			t.Fatal("expected an error, but got nil")
		}
	})

	t.Run("If the pair is empty, GetStandardRate returns the btc_jpy rate", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/rate/btc_jpy"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			result := GetRateResponse{
				Rate: "1000000",
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		got, err := client.GetRate(context.Background(), GetRateInput{Pair: ""})
		if err != nil {
			t.Fatal(err)
		}
		want := &GetRateResponse{
			Rate: "1000000",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})
}

func TestClient_GetExchangeOrdersRate(t *testing.T) {
	t.Run("GetExchangeOrdersRate returns the exchange orders rate (if input valuer set amount)", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/exchange/orders/rate"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			wantPair := PairBTCJPY
			if got := r.URL.Query().Get("pair"); got != wantPair.String() {
				t.Errorf("pair: got %v, want %v", got, wantPair)
			}
			wantOrderType := OrderTypeBuy
			if got := r.URL.Query().Get("order_type"); got != wantOrderType.String() {
				t.Errorf("order_type: got %v, want %v", got, wantOrderType)
			}
			wantAmount := "1.200000"
			if got := r.URL.Query().Get("amount"); got != wantAmount {
				t.Errorf("amount: got %v, want %v", got, wantAmount)
			}

			result := GetExchangeOrdersRateResponse{
				Success: true,
				Rate:    "9118315.44305",
				Price:   "10941978.53166054",
				Amount:  "1.2",
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

		// Start testing
		input := GetExchangeOrdersRateInput{
			OrderType: OrderTypeBuy,
			Pair:      PairBTCJPY,
			Amount:    pointer.Float64(1.2),
		}
		got, err := client.GetExchangeOrdersRate(context.Background(), input)
		if err != nil {
			t.Fatal(err)
		}
		want := &GetExchangeOrdersRateResponse{
			Success: true,
			Rate:    "9118315.44305",
			Price:   "10941978.53166054",
			Amount:  "1.2",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("GetExchangeOrdersRate returns the exchange orders rate (if input valuer set price)", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/exchange/orders/rate"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			wantPair := PairBTCJPY
			if got := r.URL.Query().Get("pair"); got != wantPair.String() {
				t.Errorf("pair: got %v, want %v", got, wantPair)
			}
			wantOrderType := OrderTypeBuy
			if got := r.URL.Query().Get("order_type"); got != wantOrderType.String() {
				t.Errorf("order_type: got %v, want %v", got, wantOrderType)
			}
			wantPrice := "1200000.000000"
			if got := r.URL.Query().Get("price"); got != wantPrice {
				t.Errorf("price: got %v, want %v", got, wantPrice)
			}

			result := GetExchangeOrdersRateResponse{
				Success: true,
				Rate:    "9118315.44305",
				Price:   "1200000",
				Amount:  "1.1",
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

		// Start testing
		input := GetExchangeOrdersRateInput{
			OrderType: OrderTypeBuy,
			Pair:      PairBTCJPY,
			Price:     pointer.Float64(1200000),
		}
		got, err := client.GetExchangeOrdersRate(context.Background(), input)
		if err != nil {
			t.Fatal(err)
		}
		want := &GetExchangeOrdersRateResponse{
			Success: true,
			Rate:    "9118315.44305",
			Price:   "1200000",
			Amount:  "1.1",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("If the amount and pair are empty, GetExchangeOrdersRate returns an error", func(t *testing.T) {
		client, err := NewClient()
		if err != nil {
			t.Fatal(err)
		}

		if _, err = client.GetExchangeOrdersRate(context.Background(), GetExchangeOrdersRateInput{}); err == nil {
			t.Fatal("expected an error, but got nil")
		}
	})

	t.Run("GetExchangeOrdersRate returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		if _, err = client.GetExchangeOrdersRate(context.Background(), GetExchangeOrdersRateInput{
			OrderType: OrderTypeBuy,
			Pair:      PairBTCJPY,
			Amount:    pointer.Float64(1.2),
		}); err == nil {
			t.Fatal("expected an error, but got nil")
		}
	})
}
