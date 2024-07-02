package jamfprointegration

import (
	"net/http"

	"go.uber.org/zap"
)

// JamfAPIHandler implements the APIHandler interface for the Jamf Pro API.
type Integration struct {
	BaseDomain           string
	AuthMethodDescriptor string
	Sugar                *zap.SugaredLogger
	auth                 authInterface
}

// Info

// TODO migrate strings
func (j *Integration) GetFQDN() string {
	return j.BaseDomain
}

// TODO this comment
func (j *Integration) ConstructURL(endpoint string) string {
	return j.GetFQDN() + endpoint
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
	domain := j.GetFQDN()
	return j.getSessionCookies(domain)
}
