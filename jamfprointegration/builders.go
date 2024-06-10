package jamfprointegration

import "github.com/deploymenttheory/go-api-http-client/logger"

func BuildIntegrationWithOAuth(jamfBaseDomain string, jamfInstanceName string, logger logger.Logger, clientId string, clientSecret string) Integration {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		InstanceName:         jamfInstanceName,
		Logger:               logger,
		AuthMethodDescriptor: "oauth2",
	}

	integration.BuildOAuth(clientId, clientSecret)

	return integration
}

func BuildIntegrationWithBasicAuth(jamfBaseDomain string, jamfInstanceName string, logger logger.Logger, username string, password string) Integration {
	integration := Integration{
		BaseDomain:           jamfBaseDomain,
		InstanceName:         jamfInstanceName,
		Logger:               logger,
		AuthMethodDescriptor: "basic",
	}

	integration.BuildBasicAuth(username, password)

	return integration
}

func (j *Integration) BuildOAuth(clientId string, clientSecret string) {
	authInterface := oauth{
		// args
		clientId:     clientId,
		clientSecret: clientSecret,

		// integration
		baseDomain: j.BaseDomain,
		Logger:     j.Logger,
	}

	j.auth = &authInterface
	j.CheckRefreshToken()

}

func (j *Integration) BuildBasicAuth(username string, password string) {
	authInterface := basicAuth{
		username: username,
		password: password,

		logger:     j.Logger,
		baseDomain: j.BaseDomain,
	}

	j.auth = &authInterface
	j.CheckRefreshToken()
}
