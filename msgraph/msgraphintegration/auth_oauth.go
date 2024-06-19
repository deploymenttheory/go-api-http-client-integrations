package msgraphintegration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

type oauth struct {
	// Set
	baseDomain   string
	clientId     string
	clientSecret string
	tenantID     string
	bufferPeriod time.Duration
	Logger       logger.Logger

	// Computed
	expiryTime time.Time
	token      string
}

// OAuthResponse represents the response structure when obtaining an OAuth access token from Microsoft Graph.
type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// Operations

// getNewToken obtains a new bearer token from the Microsoft Graph authentication server.
// This function constructs a new HTTP request to the OAuth2.0 token endpoint using the client credentials,
// sends the request, and updates the oauth instance with the new bearer token and its expiry time.
//
// Returns:
//   - error: Any error encountered during the request, response handling, or JSON decoding. Returns nil if no errors occur.
//
// Functionality:
//   - Constructs the complete OAuth2.0 token endpoint URL using the tenantID.
//   - Creates a new HTTP POST request and sets the form data with client ID, client secret, and grant type.
//   - Sends the request using an HTTP client and checks the response status.
//   - Decodes the JSON response to obtain the bearer token and its expiry time.
//   - Updates the oauth instance with the new bearer token and its expiry time.
func (a *oauth) getNewToken() error {
	client := http.Client{}
	data := url.Values{}

	data.Set("client_id", a.clientId)
	data.Set("client_secret", a.clientSecret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", oAuthTokenScope)

	oauthCompleteEndpoint := fmt.Sprintf("%s/%s%s", baseAuthURL, a.tenantID, oAuthTokenEndpoint)
	req, err := http.NewRequest("POST", oauthCompleteEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("bad request: %v", resp)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	oauthResp := &OAuthResponse{}
	err = json.Unmarshal(bodyBytes, oauthResp)
	if err != nil {
		return fmt.Errorf("failed to decode OAuth response: %w", err)
	}

	if oauthResp.AccessToken == "" {
		return fmt.Errorf("empty access token received")
	}

	expiresIn := time.Duration(oauthResp.ExpiresIn) * time.Second
	a.expiryTime = time.Now().Add(expiresIn)
	a.token = oauthResp.AccessToken

	return nil
}

// getTokenString returns the current bearer token as a string.
// This function provides access to the current bearer token stored in the oauth instance.
//
// Returns:
//   - string: The current bearer token.
func (a *oauth) getTokenString() string {
	return a.token
}

// getExpiryTime returns the expiry time of the current bearer token.
// This function provides access to the expiry time of the current bearer token stored in the oauth instance.
//
// Returns:
//   - time.Time: The expiry time of the current bearer token.
func (a *oauth) getExpiryTime() time.Time {
	return a.expiryTime
}

// Utils

// tokenExpired checks if the current bearer token has expired.
// This function compares the current time with the bearer token's expiry time to determine if the token has expired.
//
// Returns:
//   - bool: True if the bearer token has expired, false otherwise.
func (a *oauth) tokenExpired() bool {
	return a.expiryTime.Before(time.Now())
}

// tokenInBuffer checks if the current bearer token is within the buffer period before expiry.
// This function calculates the remaining time until the token expires and compares it with the buffer period.
//
// Returns:
//   - bool: True if the bearer token is within the buffer period, false otherwise.
func (a *oauth) tokenInBuffer() bool {
	return time.Until(a.expiryTime) <= a.bufferPeriod
}

// tokenEmpty checks if the current bearer token is empty.
// This function determines if the bearer token string stored in the oauth instance is empty.
//
// Returns:
//   - bool: True if the bearer token is empty, false otherwise.
func (a *oauth) tokenEmpty() bool {
	return a.token == ""
}
