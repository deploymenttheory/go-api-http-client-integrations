package jamfprointegration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type basicAuth struct {
	// Set
	baseDomain   string
	username     string
	password     string
	bufferPeriod time.Duration
	Sugar        zap.SugaredLogger

	// Computed
	// basicToken            string
	bearerToken           string
	bearerTokenExpiryTime time.Time
}

type basicAuthResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// Operations

// getNewToken obtains a new bearer token from the authentication server.
// This function constructs a new HTTP request to the bearer token endpoint using the basic authentication credentials,
// sends the request, and updates the basicAuth instance with the new bearer token and its expiry time.
//
// Returns:
//   - error: Any error encountered during the request, response handling, or JSON decoding. Returns nil if no errors occur.
//
// Functionality:
//   - Constructs the complete bearer token endpoint URL.
//   - Creates a new HTTP POST request and sets the basic authentication headers.
//   - Sends the request using an HTTP client and checks the response status.
//   - Decodes the JSON response to obtain the bearer token and its expiry time.
//   - Updates the basicAuth instance with the new bearer token and its expiry time.
//   - Logs the successful token retrieval with the expiry time and duration.
//
// TODO migrate strings
func (a *basicAuth) getNewToken() error {
	client := http.Client{}

	completeBearerEndpoint := a.baseDomain + bearerTokenEndpoint
	req, err := http.NewRequest("POST", completeBearerEndpoint, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(a.username, a.password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response status: %d", resp.StatusCode)
	}

	tokenResp := &basicAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(tokenResp)
	if err != nil {
		return err
	}

	a.bearerToken = tokenResp.Token
	a.bearerTokenExpiryTime = tokenResp.Expires
	tokenDuration := time.Until(a.bearerTokenExpiryTime)

	a.Sugar.Info("Token obtained successfully", zap.Time("Expiry", a.bearerTokenExpiryTime), zap.Duration("Duration", tokenDuration))

	return nil
}

// getTokenString returns the current bearer token as a string.
// This function provides access to the current bearer token stored in the basicAuth instance.
//
// Returns:
//   - string: The current bearer token.
func (a *basicAuth) getTokenString() string {
	return a.bearerToken
}

// getExpiryTime returns the expiry time of the current bearer token.
// This function provides access to the expiry time of the current bearer token stored in the basicAuth instance.
//
// Returns:
//   - time.Time: The expiry time of the current bearer token.
func (a *basicAuth) getExpiryTime() time.Time {
	return a.bearerTokenExpiryTime
}

// Utils

// tokenExpired checks if the current bearer token has expired.
// This function compares the current time with the bearer token's expiry time to determine if the token has expired.
//
// Returns:
//   - bool: True if the bearer token has expired, false otherwise.
func (a *basicAuth) tokenExpired() bool {
	return a.bearerTokenExpiryTime.Before(time.Now())
}

// tokenInBuffer checks if the current bearer token is within the buffer period before expiry.
// This function calculates the remaining time until the token expires and compares it with the buffer period.
//
// Returns:
//   - bool: True if the bearer token is within the buffer period, false otherwise.
func (a *basicAuth) tokenInBuffer() bool {
	return time.Until(a.bearerTokenExpiryTime) <= a.bufferPeriod
}

// tokenEmpty checks if the current bearer token is empty.
// This function determines if the bearer token string stored in the basicAuth instance is empty.
//
// Returns:
//   - bool: True if the bearer token is empty, false otherwise.
func (a *basicAuth) tokenEmpty() bool {
	return a.bearerToken == ""
}
