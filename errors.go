package coincheck

import "errors"

var (
	// ErrNilHTTPClient means specified http client is nil.
	ErrNilHTTPClient = errors.New("coincheck: specified HTTP client is nil")
	// ErrInvalidBaseURL means specified base URL is invalid.
	ErrInvalidBaseURL = errors.New("coincheck: specified base URL is invalid")
	// ErrGenerateRequestHeaders means failed to generate request headers.
	// If this error occurs, you should check the API key and API secret.
	ErrGenerateRequestHeaders = errors.New("coincheck: failed to generate request headers")
	// ErrNoCredentials means specified credentials is nil.
	ErrNoCredentials = errors.New("coincheck: specified credentials is nil")
)

// withPrefixError returns an error with the package prefix.
func withPrefixError(err error) error {
	const prefix = "coincheck"
	return errors.New(prefix + ": " + err.Error())
}
