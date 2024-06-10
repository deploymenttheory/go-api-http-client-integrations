package jamfprointegration

import (
	"errors"
	"time"
)

type auth interface {
	tokenExpired() bool
	tokenInBuffer(tokenRefreshBufferPeriod time.Duration) bool
	tokenEmpty() bool
	getNewToken() (string, error)
}

func (j *Integration) token(bufferPeriod time.Duration) (string, error) {
	var err error
	var token string

	if j.auth.tokenEmpty() {
		j.Logger.Warn("token empty before processing - disregard if first run")
	}

	if j.auth.tokenExpired() || j.auth.tokenInBuffer(bufferPeriod) || j.auth.tokenEmpty() {
		token, err = j.auth.getNewToken()

		if err != nil {
			return "", err
		}

		if j.auth.tokenExpired() || j.auth.tokenInBuffer(bufferPeriod) {
			return "", errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
		}

		return "", err
	}

	return token, nil
}
