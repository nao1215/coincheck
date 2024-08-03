package coincheck

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClientGetTicker(t *testing.T) {
	t.Run("In the case of a successful GET /api/ticker request", func(t *testing.T) {
		// Create a new test server
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if diff := cmp.Diff(wantMethod, r.Method); diff != "" {
				printDiff(t, diff)
			}

			wantEndpoint := "/api/ticker"
			if diff := cmp.Diff(wantEndpoint, r.URL.Path); diff != "" {
				printDiff(t, diff)
			}

			wantPair := PairETCJPY
			if got := r.URL.Query().Get("pair"); got != wantPair.String() {
				t.Errorf("pair: got %v, want %v", got, wantPair)
			}

			result := GetTickerResponse{
				Last:      1000000,
				Bid:       999000,
				Ask:       1001000,
				High:      1002000,
				Low:       998000,
				Volume:    100,
				Timestamp: 1609459200,
			}
			if err := json.NewEncoder(w).Encode(result); err != nil {
				t.Fatal(err)
			}
		}))
		defer testServer.Close()

		// Create a new client
		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		// Start testing
		input := GetTickerInput{
			Pair: PairETCJPY,
		}
		got, err := client.GetTicker(context.Background(), input)
		if err != nil {
			t.Fatal(err)
		}

		// Check the result
		want := &GetTickerResponse{
			Last:      1000000,
			Bid:       999000,
			Ask:       1001000,
			High:      1002000,
			Low:       998000,
			Volume:    100,
			Timestamp: 1609459200,
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("In the case of a failed GET /api/ticker request", func(t *testing.T) {
		// Create a new test server
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer testServer.Close()

		// Create a new client
		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		// Start testing
		input := GetTickerInput{
			Pair: PairETCJPY,
		}
		if _, err = client.GetTicker(context.Background(), input); err == nil {
			t.Error("want error, but got nil")
		}
	})
}
