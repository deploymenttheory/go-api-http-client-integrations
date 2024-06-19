package msgraphintegration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deploymenttheory/go-api-http-client/logger"
	"go.uber.org/zap"
)

type basicAuth struct {
	// Set
	baseDomain   string
	username     string
	password     string
	tenantID     string
	bufferPeriod time.Duration
	logger       logger.Logger

	// Computed
	basicToken            string
	bearerToken           string
	bearerTokenExpiryTime time.Time
}

type basicAuthResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// Operations

// getNewToken obtains a new bearer token from the Microsoft Graph API authentication server.
// This function constructs a new HTTP request to the OAuth2.0 token endpoint using the basic authentication credentials,
// sends the request, and updates the basicAuth instance with the new bearer token and its expiry time.
//
// Returns:
//   - error: Any error encountered during the request, response handling, or JSON decoding. Returns nil if no errors occur.
//
// Functionality:
//   - Constructs the complete OAuth2.0 token endpoint URL using the tenantID.
//   - Logs the constructed authentication URL.
//   - Creates a new HTTP POST request and sets the form data with grant type, scope, username, and password for the request body.
//   - Sends the request using an HTTP client and checks the response status.
//   - Decodes the JSON response to obtain the bearer token and its expiry time.
//   - Updates the basicAuth instance with the new bearer token and its expiry time.
//   - Logs the successful token retrieval with the expiry time and duration.
func (a *basicAuth) getNewToken() error {
	client := http.Client{}

	constructedBearerAuthEndpoint := fmt.Sprintf("%s/%s%s", baseAuthURL, a.tenantID, oAuthTokenEndpoint)

	a.logger.Info("constructed Microsoft Graph API authentication URL", zap.String("URL", constructedBearerAuthEndpoint))

	formData := url.Values{
		"grant_type": {"password"},
		"scope":      {oAuthTokenScope},
		"username":   {a.username},
		"password":   {a.password},
	}

	req, err := http.NewRequest("POST", constructedBearerAuthEndpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

	a.logger.Info("Token obtained successfully", zap.Time("Expiry", a.bearerTokenExpiryTime), zap.Duration("Duration", tokenDuration))

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
	if time.Until(a.bearerTokenExpiryTime) <= a.bufferPeriod {
		return true
	}

	return false
}

// tokenEmpty checks if the current bearer token is empty.
// This function determines if the bearer token string stored in the basicAuth instance is empty.
//
// Returns:
//   - bool: True if the bearer token is empty, false otherwise.
func (a *basicAuth) tokenEmpty() bool {
	if a.bearerToken == "" {
		return true
	}
	return false
}
