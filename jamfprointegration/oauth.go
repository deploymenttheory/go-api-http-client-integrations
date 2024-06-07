package jamfprointegration

import (
	"time"
)

// OAuthResponse represents the response structure when obtaining an OAuth access token from JamfPro.
type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type oauth struct {
	clientId     string
	clientSecret string
	expiryTime   time.Time
}
