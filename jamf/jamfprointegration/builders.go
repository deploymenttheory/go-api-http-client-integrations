package jamfprointegration

import (
	"time"

	"go.uber.org/zap"
)

// BuildWithOAuth is a helper function allowing the full construct of a Jamf Integration using OAuth2
func BuildWithOAuth(jamfBaseDomain string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, clientId string, clientSecret string, hideSensitiveData bool) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		Sugar:                Sugar,
		AuthMethodDescriptor: "oauth2",
	}

	integration.BuildOAuth(clientId, clientSecret, bufferPeriod, hideSensitiveData)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// BuildWithBasicAuth is a helper function allowing the full construct of a Jamf Integration using BasicAuth
func BuildWithBasicAuth(jamfBaseDomain string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, username string, password string, hideSensitiveData bool) (*Integration, error) {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		Sugar:                Sugar,
		AuthMethodDescriptor: "basic",
	}

	integration.BuildBasicAuth(username, password, bufferPeriod, hideSensitiveData)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// BuildOAuth is a helper which returns just a configured OAuth interface
func (j *Integration) BuildOAuth(clientId string, clientSecret string, bufferPeriod time.Duration, hideSensitiveData bool) {
	authInterface := oauth{
		clientId:          clientId,
		clientSecret:      clientSecret,
		bufferPeriod:      bufferPeriod,
		baseDomain:        j.BaseDomain,
		Sugar:             j.Sugar,
		hideSensitiveData: hideSensitiveData,
	}

	j.auth = &authInterface
}

// BuildBasicAuth is a helper which returns just a configured Basic Auth interface/
func (j *Integration) BuildBasicAuth(username string, password string, bufferPeriod time.Duration, hideSensitiveData bool) {
	authInterface := basicAuth{
		username:          username,
		password:          password,
		bufferPeriod:      bufferPeriod,
		Sugar:             j.Sugar,
		baseDomain:        j.BaseDomain,
		hideSensitiveData: hideSensitiveData,
	}

	j.auth = &authInterface
}
