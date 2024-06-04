package jamfprointegration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

// OAuthResponse represents the response structure when obtaining an OAuth access token.
type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (j *Integration) token(bufferPeriod time.Duration) (string, error) {
	var err error
	var token string
	switch j.AuthMethodDescriptor {
	case "oauth2":
		if j.tokenExpired() || j.tokenInBuffer(bufferPeriod) || j.oauthTokenString == "" {
			token, err = j.getOauthToken()
			if j.tokenExpired() || j.tokenInBuffer(bufferPeriod) {
				return "", errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
			}

			if err != nil {
				return "", err
			}

			return token, nil
		}

	case "bearer":
		return "", errors.New("Not implemented")
		// token, err = j.getBasicToken()

	default:
		return "", errors.New("invalid auth method")
	}

	return token, nil
}

func (j *Integration) getOauthToken() (string, error) {

	client := http.Client{}

	data := url.Values{}
	data.Set("client_id", j.ClientId)
	data.Set("client_secret", j.ClientSecret)
	data.Set("grant_type", "client_credentials")

	j.Logger.Debug("Attempting to obtain OAuth token", zap.String("ClientID", j.ClientId))

	req, err := http.NewRequest("POST", oAuthTokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		j.Logger.Error("Failed to create request for OAuth token", zap.Error(err))
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		j.Logger.Error("Failed to execute request for OAuth token", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		j.Logger.Error("Failed to read response body", zap.Error(err))
		return "", err
	}

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	oauthResp := &OAuthResponse{}
	err = json.Unmarshal(bodyBytes, oauthResp)
	if err != nil {
		j.Logger.Error("Failed to decode OAuth response", zap.Error(err))
		return "", fmt.Errorf("failed to decode OAuth response: %w", err)
	}

	if oauthResp.AccessToken == "" {
		j.Logger.Error("Empty access token received")
		return "", fmt.Errorf("empty access token received")
	}

	expiresIn := time.Duration(oauthResp.ExpiresIn) * time.Second
	expirationTime := time.Now().Add(expiresIn)

	j.oauthTokenString = oauthResp.AccessToken
	j.tokenExpiry = expirationTime

	return j.oauthTokenString, nil
}

func (j *Integration) getBasicToken() (string, error) {
	return "", nil
}

func (j *Integration) invalidateToken() (string, error) {
	return "", nil
}

func (j *Integration) keepAliveToken() (string, error) {
	return "", nil
}

func (j *Integration) tokenInBuffer(bufferPeriod time.Duration) bool {
	j.Logger.Warn(j.tokenExpiry.String())
	j.Logger.Warn(bufferPeriod.String())
	log.Println(j.tokenExpiry.String())
	log.Println(bufferPeriod.String())
	if time.Until(j.tokenExpiry) >= bufferPeriod {
		return false
	}

	return true
}

func (j *Integration) tokenExpired() bool {
	if j.tokenExpiry.Before(time.Now()) {
		return true
	}

	return false
}
