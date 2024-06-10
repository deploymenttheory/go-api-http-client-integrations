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
	username              string
	password              string
	basicToken            string
	bearerToken           string
	bearerTokenExpiryTime time.Time
	logger                logger.Logger
}

type basicAuthResponse struct {
	token  string
	expiry time.Time
}

// Operations

func (a *basicAuth) getNewToken() error {
	client := http.Client{}

	req, err := http.NewRequest("POST", bearerTokenEndpoint, nil)
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

	tokenResp := &TokenResponse{}
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

func (a *basicAuth) tokenString() string {
	return a.bearerToken
}

// Utils

func (a *basicAuth) tokenExpired() bool {
	return false
}

func (a *basicAuth) tokenInBuffer() bool {
	return false
}

func (a *basicAuth) tokenEmpty() bool {
	if a.bearerToken == "" {
		return true
	}
	return false
}
