package jamfprointegration

import (
	"errors"
	"time"

	"go.uber.org/zap"
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

	j.Logger.Debug("Checking and refreshing token")

	if j.auth.tokenEmpty() {
		j.Logger.Warn(tokenEmptyWarnString)
	}

	j.Logger.Debug("Bools:", zap.Bool("expired", j.auth.tokenExpired()), zap.Bool("in buffer", j.auth.tokenInBuffer()), zap.Bool("empty", j.auth.tokenEmpty()))
	j.Logger.Debug("Vars", zap.String("exp time", j.auth.getExpiryTime().String()))

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
