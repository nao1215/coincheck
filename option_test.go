package coincheck

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	t.Run("WithHTTPClient sets the HTTP client which will be used to make requests", func(t *testing.T) {
		t.Parallel()

		client := &http.Client{}
		c, err := NewClient(WithHTTPClient(client))
		if err != nil {
			t.Fatalf("NewClient returned unexpected error: %v", err)
		}

		if diff := cmp.Diff(c.client, client); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("WithHTTPClient returns an error if the HTTP client is nil", func(t *testing.T) {
		t.Parallel()

		if _, err := NewClient(WithHTTPClient(nil)); !errors.Is(err, ErrNilHTTPClient) {
			t.Errorf("error is not ErrNilHTTPClient: %v", err)
		}
	})

	t.Run("WithBaseURL sets the base URL for the client to make requests to", func(t *testing.T) {
		t.Parallel()

		baseURL := "https://example.com"
		c, err := NewClient(WithBaseURL(baseURL))
		if err != nil {
			t.Fatalf("NewClient returned unexpected error: %v", err)
		}

		if diff := cmp.Diff(c.baseURL.String(), baseURL); diff != "" {
			printDiff(t, diff)
		}
	})

	t.Run("WithBaseURL returns an error if the base URL is invalid", func(t *testing.T) {
		t.Parallel()

		if _, err := NewClient(WithBaseURL(":")); !errors.Is(err, ErrInvalidBaseURL) {
			t.Errorf("error is not ErrInvalidBaseURL: %v", err)
		}
	})
}
