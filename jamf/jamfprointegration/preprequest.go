package jamfprointegration

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
func (j *Integration) prepRequest(req *http.Request) error {

	j.Sugar.Debugw("LOG-CONTENT-TYPE", "METHOD", req.Method)
	if req.Method != "READ" && req.Method != "DELETE" {
		req.Header.Add("Content-Type", j.getContentTypeHeader(req.URL.String()))
	}

	req.Header.Add("Accept", j.getAcceptHeader())
	req.Header.Add("User-Agent", j.getUserAgentHeader())

	j.Sugar.Debug("request headers added, refreshing token")

	err := j.checkRefreshToken()
	if err != nil {
		j.Sugar.Warnw("error refreshing token", "error", err)
		return err
	}

	j.Sugar.Debug("token refreshed, setting header")

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", j.auth.getTokenString()))

	return nil
}
