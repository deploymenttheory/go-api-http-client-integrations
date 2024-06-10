package jamfprointegration

import (
	"errors"

	"go.uber.org/zap"
)

const (
	tokenEmptyWarnString = "token empty before processing - disregard if first run"
)

type authInterface interface {
	// Token Operations
	getNewToken() error
	getTokenString() string

	// Token Utils
	tokenExpired() bool
	tokenInBuffer() bool
	tokenEmpty() bool
}

func (j *Integration) checkRefreshToken() error {
	var err error

	j.Logger.Debug("Checking and refreshing token")

	if j.auth.tokenEmpty() {
		j.Logger.Warn(tokenEmptyWarnString)
	}

	if j.auth.tokenExpired() || j.auth.tokenInBuffer() || j.auth.tokenEmpty() {
		err = j.auth.getNewToken()

		if err != nil {
			j.Logger.Warn("WARNING", zap.Error(err))
			return err
		}
		j.Logger.Warn("VARS: ", zap.Bool("expired:", j.auth.tokenExpired()), zap.Bool("buffer: ", j.auth.tokenInBuffer()))
		// Protects against bad token lifetime/buffer combinations (infinite loops)
		if j.auth.tokenExpired() || j.auth.tokenInBuffer() {
			j.Logger.Warn("INSIDE CATCH")
			return errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
		}

		return nil
	}

	return nil
}
