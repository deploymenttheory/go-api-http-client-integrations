package jamfprointegration

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// JamfAPIHandler implements the APIHandler interface for the Jamf Pro API.
type Integration struct {
	JamfProFQDN          string
	AuthMethodDescriptor string
	Sugar                *zap.SugaredLogger
	auth                 authInterface
	http                 http.Client
	TenantID             string // Jamf Pro instance UUID (required for platform gateway auth)
}

// getFQDN returns just the FQDN // TODO remove the "get"
func (j *Integration) GetFQDN() string {
	return j.JamfProFQDN
}

// constructURL appends any endpoint to the FQDN, rewriting paths when in gateway mode.
func (j *Integration) ConstructURL(endpoint string) string {
	if j.AuthMethodDescriptor == "platform" {
		endpoint = j.rewriteEndpointForGateway(endpoint)
	}
	return j.GetFQDN() + endpoint
}

// rewriteEndpointForGateway translates direct Jamf Pro API paths to platform gateway paths
// with the tenant ID embedded in the URL path.
//   - /JSSResource/... → /api/proclassic/tenant/{id}/...
//   - /api/v{x}/...    → /api/pro/tenant/{id}/v{x}/...
func (j *Integration) rewriteEndpointForGateway(endpoint string) string {
	if strings.HasPrefix(endpoint, "/JSSResource") {
		return fmt.Sprintf("/api/proclassic/tenant/%s%s", j.TenantID, endpoint[len("/JSSResource"):])
	}

	if strings.HasPrefix(endpoint, "/api/v") {
		return fmt.Sprintf("/api/pro/tenant/%s%s", j.TenantID, endpoint[len("/api"):])
	}

	return endpoint
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
func (j *Integration) PrepRequestBody(body any, method string, endpoint string) ([]byte, error) {
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
