// Package coincheck provides a client for using the coincheck API.
// Ref. https://coincheck.com/documents/exchange/api
//
// The coincheck allows roughly two kinds of APIs; Public API and Private API.
// Public API allows you to browse order status and order book.
// Private API allows you to create and cancel new orders, and to check your balance.
// If you use Private API, you need to get your API key and API secret from the coincheck website.
package coincheck

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// BaseURL is the base URL for the coincheck API.
	BaseURL = "https://coincheck.com"
)

// Client represents a coincheck client.
type Client struct {
	// client is HTTP client that used to communicate with the Coincheck API.
	client *http.Client
	// baseURL is the base URL for the coincheck API.
	baseURL *url.URL
	// credentials is the credentials used to authenticate with the coincheck API.
	credentials *credentials
}

// NewClient returns a new coincheck client.
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		client: http.DefaultClient,
	}

	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		return nil, withPrefixError(err)
	}
	c.baseURL = baseURL

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// hasCredentials returns true if the client has credentials.
func (c *Client) hasCredentials() bool {
	return c.credentials != nil
}

// setAuthHeaders sets the authentication headers to the request.
// If you use Private API, you need to get your API key and API secret from the coincheck website.
func (c *Client) setAuthHeaders(req *http.Request, body string) error {
	if !c.hasCredentials() {
		return ErrNoCredentials
	}

	headers, err := c.credentials.generateRequestHeaders(req.URL, body)
	if err != nil {
		return withPrefixError(err)
	}
	req.Header.Set("ACCESS-KEY", headers.AccessKey)
	req.Header.Set("ACCESS-NONCE", headers.AccessNonce)
	req.Header.Set("ACCESS-SIGNATURE", headers.AccessSignature)
	return nil
}

// createRequestInput represents the input parameters for createRequest.
type createRequestInput struct {
	method     string            // HTTP method (e.g. GET, POST)
	path       string            // API path (e.g. /api/orders)
	body       io.Reader         // Request body. If you don't need it, set nil.
	queryParam map[string]string // Query parameters (e.g. {"pair": "btc_jpy"}) If you don't need it, set nil.
	private    bool              // If true, it's a private API.
}

// createRequest creates a new HTTP request.
func (c *Client) createRequest(ctx context.Context, input createRequestInput) (*http.Request, error) {
	u, err := url.JoinPath(c.baseURL.String(), input.path)
	if err != nil {
		return nil, withPrefixError(err)
	}

	endpoint, err := url.Parse(u)
	if err != nil {
		return nil, withPrefixError(err)
	}

	if input.queryParam != nil {
		q := endpoint.Query()
		for k, v := range input.queryParam {
			q.Set(k, v)
		}
		endpoint.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, input.method, endpoint.String(), input.body)
	if err != nil {
		return nil, withPrefixError(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	if input.private {
		if err := c.setAuthHeaders(req, ""); err != nil {
			return nil, err
		}
	}

	return req, nil
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) do(req *http.Request, output any) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return withPrefixError(err)
	}
	defer resp.Body.Close() //nolint: errcheck // ignore error

	if resp.StatusCode != http.StatusOK {
		return withPrefixError(fmt.Errorf("unexpected status code=%d", resp.StatusCode))
	}

	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return withPrefixError(err)
	}
	return nil
}
