package jamfprointegration

import (
	"errors"
)

const (
	tokenEmptyWarnString = "token empty before processing - disregard if first run"
)

type authInterface interface {
	// Token Operations
	getNewToken() (string, error)
	tokenString() string

	// Token Utils
	tokenExpired() bool
	tokenInBuffer() bool
	tokenEmpty() bool
}

func (j *Integration) token() (string, error) {
	var err error
	var token string

	if j.auth.tokenEmpty() {
		j.Logger.Warn(tokenEmptyWarnString)
	}

	if j.auth.tokenExpired() || j.auth.tokenInBuffer() || j.auth.tokenEmpty() {
		token, err = j.auth.getNewToken()

		if err != nil {
			return "", err
		}

		if j.auth.tokenExpired() || j.auth.tokenInBuffer() {
			return "", errors.New("token lifetime is shorter than buffer period. please adjust parameters.")
		}

		return token, nil
	}

	token = j.auth.tokenString()

	return token, nil
}
