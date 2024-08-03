package coincheck

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
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
