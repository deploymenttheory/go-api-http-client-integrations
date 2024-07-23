package jamfprointegration

import (
	"time"
)

// oauth implements the authInterface for Oauth2 support
type mockauth struct{}

// TODO migrate strings

// getNewToken updates the held token and expiry information
func (a *mockauth) getNewToken() error {
	return nil
}

// getTokenString returns the current token as a string
func (a *mockauth) getTokenString() string {
	return "mocktoken"
}

// getExpiryTime returns the current token's expiry time as a time.Time var.
func (a *mockauth) getExpiryTime() time.Time {
	out := time.Now().Add(100 * time.Minute)
	return out
}

// tokenExpired returns a bool denoting if the current token expiry time has passed.
func (a *mockauth) tokenExpired() bool {
	return false
}

// tokenInBuffer returns a bool denoting if the current token's duration until expiry is within the buffer period
func (a *mockauth) tokenInBuffer() bool {
	return false
}

// tokenEmpty returns a bool denoting if the current token string is empty.
func (a *mockauth) tokenEmpty() bool {
	return false
}
