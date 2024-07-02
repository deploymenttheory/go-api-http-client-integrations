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

type oauth struct {
	Sugar *zap.SugaredLogger

	// Set
	baseDomain        string
	clientId          string
	clientSecret      string
	bufferPeriod      time.Duration
	hideSensitiveData bool

	// Computed
	expiryTime time.Time
	token      string
}

// OAuthResponse represents the response structure when obtaining an OAuth access token from JamfPro.
type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Operations

// TODO migrate strings
func (a *oauth) getNewToken() error {
	client := http.Client{}
	data := url.Values{}

	data.Set("client_id", a.clientId)
	data.Set("client_secret", a.clientSecret)
	data.Set("grant_type", "client_credentials")

	oauthComlpeteEndpoint := a.baseDomain + oAuthTokenEndpoint
	a.Sugar.Debugf("oauth endpoint constructed: %s", oauthComlpeteEndpoint)

	req, err := http.NewRequest("POST", oauthComlpeteEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	a.Sugar.Debugf("oauth token request constructed: %+v", req)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("bad request getting auth token: %v", resp)
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

	if !a.hideSensitiveData {
		a.Sugar.Debug("token recieved: %+v", oauthResp)
	}

	if oauthResp.AccessToken == "" {
		return fmt.Errorf("empty access token received")
	}

	expiresIn := time.Duration(oauthResp.ExpiresIn) * time.Second
	a.expiryTime = time.Now().Add(expiresIn)
	a.token = oauthResp.AccessToken

	a.Sugar.Info("Token obtained successfully", zap.Time("Expiry", a.expiryTime))
	return nil
}

// TODO func comment
func (a *oauth) getTokenString() string {
	return a.token
}

// TODO func comment
func (a *oauth) getExpiryTime() time.Time {
	return a.expiryTime
}

// Utils

// TODO func comment
func (a *oauth) tokenExpired() bool {
	return a.expiryTime.Before(time.Now())
}

// TODO func comment
func (a *oauth) tokenInBuffer() bool {
	return time.Until(a.expiryTime) <= a.bufferPeriod
}

// TODO func comment
func (a *oauth) tokenEmpty() bool {
	return a.token == ""
}
