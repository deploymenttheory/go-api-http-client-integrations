// microsoft/msgraphintegration/interface.go
package msgraphintegration

import (
	"net/http"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

// Integration implements the APIHandler interface for the Microsoft Graph API.
type Integration struct {
	TenantID             string // TenantID used for constructing the authentication endpoint.
	TenantName           string // TenantName used for constructing the authentication endpoint.
	AuthMethodDescriptor string
	Logger               logger.Logger
	auth                 authInterface
}

// Info

// GetFQDN returns the fully qualified domain name (FQDN) for Microsoft Graph.
func (m *Integration) GetFQDN() string {
	return m.getFQDN()
}

// ConstructURL constructs a full URL for a given Microsoft Graph endpoint.
//
// Parameters:
//   - endpoint: The API endpoint to be appended to the FQDN.
//
// Returns:
//   - string: The fully constructed URL.
func (j *Integration) ConstructURL(endpoint string) string {
	return j.GetFQDN() + endpoint
}

// GetAuthMethodDescriptor returns the authentication method descriptor.
//
// Returns:
//   - string: The descriptor of the authentication method used.
func (m *Integration) GetAuthMethodDescriptor() string {
	return m.AuthMethodDescriptor
}

// Utilities

// CheckRefreshToken checks and refreshes the authentication token if necessary.
//
// Returns:
//   - error: Any error encountered during the token refresh process. Returns nil if no errors occur.
func (m *Integration) CheckRefreshToken() error {
	return m.checkRefreshToken()
}

// PrepRequestParamsAndAuth prepares the request parameters and handles authentication.
//
// Parameters:
//   - req: A pointer to the http.Request that needs to be prepared.
//
// Returns:
//   - error: Any error encountered during request preparation. Returns nil if no errors occur.
func (m *Integration) PrepRequestParamsAndAuth(req *http.Request) error {
	err := m.prepRequest(req)
	return err
}

// PrepRequestBody marshals the request body as JSON.
//
// Parameters:
//   - body: The request body to be marshaled.
//   - method: The HTTP method being used for the request (e.g., "POST", "PUT", "PATCH").
//   - endpoint: The API endpoint for the request.
//
// Returns:
//   - []byte: The marshaled JSON byte slice of the request body.
//   - error: Any error encountered during the marshaling process.
func (m *Integration) PrepRequestBody(body interface{}, method string, endpoint string) ([]byte, error) {
	return m.marshalRequest(body, method, endpoint)
}

// MarshalMultipartRequest handles multipart form data encoding.
//
// Parameters:
//   - fields: A map of form fields to be included in the request.
//   - files: A map of file paths to be included in the request.
//
// Returns:
//   - []byte: The marshaled multipart form data.
//   - string: The content type of the multipart form data.
//   - error: Any error encountered during the marshaling process.
func (m *Integration) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	return m.marshalMultipartRequest(fields, files)
}
