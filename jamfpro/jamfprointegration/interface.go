// jamfpro_api_handler.go

package jamfprointegration

import (
	"net/http"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

// JamfAPIHandler implements the APIHandler interface for the Jamf Pro API.
type Integration struct {
	BaseDomain           string
	AuthMethodDescriptor string
	Logger               logger.Logger
	auth                 authInterface
}

// Info

// TODO migrate strings
func (j *Integration) Domain() string {
	return j.BaseDomain
}

// TODO migrate strings
func (j *Integration) GetAuthMethodDescriptor() string {
	return j.AuthMethodDescriptor
}

// Utilities

// TODO migrate strings
func (j *Integration) CheckRefreshToken() error {
	return j.checkRefreshToken()
}

// TODO migrate strings
func (j *Integration) PrepRequestParamsAndAuth(req *http.Request) error {
	err := j.prepRequest(req)
	return err
}

// TODO migrate strings
func (j *Integration) PrepRequestBody(body interface{}, method string, endpoint string) ([]byte, error) {
	return j.marshalRequest(body, method, endpoint)
}

// TODO migrate strings
func (j *Integration) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	return j.marshalMultipartRequest(fields, files)
}

// TODO migrate strings
func (j *Integration) GetSessionCookies() ([]*http.Cookie, error) {
	domain := j.Domain()
	return j.getSessionCookies(domain)
}
