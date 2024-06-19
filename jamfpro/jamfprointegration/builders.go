package jamfprointegration

import (
	"time"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

// TODO migrate strings
func BuildIntegrationWithOAuth(jamfBaseDomain string, logger logger.Logger, bufferPeriod time.Duration, clientId string, clientSecret string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		Logger:               logger,
		AuthMethodDescriptor: "oauth2",
	}

	integration.BuildOAuth(clientId, clientSecret, bufferPeriod)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// TODO migrate strings
func BuildIntegrationWithBasicAuth(jamfBaseDomain string, logger logger.Logger, bufferPeriod time.Duration, username string, password string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		Logger:               logger,
		AuthMethodDescriptor: "basic",
	}

	integration.BuildBasicAuth(username, password, bufferPeriod)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// TODO migrate strings
func (j *Integration) BuildOAuth(clientId string, clientSecret string, bufferPeriod time.Duration) {
	authInterface := oauth{
		clientId:     clientId,
		clientSecret: clientSecret,
		bufferPeriod: bufferPeriod,
		baseDomain:   j.BaseDomain,
		Logger:       j.Logger,
	}

	j.auth = &authInterface
}

// TODO migrate strings
func (j *Integration) BuildBasicAuth(username string, password string, bufferPeriod time.Duration) {
	authInterface := basicAuth{
		username:     username,
		password:     password,
		bufferPeriod: bufferPeriod,
		logger:       j.Logger,
		baseDomain:   j.BaseDomain,
	}

	j.auth = &authInterface
}
