package msgraphintegration

import (
	"errors"
	"time"
)

const (
	tokenEmptyWarnString = "token empty before processing - disregard if first run"
)

// authInterface defines the methods required to satify the authentication interface.
type authInterface interface {
	// Token Operations
	getNewToken() error
	getTokenString() string
	getExpiryTime() time.Time

	// Token Utils
	tokenExpired() bool
	tokenInBuffer() bool
	tokenEmpty() bool
}

// checkRefreshToken checks and refreshes the authentication token if necessary.
// This function ensures that the authentication token is valid and not expired. If the token is empty, expired,
// or within the buffer period before expiry, it attempts to obtain a new token and validates the new token's lifetime
// against the buffer period to prevent infinite loops.
//
// Returns:
//   - error: Any error encountered during the token refresh process or if the token's lifetime is shorter than the buffer period. Returns nil if no errors occur.
//
// Functionality:
//   - Logs a warning if the token is empty.
//   - Checks if the token is expired, within the buffer period, or empty.
//   - Attempts to obtain a new token if the current token is invalid.
//   - Validates the new token's lifetime against the buffer period to prevent bad token lifetime/buffer combinations.
//   - Returns an error if the token refresh fails or if the new token's lifetime is shorter than the buffer period.
func (m *Integration) checkRefreshToken() error {
	var err error

	if m.auth.tokenEmpty() {
		m.Logger.Warn(tokenEmptyWarnString)
	}

	if m.auth.tokenExpired() || m.auth.tokenInBuffer() || m.auth.tokenEmpty() {
		err = m.auth.getNewToken()

		if err != nil {
			return err
		}

		// Protects against bad token lifetime/buffer combinations (infinite loops)
		if m.auth.tokenExpired() || m.auth.tokenInBuffer() {
			return errors.New("token lifetime is shorter than buffer period. please adjust parameters")
		}

		return nil
	}

	return nil
}
