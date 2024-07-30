package test

import (
	"net/http"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"go.uber.org/zap"
)

const (
	ENV_KEY_JAMFPRO_FQDN  = "TEST_JAMFPRO_FQDN"
	ENV_KEY_CLIENT_ID     = "TEST_JAMFPRO_CLIENT_ID"
	ENV_KEY_CLIENT_SECRET = "TEST_JAMFPRO_CLIENT_SECRET"
	ENV_KEY_USERNAME      = "TEST_JAMFPRO_USERNAME"
	ENV_KEY_PASSWORD      = "TEST_JAMFPRO_PASSWORD"
)

func NewSugaredDevelopmentLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}

func NewIntegrationFromEnv() *jamfprointegration.Integration {
	integration, err := jamfprointegration.BuildWithOAuth(
		os.Getenv(ENV_KEY_JAMFPRO_FQDN),
		NewSugaredDevelopmentLogger(),
		10*time.Second,
		os.Getenv(ENV_KEY_CLIENT_ID),
		os.Getenv(ENV_KEY_CLIENT_SECRET),
		false,
		&httpclient.ProdExecutor{Client: &http.Client{}},
	)

	if err != nil {
		panic("we have a problem")
	}

	return integration
}
