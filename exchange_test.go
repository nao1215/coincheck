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

func TestClient_GetExchangeStatus(t *testing.T) {
	t.Run("GetExchangeStatus returns the all exchange status", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/exchange_status"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			result := GetExchangeStatusResponse{
				ExchangeStatus: []ExchangeStatus{
					{
						Pair:      PairBTCJPY,
						Status:    ExchangeStatusAvailabilityAvailable,
						Timestamp: 1609459200,
						Availability: Availability{
							Order:       true,
							MarketOrder: true,
							Cancel:      true,
						},
					},
					{
						Pair:      PairBrilJPY,
						Status:    ExchangeStatusAvailabilityItayose,
						Timestamp: 1609459200,
						Availability: Availability{
							Order:       false,
							MarketOrder: false,
							Cancel:      false,
						},
					},
				},
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		got, err := client.GetExchangeStatus(context.Background(), GetExchangeStatusInput{})
		if err != nil {
			t.Fatal(err)
		}
		want := &GetExchangeStatusResponse{
			ExchangeStatus: []ExchangeStatus{
				{
					Pair:      PairBTCJPY,
					Status:    ExchangeStatusAvailabilityAvailable,
					Timestamp: 1609459200,
					Availability: Availability{
						Order:       true,
						MarketOrder: true,
						Cancel:      true,
					},
				},
				{
					Pair:      PairBrilJPY,
					Status:    ExchangeStatusAvailabilityItayose,
					Timestamp: 1609459200,
					Availability: Availability{
						Order:       false,
						MarketOrder: false,
						Cancel:      false,
					},
				},
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("If pair is PairETCJPY, GetExchangeStatus returns the exchange status of PairETCJPY", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/exchange_status"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			wantPair := PairETCJPY
			if got := r.URL.Query().Get("pair"); got != wantPair.String() {
				t.Errorf("Pair: got %v, want %v", got, wantPair)
			}

			result := GetExchangeStatusResponse{
				ExchangeStatus: []ExchangeStatus{
					{
						Pair:      PairETCJPY,
						Status:    ExchangeStatusAvailabilityAvailable,
						Timestamp: 1609459200,
						Availability: Availability{
							Order:       true,
							MarketOrder: true,
							Cancel:      true,
						},
					},
				},
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		got, err := client.GetExchangeStatus(context.Background(), GetExchangeStatusInput{
			Pair: pointer.Ptr(PairETCJPY),
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &GetExchangeStatusResponse{
			ExchangeStatus: []ExchangeStatus{
				{
					Pair:      PairETCJPY,
					Status:    ExchangeStatusAvailabilityAvailable,
					Timestamp: 1609459200,
					Availability: Availability{
						Order:       true,
						MarketOrder: true,
						Cancel:      true,
					},
				},
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("GetExchangeStatus returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer testServer.Close()

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		if _, err = client.GetExchangeStatus(context.Background(), GetExchangeStatusInput{}); err == nil {
			t.Error("want error, but got nil")
		}
	})
}
