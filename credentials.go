package coincheck

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

// credentials represents the credentials used to authenticate with the coincheck API.
// If you use Private API, you need to get your API key and API secret from the coincheck website.
type credentials struct {
	// key is the API key.
	key string
	// secret is the API secret.
	secret string
}

// requestHeaderParam represents the parameters to be included in the request header.
//
// For requests that require authentication, you have to add information below to HTTP Header.
// ACCESS-KEY: Access key you genertaed at API key
// ACCESS-NONCE: Positive integer will increase every time you send request (managed for each API key). A common practice is to use UNIX time. Maximum value is 9223372036854775807.
// ACCESS-SIGNATURE: HMAC-SHA256 encoded message containing, ACCESS-NONCE, Request URL and Request body by using API key.
type requestHeaderParam struct {
	// AccessKey is the API key.
	AccessKey string
	// AccessNonce is the positive integer that will increase every time you send a request.
	// A common practice is to use UNIX time. Maximum value is 9223372036854775807.
	AccessNonce string
	// AccessSignature is the HMAC-SHA256 encoded message containing, ACCESS-NONCE, Request URL and Request body by using API key.
	AccessSignature string
}

// generateRequestHeaders generates requestHeaderParam struct.
func (c *credentials) generateRequestHeaders(requestURL *url.URL, body string) (*requestHeaderParam, error) {
	nonce := time.Now().Unix()
	message := fmt.Sprintf("%d%s%s", nonce, requestURL, body)

	h := hmac.New(sha256.New, []byte(c.secret))
	if _, err := h.Write([]byte(message)); err != nil {
		return nil, ErrGenerateRequestHeaders
	}
	signature := hex.EncodeToString(h.Sum(nil))

	return &requestHeaderParam{
		AccessKey:       c.key,
		AccessNonce:     fmt.Sprintf("%d", nonce),
		AccessSignature: signature,
	}, nil
}
