package jamfprointegration

import (
	"net/http"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"go.uber.org/zap"
)

// JamfAPIHandler implements the APIHandler interface for the Jamf Pro API.
type Integration struct {
	BaseDomain           string
	AuthMethodDescriptor string
	Sugar                *zap.SugaredLogger
	auth                 authInterface
	httpExecutor         httpclient.HTTPExecutor
}

// getFQDN returns just the FQDN // TODO remove the "get"
func (j *Integration) GetFQDN() string {
	return j.BaseDomain
}

// constructURL appends any endpoint to the FQDN
func (j *Integration) ConstructURL(endpoint string) string {
	return j.GetFQDN() + endpoint
}

// GetAuthMethodDescriptor returns a single string describing the auth method for debug and logging purposes
func (j *Integration) GetAuthMethodDescriptor() string {
	return j.AuthMethodDescriptor
}

// CheckRefreshToken ensures the token is valid and refreshes if it is not.
func (j *Integration) CheckRefreshToken() error {
	return j.checkRefreshToken()
}

// PrepRequestParamsAndAuth applies any parameters and authentication headers to a http.Request
func (j *Integration) PrepRequestParamsAndAuth(req *http.Request) error {
	return j.prepRequest(req)
}

// PrepRequestBody formats body data to meet the API requirements.
func (j *Integration) PrepRequestBody(body interface{}, method string, endpoint string) ([]byte, error) {
	return j.marshalRequest(body, method, endpoint)
}

// TODO this comment
func (j *Integration) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	return j.marshalMultipartRequest(fields, files)
}

// GetSessionCookies retrieves all cookies from the current session
func (j *Integration) GetSessionCookies() ([]*http.Cookie, error) {
	domain := j.GetFQDN()
	return j.getSessionCookies(domain)
}
