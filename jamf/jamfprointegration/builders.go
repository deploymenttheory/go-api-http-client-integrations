package jamfprointegration

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// BuildWithOAuth is a helper function allowing the full construct of a Jamf Integration using OAuth2
func BuildWithOAuth(jamfProFQDN string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, clientId string, clientSecret string, hideSensitiveData bool, client http.Client) (*Integration, error) {
	integration := Integration{
		JamfProFQDN:          jamfProFQDN,
		Sugar:                Sugar,
		AuthMethodDescriptor: "oauth2",
		http:                 client,
	}

	integration.BuildOAuth(clientId, clientSecret, bufferPeriod, hideSensitiveData, client)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// BuildWithBasicAuth is a helper function allowing the full construct of a Jamf Integration using BasicAuth
func BuildWithBasicAuth(jamfProFQDN string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, username string, password string, hideSensitiveData bool, client http.Client) (*Integration, error) {

	integration := Integration{
		JamfProFQDN:          jamfProFQDN,
		Sugar:                Sugar,
		AuthMethodDescriptor: "basic",
		http:                 client,
	}

	integration.BuildBasicAuth(username, password, bufferPeriod, hideSensitiveData, client)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// BuildOAuth is a helper which returns just a configured OAuth interface
func (j *Integration) BuildOAuth(clientId string, clientSecret string, bufferPeriod time.Duration, hideSensitiveData bool, client http.Client) {
	authInterface := oauth{
		clientId:          clientId,
		clientSecret:      clientSecret,
		bufferPeriod:      bufferPeriod,
		baseDomain:        j.JamfProFQDN,
		Sugar:             j.Sugar,
		hideSensitiveData: hideSensitiveData,
		http:              client,
	}

	j.auth = &authInterface
}

// BuildBasicAuth is a helper which returns just a configured Basic Auth interface/
func (j *Integration) BuildBasicAuth(username string, password string, bufferPeriod time.Duration, hideSensitiveData bool, client http.Client) {
	authInterface := basicAuth{
		username:          username,
		password:          password,
		bufferPeriod:      bufferPeriod,
		Sugar:             j.Sugar,
		baseDomain:        j.JamfProFQDN,
		hideSensitiveData: hideSensitiveData,
		http:              client,
	}

	j.auth = &authInterface
}
