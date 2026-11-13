package jamfprointegration

import (
	"fmt"
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

// BuildWithPlatformGatewayOAuth constructs a Jamf Integration using the platform gateway OAuth2 flow.
// The gateway URL becomes the base FQDN, and all API paths are rewritten transparently.
// The tenantID is sent as the X-Tenant-Id request header.
func BuildWithPlatformGatewayOAuth(gatewayURL string, Sugar *zap.SugaredLogger, bufferPeriod time.Duration, clientId string, clientSecret string, tenantID string, hideSensitiveData bool, client http.Client) (*Integration, error) {
	return BuildWithPlatformGatewayOAuthWithScope(gatewayURL, Sugar, bufferPeriod, clientId, clientSecret, tenantID, "", hideSensitiveData, client)
}

// BuildWithPlatformGatewayOAuthWithScope constructs a Jamf Integration using either tenant or environment scope.
// Exactly one of tenantID and environmentID must be set; the selected ID is sent in its corresponding request header.
func BuildWithPlatformGatewayOAuthWithScope(
	gatewayURL string,
	Sugar *zap.SugaredLogger,
	bufferPeriod time.Duration,
	clientId string,
	clientSecret string,
	tenantID string,
	environmentID string,
	hideSensitiveData bool,
	client http.Client,
) (*Integration, error) {
	if (tenantID == "") == (environmentID == "") {
		return nil, fmt.Errorf("exactly one of tenant ID or environment ID must be set for platform authentication")
	}

	integration := Integration{
		JamfProFQDN:          gatewayURL,
		Sugar:                Sugar,
		AuthMethodDescriptor: "platform",
		http:                 client,
		TenantID:             tenantID,
		EnvironmentID:        environmentID,
	}

	integration.BuildPlatformOAuth(clientId, clientSecret, bufferPeriod, hideSensitiveData, client, gatewayURL)
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

// BuildBasicAuth is a helper which returns just a configured Basic Auth interface.
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

// BuildPlatformOAuth is a helper which returns just a configured platform gateway OAuth interface.
func (j *Integration) BuildPlatformOAuth(clientId string, clientSecret string, bufferPeriod time.Duration, hideSensitiveData bool, client http.Client, gatewayDomain string) {
	authInterface := platformOAuth{
		clientId:          clientId,
		clientSecret:      clientSecret,
		bufferPeriod:      bufferPeriod,
		gatewayDomain:     gatewayDomain,
		Sugar:             j.Sugar,
		hideSensitiveData: hideSensitiveData,
		http:              client,
	}

	j.auth = &authInterface
}
