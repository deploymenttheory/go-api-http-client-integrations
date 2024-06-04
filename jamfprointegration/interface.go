// jamfpro_api_handler.go

package jamfprointegration

import (
	"net/http"
	"time"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
)

// JamfAPIHandler implements the APIHandler interface for the Jamf Pro API.
type Integration struct {
	BaseDomain           string // OverrideBaseDomain is used to override the base domain for URL construction.
	InstanceName         string // InstanceName is the name of the Jamf instance.
	Logger               logger.Logger
	AuthMethod           string
	ClientId             string
	ClientSecret         string
	BasicAuthUsername    string
	BasicAuthPassword    string
	AuthMethodDescriptor string
	oauthTokenString     string
	bearerTokenString    string
	tokenExpiry          time.Time
	clientConfig         httpclient.ClientConfig
}

type TokenResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (j *Integration) Token() (string, error) {
	return j.token(j.clientConfig)
}

func (j *Integration) Domain() string {
	return j.BaseDomain
}

func (j *Integration) PrepRequestParamsForIntegration(req *http.Request) error {
	err := j.setRequestHeaders(req)
	return err
}

func (j *Integration) PrepRequestBodyForIntergration(body interface{}, method string, endpoint string) ([]byte, error) {
	return j.marshalRequest(body, method, endpoint)
}

func (j *Integration) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	return j.marshalMultipartRequest(fields, files)
}

func (j *Integration) GetAuthMethodDescriptor() string {
	return j.AuthMethodDescriptor
}
