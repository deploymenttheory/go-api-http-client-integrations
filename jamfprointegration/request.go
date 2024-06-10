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

	token, err := j.token()
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return nil
}
