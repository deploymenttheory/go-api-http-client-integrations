package jamfprointegration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

// platformOAuth implements the authInterface for Jamf platform gateway OAuth2 support.
// The platform gateway uses a different token endpoint (/auth/token) and returns opaque
// (non-JWT) bearer tokens. Tokens expire naturally — there is no invalidation endpoint.
type platformOAuth struct {
	Sugar             *zap.SugaredLogger
	gatewayDomain     string
	clientId          string
	clientSecret      string
	bufferPeriod      time.Duration
	hideSensitiveData bool
	expiryTime        time.Time
	token             string
	http              http.Client
}

// getNewToken obtains a new access token from the platform gateway auth endpoint.
func (a *platformOAuth) getNewToken() error {
	data := url.Values{}
	data.Set("client_id", a.clientId)
	data.Set("client_secret", a.clientSecret)
	data.Set("grant_type", "client_credentials")

	tokenEndpoint := a.gatewayDomain + platformOAuthTokenEndpoint
	a.Sugar.Debugf("platform oauth endpoint constructed: %s", tokenEndpoint)

	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.http.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("bad request getting platform auth token: %v", resp)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Reuse OAuthResponse — the platform gateway returns the same shape (access_token, expires_in)
	oauthResp := &OAuthResponse{}
	err = json.Unmarshal(bodyBytes, oauthResp)
	if err != nil {
		return fmt.Errorf("failed to decode platform OAuth response: %w", err)
	}

	if !a.hideSensitiveData {
		a.Sugar.Debug("platform token received: %+v", oauthResp)
	}

	if oauthResp.AccessToken == "" {
		return fmt.Errorf("empty access token received from platform gateway")
	}

	expiresIn := time.Duration(oauthResp.ExpiresIn) * time.Second
	a.expiryTime = time.Now().Add(expiresIn)
	a.token = oauthResp.AccessToken

	a.Sugar.Infow("Platform token obtained successfully", "expiry", a.expiryTime)
	return nil
}

// getTokenString returns the current token as a string.
func (a *platformOAuth) getTokenString() string {
	return a.token
}

// getExpiryTime returns the current token's expiry time.
func (a *platformOAuth) getExpiryTime() time.Time {
	return a.expiryTime
}

// tokenExpired returns true if the current token has expired.
func (a *platformOAuth) tokenExpired() bool {
	return a.expiryTime.Before(time.Now())
}

// tokenInBuffer returns true if the token's remaining lifetime is within the buffer period.
func (a *platformOAuth) tokenInBuffer() bool {
	return time.Until(a.expiryTime) <= a.bufferPeriod
}

// tokenEmpty returns true if no token has been obtained yet.
func (a *platformOAuth) tokenEmpty() bool {
	return a.token == ""
}
