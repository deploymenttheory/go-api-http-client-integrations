// apiintegrations/msgraph/request.go
package msgraphintegration

import (
	"fmt"
	"net/http"
)

// TODO func comment
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
