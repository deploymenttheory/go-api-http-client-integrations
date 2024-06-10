package jamfprointegration

import (
	"time"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

func BuildIntegrationWithOAuth(jamfBaseDomain string, jamfInstanceName string, logger logger.Logger, bufferPeriod time.Duration, clientId string, clientSecret string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		InstanceName:         jamfInstanceName,
		Logger:               logger,
		AuthMethodDescriptor: "oauth2",
	}

	integration.BuildOAuth(clientId, clientSecret, bufferPeriod)
	err := integration.CheckRefreshToken()

	return &integration, err
}

func BuildIntegrationWithBasicAuth(jamfBaseDomain string, jamfInstanceName string, logger logger.Logger, bufferPeriod time.Duration, username string, password string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		InstanceName:         jamfInstanceName,
		Logger:               logger,
		AuthMethodDescriptor: "basic",
	}

	integration.BuildBasicAuth(username, password, bufferPeriod)
	err := integration.CheckRefreshToken()

	return &integration, err
}

func (j *Integration) BuildOAuth(clientId string, clientSecret string, bufferPeriod time.Duration) {
	authInterface := oauth{
		// args
		clientId:     clientId,
		clientSecret: clientSecret,
		bufferPeriod: bufferPeriod,

		// integration
		baseDomain: j.BaseDomain,
		Logger:     j.Logger,
	}

	j.auth = &authInterface
}

func (j *Integration) BuildBasicAuth(username string, password string, bufferPeriod time.Duration) {
	authInterface := basicAuth{
		username:     username,
		password:     password,
		bufferPeriod: bufferPeriod,

		logger:     j.Logger,
		baseDomain: j.BaseDomain,
	}

	j.auth = &authInterface
}
