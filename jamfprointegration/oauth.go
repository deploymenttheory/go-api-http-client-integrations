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

	"github.com/deploymenttheory/go-api-http-client/logger"
	"go.uber.org/zap"
)

// OAuthResponse represents the response structure when obtaining an OAuth access token from JamfPro.
type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type oauth struct {
	baseDomain   string
	clientId     string
	clientSecret string
	expiryTime   time.Time
	Logger       logger.Logger
}

func (a *oauth) getOauthToken(tokenHome *string, expiryHome *time.Time) error {
	client := http.Client{}
	data := url.Values{}

	data.Set("client_id", a.clientId)
	data.Set("client_secret", a.clientSecret)
	data.Set("grant_type", "client_credentials")

	a.Logger.Debug("Attempting to obtain OAuth token", zap.String("ClientID", a.clientId))

	oauthComlpeteEndpoint := a.baseDomain + oAuthTokenEndpoint
	req, err := http.NewRequest("POST", oauthComlpeteEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
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
	expirationTime := time.Now().Add(expiresIn)

	*tokenHome = oauthResp.AccessToken
	*expiryHome = expirationTime

	return nil
}

func (a *oauth) tokenInBuffer(bufferPeriod time.Duration) bool {
	if time.Until(a.expiryTime) >= bufferPeriod {
		return false
	}
	return true
}

func (a *oauth) tokenExpired() bool {
	if a.expiryTime.Before(time.Now()) {
		return true
	}
	return false
}
