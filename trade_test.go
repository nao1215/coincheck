package coincheck

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetTrades(t *testing.T) {
	t.Run("In the case of a successful GET /api/trades request", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wantMethod := http.MethodGet
			if got := r.Method; got != wantMethod {
				t.Errorf("Method: got %v, want %v", got, wantMethod)
			}

			wantEndpoint := "/api/trades"
			if got := r.URL.Path; got != wantEndpoint {
				t.Errorf("Endpoint: got %v, want %v", got, wantEndpoint)
			}

			wantPair := PairETCJPY
			if got := r.URL.Query().Get("pair"); got != wantPair.String() {
				t.Errorf("Pair: got %v, want %v", got, wantPair)
			}

			result := GetTradesResponse{
				Success: true,
				Pagination: Pagination{
					Limit:           1,
					PaginationOrder: "desc",
					StartingAfter:   0,
					EndingBefore:    0,
				},
				Data: []Trade{
					{
						ID:        1,
						Amount:    1,
						Rate:      1000000,
						Pair:      PairETCJPY,
						OrderType: OrderTypeBuy,
						CreatedAt: "2021-01-01T00:00:00Z",
					},
					{
						ID:        2,
						Amount:    2,
						Rate:      2000000,
						Pair:      PairETCJPY,
						OrderType: OrderTypeSell,
						CreatedAt: "2021-01-02T00:00:00Z",
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

		input := GetTradesInput{
			Pair: PairETCJPY,
		}
		got, err := client.GetTrades(context.Background(), input)
		if err != nil {
			t.Fatal(err)
		}

		want := &GetTradesResponse{
			Success: true,
			Pagination: Pagination{
				Limit:           1,
				PaginationOrder: PaginationOrderDesc,
				StartingAfter:   0,
				EndingBefore:    0,
			},
			Data: []Trade{
				{
					ID:        1,
					Amount:    1,
					Rate:      1000000,
					Pair:      PairETCJPY,
					OrderType: OrderTypeBuy,
					CreatedAt: "2021-01-01T00:00:00Z",
				},
				{
					ID:        2,
					Amount:    2,
					Rate:      2000000,
					Pair:      PairETCJPY,
					OrderType: OrderTypeSell,
					CreatedAt: "2021-01-02T00:00:00Z",
				},
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("In the case of a failed GET /api/trades request", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		client, err := NewClient(WithBaseURL(testServer.URL))
		if err != nil {
			t.Fatal(err)
		}

		input := GetTradesInput{
			Pair: PairETCJPY,
		}
		if _, err = client.GetTrades(context.Background(), input); err == nil {
			t.Fatal("err must not be nil")
		}
	})
}
