// msgraph/msgraphintegration/builders.go
package msgraphintegration

import (
	"time"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

// BuildIntegrationWithOAuth constructs an Integration instance using OAuth2.0 authentication.
// It sets up the OAuth2.0 authentication method with the provided client ID, client secret, and tenant ID.
// Checks the token refresh status upon creation.
//
// Parameters:
//   - msGraphBaseDomain: The base domain for the Microsoft Graph API.
//   - logger: A logger instance for logging purposes.
//   - bufferPeriod: The buffer period before token expiry to refresh the token.
//   - clientId: The client ID for OAuth2.0 authentication.
//   - clientSecret: The client secret for OAuth2.0 authentication.
//   - tenantID: The tenant ID for the Microsoft Graph API.
//
// Returns:
//   - *Integration: A pointer to the constructed Integration instance.
//   - error: Any error encountered during the token refresh check.
func BuildIntegrationWithOAuth(msGraphBaseDomain string, logger logger.Logger, bufferPeriod time.Duration, clientId string, clientSecret string, tenantID string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           msGraphBaseDomain,
		Logger:               logger,
		AuthMethodDescriptor: "oauth2",
	}

	integration.BuildOAuth(clientId, clientSecret, bufferPeriod, tenantID)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// BuildIntegrationWithBasicAuth constructs an Integration instance using Basic Authentication.
// It sets up the basic authentication method with the provided username, password, and tenant ID.
// Checks the token refresh status upon creation.
//
// Parameters:
//   - msGraphBaseDomain: The base domain for the Microsoft Graph API.
//   - logger: A logger instance for logging purposes.
//   - bufferPeriod: The buffer period before token expiry to refresh the token.
//   - username: The username for basic authentication.
//   - password: The password for basic authentication.
//   - tenantID: The tenant ID for the Microsoft Graph API.
//
// Returns:
//   - *Integration: A pointer to the constructed Integration instance.
//   - error: Any error encountered during the token refresh check.
func BuildIntegrationWithBasicAuth(msGraphBaseDomain string, logger logger.Logger, bufferPeriod time.Duration, username string, password string, tenantID string) (*Integration, error) {
	integration := Integration{
		BaseDomain:           msGraphBaseDomain,
		Logger:               logger,
		AuthMethodDescriptor: "basic",
	}

	integration.BuildBasicAuth(username, password, bufferPeriod, tenantID)
	err := integration.CheckRefreshToken()

	return &integration, err
}

// BuildOAuth sets up the OAuth2.0 authentication method for the Integration instance.
//
// Parameters:
//   - clientId: The client ID for OAuth2.0 authentication.
//   - clientSecret: The client secret for OAuth2.0 authentication.
//   - bufferPeriod: The buffer period before token expiry to refresh the token.
//   - tenantID: The tenant ID for the Microsoft Graph API.
func (j *Integration) BuildOAuth(clientId string, clientSecret string, bufferPeriod time.Duration, tenantID string) {
	authInterface := oauth{
		clientId:     clientId,
		clientSecret: clientSecret,
		bufferPeriod: bufferPeriod,
		baseDomain:   j.BaseDomain,
		Logger:       j.Logger,
		tenantID:     tenantID,
	}

	j.auth = &authInterface
}

// BuildBasicAuth sets up the basic authentication method for the Integration instance.
//
// Parameters:
//   - username: The username for basic authentication.
//   - password: The password for basic authentication.
//   - bufferPeriod: The buffer period before token expiry to refresh the token.
//   - tenantID: The tenant ID for the Microsoft Graph API.
func (j *Integration) BuildBasicAuth(username string, password string, bufferPeriod time.Duration, tenantID string) {
	authInterface := basicAuth{
		username:     username,
		password:     password,
		bufferPeriod: bufferPeriod,
		logger:       j.Logger,
		baseDomain:   j.BaseDomain,
		tenantID:     tenantID,
	}

	j.auth = &authInterface
}
