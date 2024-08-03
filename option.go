package coincheck

import (
	"fmt"
	"net/http"
	"net/url"
)

// Option is a parameter to be specified when creating a Coincheck
// client to configure the http client details.
type Option func(*Client) error

// WithHTTPClient sets the HTTP client which will be used to make requests.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		if client == nil {
			return ErrNilHTTPClient
		}
		c.client = client
		return nil
	}
}

// WithBaseURL sets the base URL for the client to make requests to.
func WithBaseURL(u string) Option {
	return func(c *Client) error {
		baseURL, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrInvalidBaseURL, err.Error())
		}
		c.baseURL = baseURL
		return nil
	}
}

// WithCredentials sets the credentials to be used to authenticate with the Coincheck API.
func WithCredentials(key, secret string) Option {
	return func(c *Client) error {
		c.credentials = &credentials{
			key:    key,
			secret: secret,
		}
		return nil
	}
}
