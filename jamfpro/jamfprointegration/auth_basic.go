package jamfprointegration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deploymenttheory/go-api-http-client/logger"
	"go.uber.org/zap"
)

type basicAuth struct {
	// Set
	baseDomain   string
	username     string
	password     string
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

// TODO comment
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

	a.logger.Info("Token obtained successfully", zap.Time("Expiry", a.bearerTokenExpiryTime), zap.Duration("Duration", tokenDuration))

	return nil
}

// TODO comment
func (a *basicAuth) getTokenString() string {
	return a.bearerToken
}

// TODO comment
func (a *basicAuth) getExpiryTime() time.Time {
	return a.bearerTokenExpiryTime
}

// Utils

// TODO comment
func (a *basicAuth) tokenExpired() bool {
	return a.bearerTokenExpiryTime.Before(time.Now())
}

// TODO comment
func (a *basicAuth) tokenInBuffer() bool {
	if time.Until(a.bearerTokenExpiryTime) <= a.bufferPeriod {
		return true
	}

	return false
}

// TODO comment
func (a *basicAuth) tokenEmpty() bool {
	if a.bearerToken == "" {
		return true
	}
	return false
}
