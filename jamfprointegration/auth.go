package jamfprointegration

import (
	"errors"
	"time"
)

type auth interface {
	tokenExpired() bool
	tokenInBuffer(tokenRefreshBufferPeriod time.Duration) bool
	tokenEmpty() bool
	getNewToken() error
}

func (j *Integration) token(bufferPeriod time.Duration) error {
	var err error
	if j.auth.tokenEmpty() {
		j.Logger.Warn("token empty - disregard if first run")
		if j.auth.tokenExpired() || j.auth.tokenInBuffer(bufferPeriod) || j.auth.tokenEmpty() {
			err = j.auth.getNewToken()

			if j.auth.tokenExpired() || j.auth.tokenInBuffer(bufferPeriod) {
				return errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
			}

			if err != nil {
				return err
			}

			return nil
		}
	}
	return nil
}
