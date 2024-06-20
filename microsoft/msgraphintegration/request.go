// microsoft/msgraphintegration/request.go
package msgraphintegration

import (
	"fmt"
	"net/http"
)

// prepRequest prepares an HTTP request by setting the necessary headers and handling authorization.
// This function adds headers for Accept, Content-Type, User-Agent, and Authorization based on the Integration's methods
// and checks for token refresh if needed.
//
// Parameters:
//   - req: A pointer to the http.Request that needs to be prepared.
//
// Returns:
//   - error: Any error encountered while checking the refresh token or setting headers. Returns nil if no errors occur.
//
// Functionality:
//   - Adds an "Accept" header based on the Integration's getAcceptHeader method.
//   - Adds a "Content-Type" header based on the Integration's getContentTypeHeader method, which depends on the request URL.
//   - Adds a "User-Agent" header based on the Integration's getUserAgentHeader method.
//   - Checks and refreshes the token if necessary using the Integration's checkRefreshToken method.
//   - Adds an "Authorization" header with a Bearer token obtained from the Integration's auth.getTokenString method.
func (m *Integration) prepRequest(req *http.Request) error {
	req.Header.Add("Accept", m.getAcceptHeader())
	req.Header.Add("Content-Type", m.getContentTypeHeader(req.URL.String()))
	req.Header.Add("User-Agent", m.getUserAgentHeader())

	err := m.checkRefreshToken()
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.auth.getTokenString()))

	return nil
}
