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

func (j *Integration) token(bufferPeriod time.Duration) (string, error) {
	var err error
	var token string
	if j.auth.tokenEmpty() {
		j.Logger.Warn("token empty - disregard if first run")
		if j.auth.tokenExpired() || j.auth.tokenInBuffer(bufferPeriod) || j.auth.tokenEmpty() {
			token, err = j.getOauthToken()

			if j.auth.tokenExpired() || j.auth.tokenInBuffer(bufferPeriod) {
				return "", errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
			}

			if err != nil {
				return "", err
			}

			return token, nil
		}
	}
	return token, nil
}
