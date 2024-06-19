package jamfprointegration

import (
	"fmt"
	"net/http"
)

// TODO func comment
func (j *Integration) prepRequest(req *http.Request) error {
	req.Header.Add("Accept", j.getAcceptHeader())
	req.Header.Add("Content-Type", j.getContentTypeHeader(req.URL.String()))
	req.Header.Add("User-Agent", j.getUserAgentHeader())

	err := j.checkRefreshToken()
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", j.auth.getTokenString()))

	return nil
}
