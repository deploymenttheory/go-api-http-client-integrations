package jamfprointegration

import (
	"errors"
	"time"
)

const (
	tokenEmptyWarnString = "token empty before processing - disregard if first run"
)

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

func (j *Integration) checkRefreshToken() error {
	var err error

	if j.auth.tokenEmpty() {
		j.Logger.Warn(tokenEmptyWarnString)
	}

	if j.auth.tokenExpired() || j.auth.tokenInBuffer() || j.auth.tokenEmpty() {
		err = j.auth.getNewToken()

		if err != nil {
			return err
		}

		// Protects against bad token lifetime/buffer combinations (infinite loops)
		if j.auth.tokenExpired() || j.auth.tokenInBuffer() {
			return errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
		}

		return nil
	}

	return nil
}
