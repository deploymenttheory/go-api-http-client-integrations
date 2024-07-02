package jamfprointegration

import (
	"time"

	"go.uber.org/zap"
)

// TODO migrate strings
func BuildWithOAuth(jamfBaseDomain string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, clientId string, clientSecret string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		Sugar:                Sugar,
		AuthMethodDescriptor: "oauth2",
	}

	integration.BuildOAuth(clientId, clientSecret, bufferPeriod)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// TODO migrate strings
func BuildWithBasicAuth(jamfBaseDomain string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, username string, password string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		Sugar:                Sugar,
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
		Sugar:        j.Sugar,
	}

	j.auth = &authInterface
}

// TODO migrate strings
func (j *Integration) BuildBasicAuth(username string, password string, bufferPeriod time.Duration) {
	authInterface := basicAuth{
		username:     username,
		password:     password,
		bufferPeriod: bufferPeriod,
		Sugar:        j.Sugar,
		baseDomain:   j.BaseDomain,
	}

	j.auth = &authInterface
}
